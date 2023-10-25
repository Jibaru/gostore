package application

import (
	"github.com/jibaru/gostore/internal/domain/repositories"
	"github.com/jibaru/gostore/internal/shared"
)

type DeleteBucketServiceInputPort interface {
	Do(bucketID string) error
}

type DeleteBucketService struct {
	bucketRepository          repositories.BucketRepository
	generateBucketPathService GenerateBucketPathServiceInputPort
	filesystem                shared.Filesystem
}

func NewDeleteBucketService(
	bucketRepository repositories.BucketRepository,
	generateBucketPathService GenerateBucketPathServiceInputPort,
	filesystem shared.Filesystem,
) DeleteBucketServiceInputPort {
	return &DeleteBucketService{
		bucketRepository,
		generateBucketPathService,
		filesystem,
	}
}

func (serv *DeleteBucketService) Do(bucketID string) error {
	path, err := serv.generateBucketPathService.Do(bucketID)
	if err != nil {
		return err
	}

	err = serv.filesystem.DeleteDirectoryOnPath(path)
	if err != nil {
		return err
	}

	err = serv.bucketRepository.DeleteByID(bucketID)
	if err != nil {
		return err
	}

	return nil
}
