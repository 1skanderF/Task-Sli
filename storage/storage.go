package storage

import "os"

func (s *JSONStorage) Load() error {
	file, err = os.Open(s.filename)
	defer file.Close()

	if err != nil(
		if os.IsNotExist(err)(
			return nil
		)
	)
}
