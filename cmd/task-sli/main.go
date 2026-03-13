package main

import (
	"fmt"

	"github.com/1skander/Task-Ski/models"
	"github.com/1skander/Task-Ski/storage"
)

func main() {
	storage, err := storage.NewJSONStorage("tasks.json")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	
	// Тест 1: Получение всех задач
	fmt.Println("=== Тест GetAllTasks ===")
	tasks, err := storage.GetAllTasks()
	if err != nil {
		fmt.Println("Ошибка GetAllTasks:", err)
	} else {
		fmt.Printf("Найдено задач: %d\n", len(tasks))
		for _, task := range tasks {
			fmt.Printf("  ID=%d: %s [%s]\n",
				task.ID, task.Description, task.Status)
		}
	}

	// Тест 2: Добавление + получение
	fmt.Println("\n=== Тест Добавление + Получение ===")
	newTask := models.NewTask("Новая задача для теста")
	newTask.ID = 50
	storage.AddTask(newTask)

	tasks, _ = storage.GetAllTasks()
	fmt.Printf("После добавления - задач: %d\n", len(tasks))
}
