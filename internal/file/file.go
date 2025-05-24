package file

import "github.com/google/uuid"

type File struct {
	Name string
	Data []byte
	ID   uuid.UUID
}

func NewFile(fileName string, blob []byte) (*File, error) {
	fileID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	return &File{
		ID:   fileID,
		Name: fileName,
		Data: blob,
	}, nil
}
