package files

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
)

func Read(root, filename string) ([]byte, error) {
	filename = strings.ReplaceAll(filename, "\\", string(os.PathSeparator))
	bts, err := ioutil.ReadFile(filepath.Join(root, filename))
	if err != nil {
		// filenames in Zusi are case-insensitive, so if we are running on a case-sensitive system we have to search for that file
		if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOENT && runtime.GOOS != "windows" {
			filename, err = findFile(root, filename)
			if err != nil {
				return nil, errors.Wrap(err, "could not find file with case-insensitive search")
			}
			bts, err = ioutil.ReadFile(filepath.Join(root, filename))
			if err != nil {
				return nil, errors.Wrap(err, "error loading file")
			}
		} else {
			return nil, errors.Wrap(err, "error loading file")
		}
	}

	return bts, nil
}

func findFile(root, filename string) (string, error) {
	resolvedPath := ""
	path := strings.Split(filename, string(filepath.Separator))

	for i := 0; i < len(path); i++ {
		dir, err := ioutil.ReadDir(filepath.Join(root, resolvedPath))
		if err != nil {
			return "", errors.Wrap(err, "error loading directory")
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
