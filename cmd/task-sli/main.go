package main

import (
	"fmt"

	"github.com/1skander/Task-Ski/models"
)

func main() {

	fmt.Println("1")
	task := models.NewTask("Доделать задачу")
	fmt.Println(task)
}
