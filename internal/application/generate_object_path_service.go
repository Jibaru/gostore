package application

import "github.com/jibaru/gostore/internal/domain/repositories"

type GenerateObjectPathServiceInputPort interface {
	Do(objectID string) (string, error)
}

type GenerateObjectPathService struct {
	objectRepository          repositories.ObjectRepository
	generateBucketPathService GenerateBucketPathServiceInputPort
}

type CallableGenerateObjectPathService struct {
	onDo func(objectID string) (string, error)
}

func NewGenerateObjectPathService(
	objectRepository repositories.ObjectRepository,
	generateBucketPathService GenerateBucketPathServiceInputPort,
) GenerateObjectPathServiceInputPort {
	return &GenerateObjectPathService{
		objectRepository,
		generateBucketPathService,
	}
}

func NewCallableGenerateObjectPathServiceForValidObject(
	bucketID string,
	extension string,
) GenerateBucketPathServiceInputPort {
	return &CallableGenerateBucketPathService{
		func(objectID string) (string, error) {
			return "/" + bucketID + "/" + objectID + extension, nil
		},
	}
}

func (serv *GenerateObjectPathService) Do(objectID string) (string, error) {
	object, err := serv.objectRepository.FindByID(objectID)
	if err != nil {
		return "", err
	}

	path, err := serv.generateBucketPathService.Do(object.BucketID)
	if err != nil {
		return "", err
	}

	return path + "/" + objectID + object.Extension, nil
}
