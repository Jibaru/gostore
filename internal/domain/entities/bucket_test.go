package entities

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBucketWithEmptyParent(t *testing.T) {
	bucketID := uuid.New().String()
	name := "test_bucket"

	bucket, err := NewBucket(
		bucketID,
		name,
		nil,
	)

	assert.Nil(t, err)
	assert.Equal(t, bucketID, bucket.ID)
	assert.Equal(t, name, bucket.Name)
	assert.Nil(t, bucket.ParentID)
}

func TestNewBucketWithParent(t *testing.T) {
	bucketID := uuid.New().String()
	name := "test_bucket"
	parentID := uuid.New().String()

	bucket, err := NewBucket(
		bucketID,
		name,
		&parentID,
	)

	assert.Nil(t, err)
	assert.Equal(t, bucketID, bucket.ID)
	assert.Equal(t, name, bucket.Name)
	assert.Equal(t, parentID, *bucket.ParentID)
}

func TestNewBucketWithEmptyIDFails(t *testing.T) {
	bucketID := ""
	name := "test_bucket"

	bucket, err := NewBucket(
		bucketID,
		name,
		nil,
	)

	assert.NotNil(t, err)
	assert.Nil(t, bucket)
}

func TestNewBucketWithEmptyNameFails(t *testing.T) {
	bucketID := uuid.New().String()
	name := ""

	bucket, err := NewBucket(
		bucketID,
		name,
		nil,
	)

	assert.NotNil(t, err)
	assert.Nil(t, bucket)
}

func TestNewBucketWithEmptyParentIDFails(t *testing.T) {
	bucketID := uuid.New().String()
	name := "test_bucket"
	parentID := ""

	bucket, err := NewBucket(
		bucketID,
		name,
		&parentID,
	)

	assert.NotNil(t, err)
	assert.Nil(t, bucket)
}

func TestNewBucketWithInvalidIDFails(t *testing.T) {
	bucketID := "random_id"
	name := "test_bucket"

	bucket, err := NewBucket(
		bucketID,
		name,
		nil,
	)

	assert.NotNil(t, err)
	assert.Nil(t, bucket)
}

func TestNewBucketWithInvalidParentIDFails(t *testing.T) {
	bucketID := uuid.New().String()
	name := "test_bucket"
	parentID := "random_id"

	bucket, err := NewBucket(
		bucketID,
		name,
		&parentID,
	)

	assert.NotNil(t, err)
	assert.Nil(t, bucket)
}
