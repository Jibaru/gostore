package entities

import (
	"errors"
	"github.com/google/uuid"
	"strings"
)

type Object struct {
	ID        string `gorm:"primaryKey;type:uuid" json:"id"`
	Name      string `json:"name"`
	Extension string `json:"extension"`
	BucketID  string `gorm:"type:uuid" json:"bucket_id"`
}

func NewObject(
	id string,
	name string,
	extension string,
	bucketID string,
) (*Object, error) {
	if len(strings.TrimSpace(id)) == 0 {
		return nil, errors.New("id should not be empty")
	}

	if _, err := uuid.Parse(id); err != nil {
		return nil, errors.New("id should be a uuid")
	}

	if len(strings.TrimSpace(name)) == 0 {
		return nil, errors.New("name should not be empty")
	}

	if len(strings.TrimSpace(bucketID)) == 0 {
		return nil, errors.New("bucket id should not be empty")
	}

	if _, err := uuid.Parse(bucketID); err != nil {
		return nil, errors.New("bucket id should be a uuid")
	}

	if len(strings.TrimSpace(extension)) == 0 && strings.TrimSpace(extension) != "." {
		return nil, errors.New("extension should not be empty")
	}

	return &Object{
		ID:        id,
		Name:      name,
		Extension: extension,
		BucketID:  bucketID,
	}, nil
}
