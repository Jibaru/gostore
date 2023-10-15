package repositories

import (
	"errors"
	"github.com/jibaru/gostore/internal/domain/entities"
	"github.com/jibaru/gostore/internal/domain/repositories"
)

type RamObjectRepository struct {
	objects []entities.Object
}

func NewRamObjectRepository() repositories.ObjectRepository {
	objects := make([]entities.Object, 0)
	objects = append(objects, entities.Object{
		ID:        "4b3622bc-d5ec-4071-927c-b649611cdb18",
		Name:      "imagen.png",
		Extension: ".png",
		BucketID:  "48fded16-34e8-45df-993d-6c0e39ca0308",
	})

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
