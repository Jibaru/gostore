package application

import (
	"errors"
	"github.com/jibaru/gostore/internal/domain/entities"
	"github.com/jibaru/gostore/internal/domain/repositories"
)

type GetBucketServiceInputPort interface {
	Do() ([]entities.Bucket, error)
}

type GetBucketsService struct {
	bucketRepository repositories.BucketRepository
}

func NewGetBucketsService(
	bucketRepository repositories.BucketRepository,
) GetBucketServiceInputPort {
	return &GetBucketsService{
		bucketRepository,
	}
}

func (serv *GetBucketsService) Do() ([]entities.Bucket, error) {
	buckets, err := serv.bucketRepository.GetAll()
	if err != nil {
		return nil, err
	}

	if len(buckets) == 0 {
		return nil, errors.New("buckets not found")
	}

	return buckets, nil
}
