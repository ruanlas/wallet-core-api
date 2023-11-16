package gainprojection

import uuid "github.com/satori/go.uuid"

type StorageProcess interface {
	Create(request CreateRequest) (*GainProjectionResponse, error)
}

type storageProcess struct {
	repository   Repository
	generateUUID func() uuid.UUID
}

func NewStorageProcess(repository Repository) StorageProcess {
	return &storageProcess{repository: repository}
}

func (sp *storageProcess) Create(request CreateRequest) (*GainProjectionResponse, error) {

}
