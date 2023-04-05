package db

import (
	"time"

	"gorm.io/gorm"
)

type Provider interface {
}

type GormProvider struct {
	db_main      *gorm.DB
	jwt_duration time.Duration
}

func NewProvider(db *gorm.DB, jwtDuration time.Duration) *GormProvider {

	return &GormProvider{db_main: db, jwt_duration: jwtDuration}

}
