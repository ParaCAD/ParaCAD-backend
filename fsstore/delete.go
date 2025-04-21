package fsstore

import (
	"os"
	"path/filepath"
)

func (fs *FSStore) DeleteFile(fileName string) error {
	filePath := filepath.Join(fs.storageDirectory, fileName)

	err := os.Remove(filePath)
	if err != nil {
		return err
	}

	return nil
}
