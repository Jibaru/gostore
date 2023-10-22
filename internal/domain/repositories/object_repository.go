package repositories

import (
	"github.com/jibaru/gostore/internal/domain/entities"
)

type ObjectRepository interface {
	Save(object entities.Object) error
	FindByID(ID string) (*entities.Object, error)
	GetByBucketID(bucketID string) ([]entities.Object, error)
	DeleteByID(objectID string) error
}
