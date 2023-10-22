package application

import (
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
	return serv.bucketRepository.GetAll()
}
