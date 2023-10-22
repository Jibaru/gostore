package application

import (
	"errors"
	"github.com/jibaru/gostore/internal/domain/entities"
	"github.com/jibaru/gostore/internal/domain/repositories"
)

type GetBucketsInBucketServiceInputPort interface {
	Do(bucketID string) ([]entities.Bucket, error)
}

type GetBucketsInBucketService struct {
	bucketsRepository repositories.BucketRepository
}

func NewGetBucketsInBucketService(
	bucketsRepository repositories.BucketRepository,
) GetBucketsInBucketServiceInputPort {
	return &GetBucketsInBucketService{
		bucketsRepository,
	}
}

func (serv *GetBucketsInBucketService) Do(bucketID string) ([]entities.Bucket, error) {
	buckets, err := serv.bucketsRepository.GetByParentID(bucketID)
	if err != nil {
		return nil, err
	}

	if len(buckets) == 0 {
		return nil, errors.New("buckets not found")
	}

	return buckets, nil
}
