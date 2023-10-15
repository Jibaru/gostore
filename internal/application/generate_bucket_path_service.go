package application

import "github.com/jibaru/gostore/internal/domain/repositories"

type GenerateBucketPathService struct {
	bucketRepository repositories.BucketRepository
}

func NewGenerateBucketPathService(
	bucketRepository repositories.BucketRepository,
) *GenerateBucketPathService {
	return &GenerateBucketPathService{
		bucketRepository,
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
