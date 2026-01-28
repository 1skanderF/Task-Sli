package main

import (
	"fmt"

	"github.com/1skander/Task-Ski/storage"
)

func main() {
	storage, err := storage.NewJSONStorage("tasks.json")
	if err != nil {
		fmt.Println("Ошибка создания хранилища:", err)
		return
	}

	// Проверяем Load
	err = storage.Load()
	if err != nil {
		fmt.Println("Ошибка загрузки:", err)
	} else {
		fmt.Println("Загрузка успешна")
	}
}
