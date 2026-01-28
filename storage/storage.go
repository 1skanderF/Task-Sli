package storage

import (
	"encoding/json"
	"fmt"
	"os"
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
	err = decoder.Decode(&s.tasks)
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
	err = encoder.Encode(s.tasks)
	if err != nil {
		os.Remove(tempFilename)
		return fmt.Errorf("ошибка записи JSON: %w", err)
	}

	// Перезаписываем файл
	err = os.Rename(tempFilename, s.filename)
	if err != nil {
		os.Remove(tempFilename)
		return err
	}

	return nil
}
