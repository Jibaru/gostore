package application

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jibaru/gostore/internal/domain/entities"
	"github.com/jibaru/gostore/internal/infrastructure/repositories"
	"github.com/stretchr/testify/assert"
	"testing"
)

type CallableGenerateBucketPathService struct {
	onDo func(bucketID string) (string, error)
}

func (s *CallableGenerateBucketPathService) Do(bucketID string) (string, error) {
	return s.onDo(bucketID)
}

func TestGenerateObjectPathService(t *testing.T) {
	objectID := uuid.New().String()
	bucketID := uuid.New().String()
	extension := ".png"
	objects := make([]entities.Object, 0)
	objects = append(objects, entities.Object{
		ID:        objectID,
		Name:      "object_test",
		Extension: extension,
		BucketID:  bucketID,
	})
	objectRepository := repositories.NewRamObjectRepository(objects)
	bucketPathService := CallableGenerateBucketPathService{
		onDo: func(bucketID string) (string, error) {
			return "/" + bucketID, nil
		},
	}
	service := NewGenerateObjectPathService(
		objectRepository,
		&bucketPathService,
	)
	expectedPath := "/" + bucketID + "/" + objectID + extension

	path, err := service.Do(objectID)

	assert.Nil(t, err)
	assert.Equal(t, expectedPath, path)
}

func TestGenerateObjectPathServiceWithInvalidBucketFails(t *testing.T) {
	objectID := uuid.New().String()
	bucketID := uuid.New().String()
	extension := ".png"
	objects := make([]entities.Object, 0)
	objects = append(objects, entities.Object{
		ID:        objectID,
		Name:      "object_test",
		Extension: extension,
		BucketID:  bucketID,
	})
	objectRepository := repositories.NewRamObjectRepository(objects)
	bucketPathService := CallableGenerateBucketPathService{
		onDo: func(bucketID string) (string, error) {
			return "", errors.New("an error has occurred")
		},
	}
	service := NewGenerateObjectPathService(
		objectRepository,
		&bucketPathService,
	)

	path, err := service.Do(objectID)

	assert.NotNil(t, err)
	assert.Empty(t, path)
}

func TestGenerateObjectPathServiceWithInvalidObjectFails(t *testing.T) {
	objectID := uuid.New().String()
	objects := make([]entities.Object, 0)
	objectRepository := repositories.NewRamObjectRepository(objects)
	bucketPathService := CallableGenerateBucketPathService{
		onDo: func(bucketID string) (string, error) {
			return "", nil
		},
	}
	service := NewGenerateObjectPathService(
		objectRepository,
		&bucketPathService,
	)

	path, err := service.Do(objectID)

	assert.NotNil(t, err)
	assert.Empty(t, path)
}
