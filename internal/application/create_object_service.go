package application

import (
	"github.com/google/uuid"
	"github.com/jibaru/gostore/internal/domain/entities"
	"github.com/jibaru/gostore/internal/domain/repositories"
	"github.com/jibaru/gostore/internal/shared"
	"io"
	"mime/multipart"
	"path/filepath"
)

type CreateObjectService struct {
	bucketRepository repositories.BucketRepository
	objectRepository repositories.ObjectRepository
	filesystem       shared.Filesystem
}

func NewCreateObjectService(
	bucketRepository repositories.BucketRepository,
	objectRepository repositories.ObjectRepository,
	filesystem shared.Filesystem,
) *CreateObjectService {
	return &CreateObjectService{
		bucketRepository,
		objectRepository,
		filesystem,
	}
}

func (serv *CreateObjectService) Do(
	file *multipart.FileHeader,
	bucketID string,
) (*entities.Object, error) {
	objectID := uuid.New().String()
	bucket, err := serv.bucketRepository.FindByID(bucketID)

	src, err := file.Open()
	defer src.Close()
	if err != nil {
		return nil, err
	}

	extension := filepath.Ext(file.Filename)

	// dst, err := os.Create("./storage/" + bucket.ID + "/" + objectID + extension)
	dst, err := serv.filesystem.MakeFileOnPath(objectID+extension, "/"+bucket.ID)
	defer dst.Close()
	if err != nil {
		return nil, err
	}

	if _, err = io.Copy(dst, src); err != nil {
		return nil, err
	}

	object, err := entities.NewObject(objectID, file.Filename, extension, bucket.ID)
	if err != nil {
		return nil, err
	}

	err = serv.objectRepository.Save(*object)
	if err != nil {
		return nil, err
	}

	return object, nil
}
