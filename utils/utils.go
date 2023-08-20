package utils

import "os"

func FindFiles(path string) ([]string, error) {
	entries, err := os.ReadDir(path)

	if err != nil {
		return nil, err
	}

	files := SliceFilter(entries, func(file os.DirEntry, _ int) bool {
		return !file.IsDir()
	})

	fileNames := SliceMap(files, func(file os.DirEntry, _ int) string {
		return path + file.Name()
	})

	return fileNames, nil
}
