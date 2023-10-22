package repositories

import (
	"github.com/jibaru/gostore/internal/domain/entities"
)

type BucketRepository interface {
	Save(bucket entities.Bucket) error
	GetAll() ([]entities.Bucket, error)
	FindByID(ID string) (*entities.Bucket, error)
	GetByParentID(parentID string) ([]entities.Bucket, error)
}
