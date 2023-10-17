package application

import (
	"github.com/google/uuid"
	"github.com/jibaru/gostore/internal/domain/entities"
	"github.com/jibaru/gostore/internal/domain/repositories"
	"github.com/jibaru/gostore/internal/shared"
)

type CreateBucketService struct {
	bucketRepository repositories.BucketRepository
	filesystem       shared.Filesystem
}

func NewCreateBucketService(
	bucketRepository repositories.BucketRepository,
	filesystem shared.Filesystem,
) *CreateBucketService {
	return &CreateBucketService{
		bucketRepository,
		filesystem,
	}
}

func (serv *CreateBucketService) Do(
	name string,
	parentID *string,
) (*entities.Bucket, error) {
	bucketID := uuid.New().String()

	//err := os.Mkdir(serv.storageRootFolder+"/"+bucketID, os.ModePerm)
	err := serv.filesystem.MakeDirectory(bucketID)
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
