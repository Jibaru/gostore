package repositories

import (
	"encoding/json"
	"errors"
	"github.com/jibaru/gostore/internal/domain/entities"
	"os"
)

type FileObjectRepository struct {
	filePath string
	objects  []entities.Object
}

func NewFileObjectRepository(filePath string) (*FileObjectRepository, error) {
	repo := &FileObjectRepository{
		filePath: filePath,
	}

	err := repo.loadFromJSONFile()
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func (r *FileObjectRepository) Save(object entities.Object) error {
	r.objects = append(r.objects, object)
	return r.saveToJSONFile()
}

func (r *FileObjectRepository) FindByID(ID string) (*entities.Object, error) {
	for _, object := range r.objects {
		if object.ID == ID {
			return &object, nil
		}
	}

	return nil, errors.New("object not found")
}

func (r *FileObjectRepository) GetByBucketID(bucketID string) ([]entities.Object, error) {
	objects := make([]entities.Object, 0)

	for _, object := range r.objects {
		if object.BucketID == bucketID {
			objects = append(objects, object)
		}
	}

	return objects, nil
}

func (r *FileObjectRepository) loadFromJSONFile() error {
	if _, err := os.Stat(r.filePath); os.IsNotExist(err) {
		r.objects = []entities.Object{}
		return nil
	}

	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &r.objects)
	if err != nil {
		return err
	}

	return nil
}

func (r *FileObjectRepository) saveToJSONFile() error {
	data, err := json.Marshal(r.objects)
	if err != nil {
		return err
	}

	err = os.WriteFile(r.filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
