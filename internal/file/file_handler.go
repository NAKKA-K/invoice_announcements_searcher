package file

import (
	"fmt"
	"os"
)

func GetFileNames(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("fail to read directory: %w", err)
	}

	files := make([]string, 0, len(entries))
	for _, entry := range entries {
		name := entry.Name()
		if IsHidden(name) || entry.IsDir() {
			continue
		}
		files = append(files, name)
	}

	return files, nil
}

func IsHidden(name string) bool {
	return name[0] == '.'
}
