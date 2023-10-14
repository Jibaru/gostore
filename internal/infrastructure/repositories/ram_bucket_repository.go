package repositories

import (
	"errors"
	"github.com/jibaru/gostore/internal/domain/entities"
	"github.com/jibaru/gostore/internal/domain/repositories"
)

type RamBucketRepository struct {
	buckets []entities.Bucket
}

func NewRamBucketRepository() repositories.BucketRepository {
	objects := make([]entities.Bucket, 0)
	objects = append(objects, entities.Bucket{
		ID:       "48fded16-34e8-45df-993d-6c0e39ca0308",
		Name:     "test",
		ParentID: nil,
	})

	return &RamBucketRepository{
		buckets: objects,
	}
}

func (r *RamBucketRepository) Save(bucket entities.Bucket) error {
	r.buckets = append(r.buckets, bucket)
	return nil
}

func (r *RamBucketRepository) GetAll() ([]entities.Bucket, error) {
	return r.buckets, nil
}

func (r *RamBucketRepository) FindByID(ID string) (*entities.Bucket, error) {
	for _, bucket := range r.buckets {
		if bucket.ID == ID {
			return &bucket, nil
		}
	}

	return nil, errors.New("bucket not found")
}
