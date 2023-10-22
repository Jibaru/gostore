package application

import (
	"github.com/jibaru/gostore/internal/domain/repositories"
	"github.com/jibaru/gostore/internal/shared"
)

type DeleteObjectServiceInputPort interface {
	Do(objectID string) error
}

type DeleteObjectService struct {
	objectRepository          repositories.ObjectRepository
	generateObjectPathService GenerateObjectPathServiceInputPort
	filesystem                shared.Filesystem
}

func NewDeleteObjectService(
	objectRepository repositories.ObjectRepository,
	generateObjectPathService GenerateObjectPathServiceInputPort,
	filesystem shared.Filesystem,
) DeleteObjectServiceInputPort {
	return &DeleteObjectService{
		objectRepository,
		generateObjectPathService,
		filesystem,
	}
}

func (serv *DeleteObjectService) Do(objectID string) error {
	path, err := serv.generateObjectPathService.Do(objectID)
	if err != nil {
		return err
	}

	err = serv.objectRepository.DeleteByID(objectID)
	if err != nil {
		return err
	}

	err = serv.filesystem.DeleteFileOnPath(path)
	if err != nil {
		return err
	}

	return nil
}
