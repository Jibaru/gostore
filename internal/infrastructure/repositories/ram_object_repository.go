package repositories

import (
	"errors"
	"github.com/jibaru/gostore/internal/domain/entities"
)

type RamObjectRepository struct {
	objects []entities.Object
}

func NewRamObjectRepository(objects []entities.Object) *RamObjectRepository {
	return &RamObjectRepository{objects: objects}
}

func (r *RamObjectRepository) Save(object entities.Object) error {
	r.objects = append(r.objects, object)
	return nil
}

func (r *RamObjectRepository) FindByID(ID string) (*entities.Object, error) {
	for _, object := range r.objects {
		if object.ID == ID {
			return &object, nil
		}
	}

	return nil, errors.New("object not found")
}

func (r *RamObjectRepository) GetByBucketID(bucketID string) ([]entities.Object, error) {
	objects := make([]entities.Object, 0)

	for _, object := range r.objects {
		if object.BucketID == bucketID {
			objects = append(objects, object)
		}
	}

	return objects, nil
}

func (r *RamObjectRepository) DeleteByID(objectID string) error {
	objects := make([]entities.Object, 0)
	objectDeleted := false

	for _, object := range r.objects {
		if object.ID != objectID {
			objects = append(objects, object)
		} else {
			objectDeleted = true
		}
	}

	if !objectDeleted {
		return errors.New("object not found")
	}

	r.objects = objects

	return nil
}

func (r *RamObjectRepository) Size() int {
	return len(r.objects)
}
