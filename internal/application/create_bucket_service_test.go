package application

import (
	"github.com/google/uuid"
	"github.com/jibaru/gostore/internal/domain/entities"
	"github.com/jibaru/gostore/internal/infrastructure/repositories"
	"github.com/jibaru/gostore/internal/shared"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateBucketWithEmptyParent(t *testing.T) {
	ramBucketRepository := repositories.NewEmptyRamBucketRepository()
	dummyFilesystem := shared.NewDummyFilesystem()
	createBucketService := NewCreateBucketService(ramBucketRepository, dummyFilesystem)

	bucket, err := createBucketService.Do("test_bucket", nil)

	assert.Nil(t, err)
	assert.NotNil(t, bucket)
}

func TestCreateBucketWithEmptyNameFails(t *testing.T) {
	ramBucketRepository := repositories.NewEmptyRamBucketRepository()
	dummyFilesystem := shared.NewDummyFilesystem()
	createBucketService := NewCreateBucketService(ramBucketRepository, dummyFilesystem)

	bucket, err := createBucketService.Do("", nil)

	assert.NotNil(t, err)
	assert.Nil(t, bucket)
}

func TestCreateBucketWithNotExistsParentFails(t *testing.T) {
	ramBucketRepository := repositories.NewEmptyRamBucketRepository()
	dummyFilesystem := shared.NewDummyFilesystem()
	createBucketService := NewCreateBucketService(ramBucketRepository, dummyFilesystem)
	parentID := uuid.New().String()

	bucket, err := createBucketService.Do("test_bucket", &parentID)

	assert.Nil(t, err)
	assert.NotNil(t, bucket)
}

func TestCreateBucketWithExistsParent(t *testing.T) {
	parentID := uuid.New().String()
	buckets := make([]entities.Bucket, 0)
	buckets = append(buckets, entities.Bucket{
		ID:       parentID,
		Name:     "test_bucket",
		ParentID: nil,
	})
	ramBucketRepository := repositories.NewRamBucketRepository(buckets)
	dummyFilesystem := shared.NewDummyFilesystem()
	createBucketService := NewCreateBucketService(ramBucketRepository, dummyFilesystem)

	bucket, err := createBucketService.Do("test_bucket", &parentID)

	assert.Nil(t, err)
	assert.NotNil(t, bucket)
}
