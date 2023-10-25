package repositories

import (
	"errors"
	"github.com/jibaru/gostore/internal/domain/entities"
	"github.com/jibaru/gostore/internal/domain/repositories"
)

type RamBucketRepository struct {
	buckets []entities.Bucket
}

func NewEmptyRamBucketRepository() repositories.BucketRepository {
	return &RamBucketRepository{
		make([]entities.Bucket, 0),
	}
}

func NewRamBucketRepository(buckets []entities.Bucket) repositories.BucketRepository {
	return &RamBucketRepository{
		buckets,
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

func (r *RamBucketRepository) GetByParentID(parentID string) ([]entities.Bucket, error) {
	buckets := make([]entities.Bucket, 0)

	for _, bucket := range r.buckets {
		if bucket.ParentID != nil && *bucket.ParentID == parentID {
			buckets = append(buckets, bucket)
		}
	}

	return buckets, nil
}

func (r *RamBucketRepository) DeleteByID(ID string) error {
	buckets := make([]entities.Bucket, 0)

	for _, bucket := range r.buckets {
		if bucket.ID != ID {
			if bucket.ParentID != nil && *bucket.ParentID != ID {
				buckets = append(buckets, bucket)
			}
		}
	}

	r.buckets = buckets

	return nil
}
