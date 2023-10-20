package application

import (
	"github.com/google/uuid"
	"github.com/jibaru/gostore/internal/domain/entities"
	"github.com/jibaru/gostore/internal/infrastructure/repositories"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateBucketPathServiceWithBucketWithEmptyParent(t *testing.T) {
	bucketID := uuid.New().String()
	buckets := make([]entities.Bucket, 0)
	buckets = append(buckets, entities.Bucket{
		ID:       bucketID,
		Name:     "bucket_test",
		ParentID: nil,
	})
	bucketRepository := repositories.NewRamBucketRepository(buckets)
	service := NewGenerateBucketPathService(
		bucketRepository,
	)
	expectedPath := "/" + bucketID

	path, err := service.Do(bucketID)

	assert.Nil(t, err)
	assert.Equal(t, expectedPath, path)
}

func TestGenerateBucketPathServiceWithBucketWithNestedParent(t *testing.T) {
	bucketID := uuid.New().String()
	parentID := uuid.New().String()
	rootID := uuid.New().String()
	buckets := make([]entities.Bucket, 0)
	buckets = append(buckets, entities.Bucket{
		ID:       rootID,
		Name:     "bucket_test",
		ParentID: nil,
	})
	buckets = append(buckets, entities.Bucket{
		ID:       parentID,
		Name:     "parent_bucket_test",
		ParentID: &rootID,
	})
	buckets = append(buckets, entities.Bucket{
		ID:       bucketID,
		Name:     "bucket_test",
		ParentID: &parentID,
	})
	bucketRepository := repositories.NewRamBucketRepository(buckets)
	service := NewGenerateBucketPathService(
		bucketRepository,
	)
	expectedPath := "/" + rootID + "/" + parentID + "/" + bucketID

	path, err := service.Do(bucketID)

	assert.Nil(t, err)
	assert.Equal(t, expectedPath, path)
}

func TestGenerateBucketPathServiceWithInvalidBucketIDFails(t *testing.T) {
	bucketRepository := repositories.NewRamBucketRepository(make([]entities.Bucket, 0))
	service := NewGenerateBucketPathService(
		bucketRepository,
	)
	bucketID := uuid.New().String()

	path, err := service.Do(bucketID)

	assert.NotNil(t, err)
	assert.Empty(t, path)
}
