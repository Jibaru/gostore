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

func TestServerFilesystem_DeleteFileOnPath(t *testing.T) {
	serverFilesystem := NewServerFilesystem("../../storage")
	fileName := uuid.New().String() + ".png"
	file, err := serverFilesystem.MakeFileOnPath(fileName, "/")
	assert.Nil(t, err)
	assert.NotNil(t, file)

	err = serverFilesystem.DeleteFileOnPath("/" + fileName)

	assert.Nil(t, err)
}

func TestServerFilesystem_DeleteDirectoryOnPath(t *testing.T) {
	serverFilesystem := NewServerFilesystem("../../storage")
	dirName := uuid.New().String()
	err := serverFilesystem.MakeDirectoryOnPath(dirName, "/")
	assert.Nil(t, err)

	err = serverFilesystem.DeleteDirectoryOnPath("/" + dirName)

	assert.Nil(t, err)
}
