package main

import (
	"fmt"

	"github.com/1skander/Task-Ski/models"
)

func main() {

	task := models.NewTask("Доделать задачу")
	fmt.Println(task)
}
