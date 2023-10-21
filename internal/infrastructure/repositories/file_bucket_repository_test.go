package repositories

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/jibaru/gostore/internal/domain/entities"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

const testingStorageBucketPath string = "../../../storage/buckets_test.json"

func setUpFileBucketRepositoryTest() {
	err := os.Remove(testingStorageBucketPath)
	if err != nil {
		log.Println(err.Error())
	}
}

func mockBuckets(buckets []entities.Bucket) error {
	data, err := json.Marshal(buckets)
	if err != nil {
		return err
	}

	err = os.WriteFile(testingStorageBucketPath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func TestCreateNewFileBucketPathRepository(t *testing.T) {
	setUpFileBucketRepositoryTest()
	repository, err := NewFileBucketRepository(testingStorageBucketPath)

	assert.Nil(t, err)
	assert.NotNil(t, repository)
}

func TestFileBucketRepositorySave(t *testing.T) {
	setUpFileBucketRepositoryTest()
	repository, err := NewFileBucketRepository(testingStorageBucketPath)
	assert.Nil(t, err)
	assert.NotNil(t, repository)

	err = repository.Save(entities.Bucket{
		ID:       uuid.New().String(),
		Name:     "test_bucket",
		ParentID: nil,
	})

	assert.Nil(t, err)
}

func TestFileBucketRepositoryGetAll(t *testing.T) {
	firstBucketID := uuid.New().String()
	mockedBuckets := []entities.Bucket{
		{ID: firstBucketID, Name: "test_bucket_1", ParentID: nil},
		{ID: uuid.New().String(), Name: "test_bucket_2", ParentID: &firstBucketID},
	}

	err := mockBuckets(mockedBuckets)
	assert.Nil(t, err)

	repository, err := NewFileBucketRepository(testingStorageBucketPath)
	assert.Nil(t, err)
	assert.NotNil(t, repository)

	buckets, err := repository.GetAll()

	assert.Nil(t, err)
	assert.Len(t, buckets, len(mockedBuckets))
	assert.Equal(t, mockedBuckets, buckets)
}

func TestFileBucketRepositoryFindByID(t *testing.T) {
	firstBucketID := uuid.New().String()
	mockedBuckets := []entities.Bucket{
		{ID: firstBucketID, Name: "test_bucket_1", ParentID: nil},
		{ID: uuid.New().String(), Name: "test_bucket_2", ParentID: &firstBucketID},
	}

	err := mockBuckets(mockedBuckets)
	assert.Nil(t, err)

	repository, err := NewFileBucketRepository(testingStorageBucketPath)
	assert.Nil(t, err)
	assert.NotNil(t, repository)

	bucket, err := repository.FindByID(firstBucketID)

	assert.Nil(t, err)
	assert.NotNil(t, bucket)
	assert.Equal(t, firstBucketID, bucket.ID)
	assert.Equal(t, mockedBuckets[0], *bucket)
}
