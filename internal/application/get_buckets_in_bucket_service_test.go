package application

import (
	"github.com/google/uuid"
	"github.com/jibaru/gostore/internal/domain/entities"
	"github.com/jibaru/gostore/internal/infrastructure/repositories"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBucketsInBucketService(t *testing.T) {
	parentID := uuid.New().String()
	anotherParentID := uuid.New().String()
	buckets := []entities.Bucket{
		{ID: uuid.New().String(), Name: "test_bucket_1", ParentID: &parentID},
		{ID: uuid.New().String(), Name: "test_bucket_2", ParentID: nil},
		{ID: uuid.New().String(), Name: "test_bucket_3", ParentID: &parentID},
		{ID: uuid.New().String(), Name: "test_bucket_4", ParentID: &anotherParentID},
	}

	bucketRepository := repositories.NewRamBucketRepository(buckets)
	expectedBuckets := []entities.Bucket{
		buckets[0],
		buckets[2],
	}
	service := NewGetBucketsInBucketService(bucketRepository)

	buckets, err := service.Do(parentID)

	assert.Nil(t, err)
	assert.NotNil(t, buckets)
	assert.NotEmpty(t, buckets)
	assert.Len(t, buckets, 2)
	assert.Equal(t, expectedBuckets, buckets)
}

func TestGetBucketsInBucketServiceFails(t *testing.T) {
	parentID := uuid.New().String()
	anotherParentID := uuid.New().String()

	bucketRepository := repositories.NewRamBucketRepository([]entities.Bucket{
		{ID: uuid.New().String(), Name: "test_bucket_1", ParentID: &parentID},
		{ID: uuid.New().String(), Name: "test_bucket_2", ParentID: nil},
		{ID: uuid.New().String(), Name: "test_bucket_3", ParentID: &parentID},
	})
	service := NewGetBucketsInBucketService(bucketRepository)

	buckets, err := service.Do(anotherParentID)

	assert.NotNil(t, err)
	assert.Nil(t, buckets)
}
