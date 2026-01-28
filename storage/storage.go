package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

func (s *JSONStorage) Load() error {

	file, err1 := os.Open(s.filename)

	// Проверяем файл найден?
	if err1 != nil {
		if os.IsNotExist(err1) {
			// Если файла нет это не ошибка
			return nil
		}
		return err1
	}

	defer file.Close()

	fileInfo, err2 := file.Stat()

	if err2 != nil {
		return err2
	}
	if fileInfo.Size() == 0 {
		// Если файл пустой это нормально
		return nil
	}

	decoder := json.NewDecoder(file)
	err3 := decoder.Decode(&s.tasks)
	if err3 != nil {
		return fmt.Errorf("ошибка чтения JSON: %w", err3)
	}

	return nil
}
