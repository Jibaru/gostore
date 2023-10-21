package application

import (
	"github.com/google/uuid"
	"github.com/jibaru/gostore/internal/domain/entities"
	"github.com/jibaru/gostore/internal/domain/repositories"
	"github.com/jibaru/gostore/internal/shared"
)

type CreateBucketService struct {
	bucketRepository          repositories.BucketRepository
	filesystem                shared.Filesystem
	generateBucketPathService GenerateBucketPathServiceInputPort
}

func NewCreateBucketService(
	bucketRepository repositories.BucketRepository,
	filesystem shared.Filesystem,
	generateBucketPathService GenerateBucketPathServiceInputPort,
) *CreateBucketService {
	return &CreateBucketService{
		bucketRepository,
		filesystem,
		generateBucketPathService,
	}
}

func (serv *CreateBucketService) Do(
	name string,
	parentID *string,
) (*entities.Bucket, error) {
	bucketID := uuid.New().String()

	if parentID != nil {
		parentPath, err := serv.generateBucketPathService.Do(*parentID)
		if err != nil {
			return nil, err
		}

		err = serv.filesystem.MakeDirectoryOnPath(bucketID, parentPath)
		if err != nil {
			return nil, err
		}
	} else {
		err := serv.filesystem.MakeDirectory(bucketID)
		if err != nil {
			return nil, err
		}
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
