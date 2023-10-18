package entities

import (
	"errors"
	"github.com/google/uuid"
	"strings"
)

type Bucket struct {
	ID       string  `json:"id" gorm:"primaryKey;type:uuid"`
	Name     string  `json:"name"`
	ParentID *string `json:"parent_id" gorm:"type:uuid"`
}

func NewBucket(
	id string,
	name string,
	parentID *string,
) (*Bucket, error) {
	if len(strings.TrimSpace(id)) == 0 {
		return nil, errors.New("id should not be empty")
	}

	if _, err := uuid.Parse(id); err != nil {
		return nil, errors.New("id should be a uuid")
	}

	if len(strings.TrimSpace(name)) == 0 {
		return nil, errors.New("name should not be empty")
	}

	if parentID != nil {
		if len(strings.TrimSpace(*parentID)) == 0 {
			return nil, errors.New("parent id should not be empty")
		}

		if _, err := uuid.Parse(*parentID); err != nil {
			return nil, errors.New("parent id should be a uuid")
		}
	}

	return &Bucket{
		ID:       id,
		Name:     name,
		ParentID: parentID,
	}, nil
}

func (that *Bucket) InRoot() bool {
	return that.ParentID == nil
}
