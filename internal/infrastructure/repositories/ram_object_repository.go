package repositories

import (
	"github.com/jibaru/gostore/internal/domain/entities"
	"github.com/jibaru/gostore/internal/domain/repositories"
)

type RamObjectRepository struct {
	objects []entities.Object
}

func NewRamObjectRepository() repositories.ObjectRepository {
	return &RamObjectRepository{objects: make([]entities.Object, 0)}
}

func (r *RamObjectRepository) Save(object entities.Object) error {
	r.objects = append(r.objects, object)
	return nil
}
