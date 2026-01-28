package main

import (
	"fmt"
	"os"

	"github.com/1skander/Task-Ski/models"
	"github.com/1skander/Task-Ski/storage"
)

func main() {
	storage, err := storage.NewJSONStorage("tasks.json")

	if err != nil {
		fmt.Println("Ошибка создания хранилища:", err)
		return
	}
	// Добавляем тестовую задачу
	task := models.NewTask("Тест Save()")
	task.ID = 1
	storage.tasks[1] = task // пока напрямую, потом через AddTask

	// Сохраняем
	err = storage.Save()
	if err != nil {
		fmt.Println("Ошибка Save():", err)
	} else {
		fmt.Println("Save() успешен!")

		// Читаем файл чтобы проверить
		data, _ := os.ReadFile("tasks.json")
		fmt.Println("Содержимое файла:")
		fmt.Println(string(data))
	}
}
