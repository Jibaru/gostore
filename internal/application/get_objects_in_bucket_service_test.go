package application

import (
	"github.com/google/uuid"
	"github.com/jibaru/gostore/internal/domain/entities"
	"github.com/jibaru/gostore/internal/infrastructure/repositories"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetObjectsInBucketService(t *testing.T) {
	bucketID := uuid.New().String()
	objects := []entities.Object{
		{ID: uuid.New().String(), Name: "test_object_1", Extension: ".png", BucketID: bucketID},
		{ID: uuid.New().String(), Name: "test_object_2", Extension: ".png", BucketID: uuid.New().String()},
		{ID: uuid.New().String(), Name: "test_object_3", Extension: ".png", BucketID: bucketID},
		{ID: uuid.New().String(), Name: "test_object_4", Extension: ".png", BucketID: uuid.New().String()},
	}

	objectRepository := repositories.NewRamObjectRepository(objects)
	expectedObjects := []entities.Object{
		objects[0],
		objects[2],
	}
	service := NewGetObjectsInBucketService(objectRepository)

	objects, err := service.Do(bucketID)

	assert.Nil(t, err)
	assert.NotNil(t, objects)
	assert.NotEmpty(t, objects)
	assert.Len(t, objects, 2)
	assert.Equal(t, expectedObjects, objects)
}

func TestGetObjectsInBucketServiceFails(t *testing.T) {
	bucketID := uuid.New().String()
	objects := []entities.Object{
		{ID: uuid.New().String(), Name: "test_object_1", Extension: ".png", BucketID: uuid.New().String()},
		{ID: uuid.New().String(), Name: "test_object_2", Extension: ".png", BucketID: uuid.New().String()},
	}

	objectRepository := repositories.NewRamObjectRepository(objects)
	service := NewGetObjectsInBucketService(objectRepository)

	objects, err := service.Do(bucketID)

	assert.NotNil(t, err)
	assert.Nil(t, objects)
}
