package repositories

import (
	"encoding/json"
	"errors"
	"github.com/jibaru/gostore/internal/domain/entities"
	"os"
)

type FileBucketRepository struct {
	filePath string
	buckets  []entities.Bucket
}

func NewFileBucketRepository(filePath string) (*FileBucketRepository, error) {
	repo := &FileBucketRepository{
		filePath: filePath,
	}

	err := repo.loadFromJSONFile()
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func (r *FileBucketRepository) Save(bucket entities.Bucket) error {
	r.buckets = append(r.buckets, bucket)
	return r.saveToJSONFile()
}

func (r *FileBucketRepository) GetAll() ([]entities.Bucket, error) {
	return r.buckets, nil
}

func (r *FileBucketRepository) FindByID(ID string) (*entities.Bucket, error) {
	for _, bucket := range r.buckets {
		if bucket.ID == ID {
			return &bucket, nil
		}
	}

	return nil, errors.New("bucket not found")
}

func (r *FileBucketRepository) GetByParentID(parentID string) ([]entities.Bucket, error) {
	buckets := make([]entities.Bucket, 0)

	for _, bucket := range r.buckets {
		if bucket.ParentID != nil && *bucket.ParentID == parentID {
			buckets = append(buckets, bucket)
		}
	}

	return buckets, nil
}

func (r *FileBucketRepository) DeleteByID(ID string) error {
	buckets := make([]entities.Bucket, 0)

	for _, bucket := range r.buckets {
		if bucket.ID != ID {
			buckets = append(buckets, bucket)
		}
	}

	r.buckets = buckets

	err := r.saveToJSONFile()
	if err != nil {
		return err
	}

	return nil
}

func (r *FileBucketRepository) loadFromJSONFile() error {
	if _, err := os.Stat(r.filePath); os.IsNotExist(err) {
		r.buckets = []entities.Bucket{}
		return nil
	}

	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &r.buckets)
	if err != nil {
		return err
	}

	return nil
}

func (r *FileBucketRepository) saveToJSONFile() error {
	data, err := json.Marshal(r.buckets)
	if err != nil {
		return err
	}

	err = os.WriteFile(r.filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
