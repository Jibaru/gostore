package application

import (
	"github.com/google/uuid"
	"github.com/jibaru/gostore/internal/domain/entities"
	"github.com/jibaru/gostore/internal/domain/repositories"
	"os"
)

type CreateBucketService struct {
	bucketRepository repositories.BucketRepository
}

func NewCreateBucketService(
	bucketRepository repositories.BucketRepository,
) *CreateBucketService {
	return &CreateBucketService{
		bucketRepository,
	}
}

func (serv *CreateBucketService) Do(
	name string,
	parentID *string,
) (*entities.Bucket, error) {
	bucketID := uuid.New().String()

	err := os.Mkdir("./storage/"+bucketID, os.ModePerm)
	if err != nil {
		return nil, err
	}

	bucket, err := entities.NewBucket(bucketID, name, parentID)
	if err != nil {
		return nil, err
	}

	err = serv.bucketRepository.Save(*bucket)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}
