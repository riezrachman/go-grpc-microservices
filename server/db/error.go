package db

import "fmt"

type NotFoundError struct {
	ResourceType string
	ID           string
}

func NotFound(resourceType string, id string) NotFoundError {
	return NotFoundError{ResourceType: resourceType, ID: id}
}

func (err NotFoundError) Error() string {
	return fmt.Sprintf("%s with id '%s' not found", err.ResourceType, err.ID)
}
