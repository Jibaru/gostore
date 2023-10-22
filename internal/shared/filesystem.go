package shared

import (
	"errors"
	"os"
	"strings"
)

type Filesystem interface {
	MakeDirectory(name string) error
	MakeDirectoryOnPath(name string, relativePath string) error
	MakeFileOnPath(name string, relativePath string) (*os.File, error)
	DeleteFileOnPath(path string) error
}

type ServerFilesystem struct {
	rootPath string
}

type DummyFilesystem struct {
}

func NewServerFilesystem(rootPath string) Filesystem {
	return &ServerFilesystem{
		rootPath,
	}
}

func NewDummyFilesystem() Filesystem {
	return &DummyFilesystem{}
}

func (s *ServerFilesystem) MakeDirectory(name string) error {
	if len(strings.TrimSpace(name)) == 0 {
		return errors.New("name should be not empty")
	}

	return os.Mkdir(s.rootPath+"/"+name, os.ModePerm)
}

func (s *ServerFilesystem) MakeDirectoryOnPath(name string, relativePath string) error {
	return os.Mkdir(s.rootPath+relativePath+"/"+name, os.ModePerm)
}

func (s *ServerFilesystem) MakeFileOnPath(name string, relativePath string) (*os.File, error) {
	return os.Create(s.rootPath + relativePath + "/" + name)
}

func (s *ServerFilesystem) DeleteFileOnPath(relativePath string) error {
	return os.Remove(s.rootPath + relativePath)
}

func (s *DummyFilesystem) MakeDirectory(name string) error {
	return nil
}

func (s *DummyFilesystem) MakeDirectoryOnPath(name string, relativePath string) error {
	return nil
}

func (s *DummyFilesystem) MakeFileOnPath(name string, relativePath string) (*os.File, error) {
	return &os.File{}, nil
}

func (s *DummyFilesystem) DeleteFileOnPath(relativePath string) error {
	return nil
}
