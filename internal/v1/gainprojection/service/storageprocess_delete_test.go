package service

import (
	"context"
	"errors"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestDeleteSuccess(t *testing.T) {
	_mockRepository := &mockRepository{}
	_mockRepository.AddRemoveCall(func(ctx context.Context, id string) error {
		return nil
	})

	uuidMock := func() uuid.UUID {
		return uuid.FromStringOrNil("cd1cc27b-28a1-47dc-ac76-70e8185e159d")
	}
	ctx := context.TODO()

	_storageProcess := NewStorageProcess(_mockRepository, uuidMock)
	err := _storageProcess.Delete(ctx, "cd1cc27b-28a1-47dc-ac76-70e8185e159d")
	assert.NoError(t, err)
}

func TestDeleteFail(t *testing.T) {
	_mockRepository := &mockRepository{}
	_mockRepository.AddRemoveCall(func(ctx context.Context, id string) error {
		return errors.New("An error has been ocurred")
	})

	uuidMock := func() uuid.UUID {
		return uuid.FromStringOrNil("cd1cc27b-28a1-47dc-ac76-70e8185e159d")
	}
	ctx := context.TODO()

	_storageProcess := NewStorageProcess(_mockRepository, uuidMock)
	err := _storageProcess.Delete(ctx, "cd1cc27b-28a1-47dc-ac76-70e8185e159d")
	assert.Error(t, err)
}
