package storage

import (
	"github.com/nukhvu/bot-bek/internal/file"
)

type Storage struct{}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Upload(fileName string, blob []byte) (*file.File, error) {
	return file.NewFile(fileName, blob)
	// if err != nil {
	// 	return nil, err
	// }

	// return newFile,n

}
