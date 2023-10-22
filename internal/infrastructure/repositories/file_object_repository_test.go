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

const testingStorageObjectPath string = "../../../storage/objects_test.json"

func setUpFileObjectRepositoryTest() {
	err := os.Remove(testingStorageObjectPath)
	if err != nil {
		log.Println(err.Error())
	}
}

func mockObjects(objects []entities.Object) error {
	data, err := json.Marshal(objects)
	if err != nil {
		return err
	}

	err = os.WriteFile(testingStorageObjectPath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func TestCreateNewFileObjectRepository(t *testing.T) {
	setUpFileObjectRepositoryTest()
	repository, err := NewFileObjectRepository(testingStorageObjectPath)

	assert.Nil(t, err)
	assert.NotNil(t, repository)
}

func TestFileObjectRepositorySave(t *testing.T) {
	setUpFileObjectRepositoryTest()
	repository, err := NewFileObjectRepository(testingStorageObjectPath)
	assert.Nil(t, err)
	assert.NotNil(t, repository)

	err = repository.Save(entities.Object{
		ID:        uuid.New().String(),
		Name:      "test_object",
		Extension: ".png",
		BucketID:  uuid.New().String(),
	})

	assert.Nil(t, err)
}

func TestFileObjectRepositoryFindByID(t *testing.T) {
	mockedObjects := []entities.Object{
		{ID: uuid.New().String(), Name: "test_object_1", Extension: ".png", BucketID: uuid.New().String()},
		{ID: uuid.New().String(), Name: "test_object_2", Extension: ".png", BucketID: uuid.New().String()},
	}

	err := mockObjects(mockedObjects)
	assert.Nil(t, err)

	repository, err := NewFileObjectRepository(testingStorageObjectPath)
	assert.Nil(t, err)
	assert.NotNil(t, repository)

	object, err := repository.FindByID(mockedObjects[0].ID)

	assert.Nil(t, err)
	assert.NotNil(t, object)
	assert.Equal(t, mockedObjects[0].ID, object.ID)
	assert.Equal(t, mockedObjects[0], *object)
}
