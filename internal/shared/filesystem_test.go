package shared

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeDirectoryOnServerFilesystem(t *testing.T) {
	serverFilesystem := NewServerFilesystem("../../storage")

	err := serverFilesystem.MakeDirectory(uuid.New().String())

	assert.Nil(t, err)
}

func TestMakeDirectoryWithEmptyNameOnServerFilesystemFails(t *testing.T) {
	serverFilesystem := NewServerFilesystem("../../storage")

	err := serverFilesystem.MakeDirectory("")

	assert.NotNil(t, err)
}

func TestMakeFileOnPathOnServerFilesystem(t *testing.T) {
	serverFilesystem := NewServerFilesystem("../../storage")

	file, err := serverFilesystem.MakeFileOnPath(uuid.New().String()+".png", "/")

	assert.NotNil(t, file)
	assert.Nil(t, err)
}
