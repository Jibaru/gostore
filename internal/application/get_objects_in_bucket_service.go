package application

import (
	"errors"
	"github.com/jibaru/gostore/internal/domain/entities"
	"github.com/jibaru/gostore/internal/domain/repositories"
)

type GetObjectsInBucketServiceInputPort interface {
	Do(bucketID string) ([]entities.Object, error)
}

type GetObjectsInBucketService struct {
	objectRepository repositories.ObjectRepository
}

func NewGetObjectsInBucketService(
	objectRepository repositories.ObjectRepository,
) GetObjectsInBucketServiceInputPort {
	return &GetObjectsInBucketService{
		objectRepository,
	}
}

func (serv *GetObjectsInBucketService) Do(bucketID string) ([]entities.Object, error) {
	objects, err := serv.objectRepository.GetByBucketID(bucketID)
	if err != nil {
		return nil, err
	}

	if len(objects) == 0 {
		return nil, errors.New("objects not found")
	}

	return objects, nil
}
