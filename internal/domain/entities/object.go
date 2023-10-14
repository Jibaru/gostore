package entities

import (
	"errors"
	"strings"
)

type Object struct {
	ID       string `gorm:"primaryKey;type:uuid"`
	Name     string
	BucketID string `gorm:"type:uuid"`
}

func NewObject(
	id string,
	name string,
	bucketID string,
) (*Object, error) {
	if len(strings.TrimSpace(name)) == 0 {
		return nil, errors.New("name should not be empty")
	}

	if len(strings.TrimSpace(bucketID)) == 0 {
		return nil, errors.New("bucket id should not be empty")
	}

	return &Object{
		ID:       id,
		Name:     name,
		BucketID: bucketID,
	}, nil
}
