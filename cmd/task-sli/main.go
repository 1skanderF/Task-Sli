package main

import (
	"fmt"

	"github.com/1skander/Task-Ski/models"
	"github.com/1skander/Task-Ski/storage"
)

func main() {
	storage, err := storage.NewJSONStorage("tasks.json")
	if err != nil {
		return
	}

	task1 := models.NewTask("Проверить AddTask")
	task1.ID = 1
	err = storage.AddTask(task1)
	if err != nil {
		return
	} else {
		fmt.Printf("Задача успешно добавлена с ID: %d", task1.ID)
	}

}
