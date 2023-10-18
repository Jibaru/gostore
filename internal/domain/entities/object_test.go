package entities

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewObject(t *testing.T) {
	objectID := uuid.New().String()
	name := "object_test"
	extension := ".png"
	bucketID := uuid.New().String()

	object, err := NewObject(
		objectID,
		name,
		extension,
		bucketID,
	)

	assert.Nil(t, err)
	assert.NotNil(t, object)
	assert.Equal(t, objectID, object.ID)
	assert.Equal(t, name, object.Name)
	assert.Equal(t, extension, object.Extension)
	assert.Equal(t, bucketID, object.BucketID)
}

func TestNewObjectWithEmptyIDFails(t *testing.T) {
	objectID := ""
	name := "object_test"
	extension := ".png"
	bucketID := uuid.New().String()

	object, err := NewObject(
		objectID,
		name,
		extension,
		bucketID,
	)

	assert.NotNil(t, err)
	assert.Nil(t, object)
}

func TestNewObjectWithEmptyNameFails(t *testing.T) {
	objectID := uuid.New().String()
	name := ""
	extension := ".png"
	bucketID := uuid.New().String()

	object, err := NewObject(
		objectID,
		name,
		extension,
		bucketID,
	)

	assert.NotNil(t, err)
	assert.Nil(t, object)
}

func TestNewObjectWithEmptyExtensionFails(t *testing.T) {
	objectID := uuid.New().String()
	name := "object_test"
	extension := ""
	bucketID := uuid.New().String()

	object, err := NewObject(
		objectID,
		name,
		extension,
		bucketID,
	)

	assert.NotNil(t, err)
	assert.Nil(t, object)
}

func TestNewObjectWithInvalidIDFails(t *testing.T) {
	objectID := "random_id"
	name := "object_test"
	extension := ".png"
	bucketID := uuid.New().String()

	object, err := NewObject(
		objectID,
		name,
		extension,
		bucketID,
	)

	assert.NotNil(t, err)
	assert.Nil(t, object)
}

func TestNewObjectWithInvalidBucketIDFails(t *testing.T) {
	objectID := uuid.New().String()
	name := "object_test"
	extension := ".png"
	bucketID := "random_id"

	object, err := NewObject(
		objectID,
		name,
		extension,
		bucketID,
	)

	assert.NotNil(t, err)
	assert.Nil(t, object)
}
