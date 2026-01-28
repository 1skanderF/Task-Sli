package storage

import (
	"sync"

	"github.com/1skander/Task-Ski/models"
)

type JSONStorage struct {
	filename string
	mu       sync.RWMutex
	Tasks    map[int]models.Task
}

func NewJSONStorage(file string) (*JSONStorage, error) {
	storage := &JSONStorage{
		filename: file,
	}
	storage.tasks = make(map[int]models.Task)
	return storage, nil
}
