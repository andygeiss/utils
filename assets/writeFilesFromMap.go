package assets

import (
	"errors"
	"os"
	"path/filepath"
)

// WriteFilesFromMap writes every file from the map to the file system.
func WriteFilesFromMap(in map[string][]byte, path string) (err error) {
	for base, content := range in {
		_, err = os.Stat(path)
		if errors.Is(err, os.ErrNotExist) {
			if err := os.MkdirAll(filepath.Dir(base), 0755); err != nil {
				return err
			}
		}
		if err := os.WriteFile(base, content, 0644); err != nil {
			return err
		}
	}
	return nil
}
