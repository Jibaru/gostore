package application

import (
	"github.com/google/uuid"
	"github.com/jibaru/gostore/internal/domain/entities"
	"github.com/jibaru/gostore/internal/infrastructure/repositories"
	"github.com/jibaru/gostore/internal/shared"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteObjectService(t *testing.T) {
	objectID := uuid.NewString()
	bucketID := uuid.NewString()
	extension := ".png"
	generateObjectPathServ := NewCallableGenerateObjectPathServiceForValidObject(
		bucketID,
		extension,
	)
	objectRepository := repositories.NewRamObjectRepository([]entities.Object{
		{ID: objectID, Name: "object_1", Extension: extension, BucketID: bucketID},
		{ID: uuid.NewString(), Name: "object_2", Extension: extension, BucketID: bucketID},
	})
	filesystem := shared.NewDummyFilesystem()
	service := NewDeleteObjectService(
		objectRepository,
		generateObjectPathServ,
		filesystem,
	)

	err := service.Do(objectID)

	assert.Nil(t, err)
	assert.Equal(t, 1, objectRepository.Size())
}

func TestDeleteObjectServiceFails(t *testing.T) {
	objectID := uuid.NewString()
	bucketID := uuid.NewString()
	extension := ".png"
	generateObjectPathServ := NewCallableGenerateObjectPathServiceForValidObject(
		bucketID,
		extension,
	)
	objectRepository := repositories.NewRamObjectRepository([]entities.Object{})
	filesystem := shared.NewDummyFilesystem()
	service := NewDeleteObjectService(
		objectRepository,
		generateObjectPathServ,
		filesystem,
	)

	err := service.Do(objectID)

	assert.NotNil(t, err)
	assert.Equal(t, 0, objectRepository.Size())
}
