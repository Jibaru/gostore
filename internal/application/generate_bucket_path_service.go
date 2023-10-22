package application

import "github.com/jibaru/gostore/internal/domain/repositories"

type GenerateBucketPathServiceInputPort interface {
	Do(bucketID string) (string, error)
}

type GenerateBucketPathService struct {
	bucketRepository repositories.BucketRepository
}

type CallableGenerateBucketPathService struct {
	onDo func(bucketID string) (string, error)
}

func NewGenerateBucketPathService(
	bucketRepository repositories.BucketRepository,
) GenerateBucketPathServiceInputPort {
	return &GenerateBucketPathService{
		bucketRepository,
	}
}

func NewCallableGenerateBucketPathServiceForRootBucket() GenerateBucketPathServiceInputPort {
	return &CallableGenerateBucketPathService{
		func(bucketID string) (string, error) {
			return "/" + bucketID, nil
		},
	}
}

func (serv *GenerateBucketPathService) Do(bucketID string) (string, error) {
	bucket, err := serv.bucketRepository.FindByID(bucketID)
	if err != nil {
		return "", err
	}

	if bucket.InRoot() {
		return "/" + bucket.ID, nil
	}

	path, err := serv.Do(*bucket.ParentID)
	if err != nil {
		return "", err
	}

	return path + "/" + bucketID, nil
}

func (s *CallableGenerateBucketPathService) Do(bucketID string) (string, error) {
	return s.onDo(bucketID)
}
