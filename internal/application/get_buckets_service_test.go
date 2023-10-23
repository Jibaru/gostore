package application

import (
	"github.com/google/uuid"
	"github.com/jibaru/gostore/internal/domain/entities"
	"github.com/jibaru/gostore/internal/infrastructure/repositories"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBucketsService(t *testing.T) {
	firstBucketID := uuid.NewString()
	expectedBuckets := []entities.Bucket{
		{ID: firstBucketID, Name: "test_bucket_1", ParentID: nil},
		{ID: uuid.NewString(), Name: "test_bucket_2", ParentID: &firstBucketID},
	}
	bucketRepository := repositories.NewRamBucketRepository(expectedBuckets)
	service := NewGetBucketsService(bucketRepository)

	buckets, err := service.Do()

	assert.Nil(t, err)
	assert.NotNil(t, buckets)
	assert.Equal(t, expectedBuckets, buckets)
}

func TestGetBucketsServiceFails(t *testing.T) {
	bucketRepository := repositories.NewEmptyRamBucketRepository()
	service := NewGetBucketsService(bucketRepository)

	buckets, err := service.Do()

	assert.NotNil(t, err)
	assert.Nil(t, buckets)
}
