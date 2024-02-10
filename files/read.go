package files

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
)

func Read(root, filename string) ([]byte, error) {
	filename = strings.ReplaceAll(filename, "\\", string(os.PathSeparator))
	bts, err := os.ReadFile(filepath.Join(root, filename))
	if err != nil {
		// filenames in Zusi are case-insensitive, so if we are running on a case-sensitive system we have to search for that file
		var e *os.PathError
		if errors.As(err, &e) && errors.Is(e.Err, syscall.ENOENT) && runtime.GOOS != "windows" {
			filename, err = findFile(root, filename)
			if err != nil {
				return nil, fmt.Errorf("could not find file with case-insensitive search: %w", err)
			}
			bts, err = os.ReadFile(filepath.Join(root, filename))
			if err != nil {
				return nil, fmt.Errorf("error loading file: %w", err)
			}
		}
	}

	return bts, nil
}

func findFile(root, filename string) (string, error) {
	if runtime.GOOS == "windows" {
		return filepath.Join(root, filename), nil
	}

	resolvedPath := ""
	path := strings.Split(filename, string(filepath.Separator))

	for i := 0; i < len(path); i++ {
		dir, err := os.ReadDir(filepath.Join(root, resolvedPath))
		if err != nil {
			return "", fmt.Errorf("error loading directory: %w", err)
		}

		for _, entry := range dir {
			if strings.EqualFold(entry.Name(), path[i]) {
				resolvedPath = filepath.Join(resolvedPath, entry.Name())
				break
			}
		}
	}

	return resolvedPath, nil
}
