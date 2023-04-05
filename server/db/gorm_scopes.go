package db

import (
	"fmt"
	"math"
	"strings"

	pb "swing/server/pb"

	"gorm.io/gorm"
)

type ListFilter struct {
	Data        interface{}
	Filter      string
	Query       string
	CustomOrder string
	Sort        *pb.Sort
	WithTrashed bool
}

func Paginate(value interface{}, v *pb.PaginationResponse, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	if v.Limit > 0 || v.Page > 0 {
		var totalRows int64
		db.Model(value).Count(&totalRows)

		v.TotalRows = totalRows
		totalPages := int(math.Ceil(float64(totalRows) / float64(v.Limit)))
		v.TotalPages = int32(totalPages)
	}

	return func(db *gorm.DB) *gorm.DB {
		if v.Limit < 1 || v.Page < 1 {
			return db
		}

		offset := (v.Page - 1) * v.Limit
		if v != nil {
			return db.Limit(int(v.Limit)).Offset(int(offset))
		}
		return db
	}
}

func Sort(v *pb.Sort) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if v == nil || v.Column == "" {
			return db
		}

		isArray := false

		isObject := false
		if len(v.Column) > 2 {
			if string(v.Column[0:2]) == "[]" {
				isArray = true
				isObject = true
				v.Column = v.Column[2:]
			}
		}

		colums := strings.Split(v.Column, ",")
		if len(colums) > 1 {
			for _, col := range colums {
				column := columnNameBuilder(col, isObject)
				if v.Direction != "" {
					if v.Direction == "DESC" {
						db = db.Order(column + " " + v.Direction + " NULLS LAST")
					} else {
						db = db.Order(column + " NULLS FIRST")
					}
				} else {
					db = db.Order(column + " NULLS FIRST")
				}
			}
		} else {
			v.Column = columnNameBuilder(v.Column, isObject)
			if isArray {
				v.Column = fmt.Sprintf("jsonb_array_length(%s)", v.Column)
				if v.Direction == "DESC" {
					return db.Order(v.Column + " " + v.Direction + " NULLS LAST")
				} else {
					return db.Order(v.Column + " NULLS FIRST")
				}
			} else {
				if v.Direction == "DESC" {
					return db.Order(v.Column + " " + v.Direction + " NULLS LAST")
				} else {
					return db.Order(v.Column + " NULLS FIRST")
				}
			}
		}
		return db
	}
}

func QueryScoop(v string) func(db *gorm.DB) *gorm.DB {
	// Example of v: "key1,key2,key3:val"
	return func(db *gorm.DB) *gorm.DB {
		if v == "" {
			return db
		}

		query := strings.Split(v, ":")
		if len(query) < 2 {
			return db
		}

		columns := strings.Split(query[0], ",")
		expression := ""
		value := query[1]

		dbQuery := db.Session(&gorm.Session{NewDB: true})

		if len(query[1]) > 2 {
			if string(query[1][0:1]) == "%" {
				expression = string(query[1][0:2])
			}
		}

		switch expression {
		case "%%":
			value := "%" + string(query[1][2:len(query[1])]) + "%"
			dbQuery = queryColumnsLoop(dbQuery, columns, "LIKE", value)

		case "%!":
			value := "%" + string(query[1][2:len(query[1])]) + "%"
			dbQuery = queryColumnsLoop(dbQuery, columns, "ILIKE", value)

		case "":
			dbQuery = queryColumnsLoop(dbQuery, columns, "=", value)
		}

		return db.Where(dbQuery)
	}
}

func WhereInScoop(v string) func(db *gorm.DB) *gorm.DB {
	// Example of v: "key:val1,val2,val3"
	return func(db *gorm.DB) *gorm.DB {
		if v == "" {
			return db
		}

		query := strings.Split(v, ":")
		if len(query) < 2 {
			return db
		}

		column := columnNameBuilder(query[0], false)
		values := query[1]
		if len(query) > 2 {
			values = strings.Join(query[1:], ":")
		}

		not := false
		if string(query[1][0:1]) == "!" {
			not = true
			values = query[1][1:len(query[1])]
		}

		inVals := strings.Split(values, ",")

		if not {
			db = db.Where(column+" NOT IN (?)", inVals)
		} else {
			db = db.Where(column+" IN (?)", inVals)
		}

		return db
	}
}

func queryColumnsLoop(db *gorm.DB, columns []string, expresion string, value string) *gorm.DB {
	for i, s := range columns {
		s = columnNameBuilder(s, false)
		if i == 0 {
			db = db.Where(fmt.Sprintf("%s %s ?", s, expresion), value)
		} else {
			db = db.Or(fmt.Sprintf("%s %s ?", s, expresion), value)
		}
	}
	return db
}

func reviewedByHandler(val string, expresion string, db *gorm.DB) *gorm.DB {
	approved := db.Session(&gorm.Session{NewDB: true})
	rejected := db.Session(&gorm.Session{NewDB: true})
	// filter := db.Session(&gorm.Session{NewDB: true})
	query := db.Session(&gorm.Session{NewDB: true})

	// filter = filter.Where("\"last_approved_by_name\" != '' AND \"status\" = '1'")
	approvedQuery := fmt.Sprintf("\"last_approved_by_name\" %s '%s'", expresion, val)
	approved = approved.Where(approvedQuery)

	rejectedQuery := fmt.Sprintf("\"last_rejected_by_name\" %s '%s' AND \"status\" != '1'", expresion, val)
	rejected = rejected.Where(rejectedQuery)

	db = db.Where(query.Where(approved).Or(rejected)).Where("\"status\" != '1' AND \"status\" != '2' AND \"status\" != '6'")
	return db
}

func NotConditionalScoope(v string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if v == "" {
			return db
		}

		query := strings.Split(v, ",")
		if len(query) < 3 {
			return db
		}

		column := query[0]
		symbol := query[1]
		value := query[1]

		allowedSymbols := []string{"?"}
		isAllowed := false
		for _, v := range allowedSymbols {
			if v == symbol {
				isAllowed = true
			}
		}
		if !isAllowed {
			return db
		}

		db = db.Where(fmt.Sprintf("NOT %s %s ?", column, symbol), value)

		return db
	}
}

func FilterScoope(v string) func(db *gorm.DB) *gorm.DB {
	// Example of v: "key1:val1,key2:val2"
	return func(db *gorm.DB) *gorm.DB {
		if v == "" {
			return db
		}

		filters := strings.Split(v, ",")
		if len(filters) < 1 {
			return db
		}

		for _, s := range filters {
			filter := strings.Split(s, ":")
			if len(filter) >= 2 {
				keyword := filter[1]
				if len(filter) > 2 {
					keyword = strings.Join(filter[1:], ":")
				}

				expression := ""
				if len(filter[1]) > 2 {
					if string(filter[1][0:1]) == "%" {
						expression = string(filter[1][0:2])
					} else if string(filter[1][0:1]) == ">" {
						expression = string(filter[1][0:1])
					} else if string(filter[1][0:1]) == "<" {
						expression = string(filter[1][0:1])
						if filter[1][1:2] == "@" {
							expression = string(filter[1][0:2])
						}
					} else if string(filter[1][0:1]) == "@" {
						expression = string(filter[1][0:2])
					}
				}

				column := columnNameBuilder(filter[0], false)
				if expression == "%%" {
					value := "%" + string(keyword[2:len(filter[1])]) + "%"
					if column != "\"reviewed_by\"" {
						db = db.Where(fmt.Sprintf("%s LIKE ?", column), value)
					} else {
						db = reviewedByHandler(value, "LIKE", db)
					}
				} else if expression == "%!" {
					value := "%" + string(keyword[2:len(filter[1])]) + "%"
					if column != "\"reviewed_by\"" {
						db = db.Where(fmt.Sprintf("%s ILIKE ?", column), value)
					} else {
						db = reviewedByHandler(value, "ILIKE", db)
					}
				} else if expression == "!%" {
					value := "%" + string(keyword[2:len(filter[1])]) + "%"
					if column != "\"reviewed_by\"" {
						db = db.Where(fmt.Sprintf("%s NOT LIKE ?", column), value)
					} else {
						db = reviewedByHandler(value, "NOT LIKE", db)
					}
				} else if expression == "@>" {
					value := string(keyword[2:])
					column := columnNameBuilder(filter[0], true)
					db = db.Where(fmt.Sprintf("%s @> ?", column), value)
				} else if expression == "<@" {
					value := string(keyword[2:])
					column := columnNameBuilder(filter[0], true)
					db = db.Where(fmt.Sprintf("%s <@ ?", column), value)
				} else if expression == ">" || expression == "<" {
					if expression == "<" && filter[1][1:2] == ">" {
						expression = string(filter[1][0:2])
						value := string(keyword[2:len(filter[1])])
						db = db.Where(fmt.Sprintf("%s %s ?", column, expression), value)
					} else if filter[1][1:2] == "=" {
						expression = string(filter[1][0:2])
						value := string(keyword[2:len(filter[1])])
						db = db.Where(fmt.Sprintf("%s %s ?", column, expression), value)
					} else {
						value := string(keyword[1:len(filter[1])])
						db = db.Where(fmt.Sprintf("%s %s ?", column, expression), value)
					}
				} else if keyword == "true" || keyword == "false" {
					value := keyword
					if value == "true" {
						db = db.Where(fmt.Sprintf("%s = ?", column), value)
					} else {
						if strings.Contains(column, "->") {
							groupQ := db.Session(&gorm.Session{NewDB: true})
							groupQ = groupQ.Where(fmt.Sprintf("%s IS NULL", column))
							groupQ = groupQ.Or(fmt.Sprintf("%s = ?", column), value)
							db = db.Where(groupQ)
						} else {
							db = db.Where(fmt.Sprintf("%s = ?", column), value)
						}
					}
				} else {
					value := keyword
					if column != "\"reviewed_by\"" {
						db = db.Where(fmt.Sprintf("%s = ?", column), value)
					} else {
						db = reviewedByHandler(value, "=", db)
					}
				}
			}
		}

		return db
	}
}

func CustomOrderScoop(v string) func(db *gorm.DB) *gorm.DB {
	// Example of v: "key|>|1,2,3,4,5" "key|<|1,2,3,4,5"
	return func(db *gorm.DB) *gorm.DB {
		if v == "" {
			return db
		}

		in := strings.Split(v, "|")
		if len(in) < 3 {
			return db
		}

		key := columnNameBuilder(in[0], false)
		direction := in[1]
		valJoin := strings.Join(in[2:], "|")
		valArray := strings.Split(valJoin, ",")
		if len(valArray) < 1 {
			return db
		}
		if direction == "<" {
			valArray = reverseArrayString(valArray)
		}

		orderByQuery := ""
		for i, v := range valArray {
			if v != "" {
				if i == 0 {
					orderByQuery += fmt.Sprintf("%s!= '%s'", key, v)
				} else {
					orderByQuery += fmt.Sprintf(", %s!= '%s'", key, v)
				}
			} else {
				if i == 0 {
					orderByQuery += fmt.Sprintf("%s", key)
				} else {
					orderByQuery += fmt.Sprintf(", %s", key)
				}
			}
		}

		db = db.Order(orderByQuery)

		return db
	}
}

func FilterOrScoope(v string, customOr string) func(db *gorm.DB) *gorm.DB {
	// Example of v: "key1:val1|key2:val2"
	return func(db *gorm.DB) *gorm.DB {
		if v == "" {
			return db
		}

		filters := strings.Split(v, "|")
		if len(filters) < 1 {
			return db
		}

		dbQuery := db.Session(&gorm.Session{NewDB: true})

		for _, s := range filters {
			filter := strings.Split(s, ":")
			if len(filter) >= 2 {
				keyword := filter[1]
				if len(filter) > 2 {
					keyword = strings.Join(filter[1:], ":")
				}

				expression := ""
				if len(filter[1]) > 2 {
					if string(filter[1][0:1]) == "%" {
						expression = string(filter[1][0:2])
					} else if string(filter[1][0:1]) == ">" {
						expression = string(filter[1][0:1])
					} else if string(filter[1][0:1]) == "<" {
						expression = string(filter[1][0:1])
						if filter[1][1:2] == "@" {
							expression = string(filter[1][0:2])
						}
					} else if string(filter[1][0:1]) == "@" {
						expression = string(filter[1][0:2])
					}
				}

				column := columnNameBuilder(filter[0], false)
				if expression == "%%" {
					value := "%" + string(keyword[2:len(filter[1])]) + "%"
					dbQuery = dbQuery.Or(fmt.Sprintf("%s LIKE ?", column), value)
				} else if expression == "%!" {
					value := "%" + string(keyword[2:len(filter[1])]) + "%"
					dbQuery = dbQuery.Or(fmt.Sprintf("%s ILIKE ?", column), value)
				} else if expression == "@>" {
					value := string(keyword[2:])
					column := columnNameBuilder(filter[0], true)
					dbQuery = dbQuery.Or(fmt.Sprintf("%s @> ?", column), value)
				} else if expression == "<@" {
					value := string(keyword[2:])
					column := columnNameBuilder(filter[0], true)
					dbQuery = dbQuery.Or(fmt.Sprintf("%s <@ ?", column), value)
				} else if expression == ">" || expression == "<" {
					if expression == "<" && filter[1][1:2] == ">" {
						expression = string(filter[1][0:2])
						value := string(keyword[2:len(filter[1])])
						dbQuery = dbQuery.Or(fmt.Sprintf("%s %s ?", column, expression), value)
					} else if filter[1][1:2] == "=" {
						expression = string(filter[1][0:2])
						value := string(keyword[2:len(filter[1])])
						dbQuery = dbQuery.Or(fmt.Sprintf("%s %s ?", column, expression), value)
					} else {
						value := string(keyword[1:len(filter[1])])
						dbQuery = dbQuery.Or(fmt.Sprintf("%s %s ?", column, expression), value)
					}
				} else if keyword == "true" || keyword == "false" {
					value := keyword
					if value == "true" {
						dbQuery = dbQuery.Or(fmt.Sprintf("%s = ?", column), value)
					} else {
						if strings.Contains(column, "->") {
							groupQ := db.Session(&gorm.Session{NewDB: true})
							groupQ = groupQ.Where(fmt.Sprintf("%s IS NULL", column))
							groupQ = groupQ.Or(fmt.Sprintf("%s = ?", column), value)
							dbQuery = dbQuery.Or(groupQ)
						} else {
							dbQuery = dbQuery.Or(fmt.Sprintf("%s = ?", column), value)
						}
					}
				} else {
					value := keyword
					dbQuery = dbQuery.Or(fmt.Sprintf("%s = ?", column), value)

				}
			}
		}

		if customOr != "" {
			dbQuery = dbQuery.Or(customOr)
		}

		db = db.Where(dbQuery)
		return db
	}
}

func DistinctScoope(v string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if v == "" {
			return db
		}

		keys := strings.Split(v, ",")

		for i, v := range keys {
			keys[i] = columnNameBuilder(v, false)
		}

		return db.Distinct(keys)
	}
}

func reverseArrayString(arr []string) []string {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

// func reverseArray(arr []int) []int {
// 	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
// 		arr[i], arr[j] = arr[j], arr[i]
// 	}
// 	return arr
// }

func columnNameBuilder(s string, isObject bool) string {
	if strings.Contains(s, "->") {
		nested := strings.Split(s, "->")
		s = ""
		for i, t := range nested {
			if i == 0 {
				s = fmt.Sprintf("\"%s\"", t)
			} else if i == len(nested)-1 && !isObject {
				s = s + fmt.Sprintf("->>'%s'", t)
			} else {
				s = s + fmt.Sprintf("->'%s'", t)
			}
		}
	} else if strings.Contains(s, ".") {
		nested := strings.Split(s, ".")
		s = ""
		for i, t := range nested {
			if i == 0 {
				s = fmt.Sprintf("\"%s\"", t)
			} else if i == len(nested)-1 && !isObject {
				s = s + fmt.Sprintf("->>'%s'", t)
			} else {
				if t == "0" {
					s = s + fmt.Sprintf("->%s", t)
				} else {
					s = s + fmt.Sprintf("->'%s'", t)
				}
			}
		}
	} else {
		s = fmt.Sprintf("\"%s\"", s)
	}
	return s
}
