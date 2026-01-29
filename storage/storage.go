package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/1skander/Task-Ski/models"
)

func (s *JSONStorage) Load() error {

	file, err := os.Open(s.filename)

	// Проверяем файл найден?
	if err != nil {
		if os.IsNotExist(err) {
			// Если файла нет это не ошибка
			return nil
		}
		return err
	}

	defer file.Close()

	fileInfo, err := file.Stat()

	if err != nil {
		return err
	}
	if fileInfo.Size() == 0 {
		// Если файл пустой это нормально
		return nil
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&s.Tasks)
	if err != nil {
		return fmt.Errorf("ошибка чтения JSON: %w", err)
	}

	return nil
}

func (s *JSONStorage) Save() error {
	// Блокируем параллельный вызов Save()
	s.mu.Lock()
	defer s.mu.Unlock()

	// Задаем имя для временного файла
	tempFilename := s.filename + ".tmp"

	// Создаем временный файл, если ошибка возращаем ошибку удаляем временный файл
	tempFile, err := os.Create(tempFilename)
	if err != nil {
		return err
	}

	// Откладываем функцию закрытия
	defer tempFile.Close()

	// Создаем JSON-энкодер для файла, настраиваем отступы, проверяем ошибку
	encoder := json.NewEncoder(tempFile)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(s.Tasks)
	if err != nil {
		os.Remove(tempFilename)
		return fmt.Errorf("ошибка записи JSON: %w", err)
	}

	// Перезаписываем файл
	tempFile.Close()
	err = os.Rename(tempFilename, s.filename)
	if err != nil {
		os.Remove(tempFilename)
		return err
	}

	return nil
}

func (s *JSONStorage) AddTask(task models.Task) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.Tasks[task.ID]; ok {
		return fmt.Errorf("Задача с ID %d уже существует", task.ID)
	}

	s.Tasks[task.ID] = task

	if err := s.Save(); err != nil {
		delete(s.Tasks, task.ID)
		return fmt.Errorf("Не удалось сохранить задачу %w", err)
	}

	return nil
}
