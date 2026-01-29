package main

import (
	"github.com/1skander/Task-Ski/models"
	"github.com/1skander/Task-Ski/storage"
)

func main() {
	task := models.NewTask("Проверить AddTask")
	task.ID = 1
	err := storage.AddTask(task)
	if err != nil {
		return
	}
}
