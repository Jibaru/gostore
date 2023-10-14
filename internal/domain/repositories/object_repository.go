package repositories

import (
	"github.com/jibaru/gostore/internal/domain/entities"
)

type ObjectRepository interface {
	Save(object entities.Object) error
}
