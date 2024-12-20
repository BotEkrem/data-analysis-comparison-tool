package utils

import (
	"errors"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func GetCurrentPath() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic(errors.New("unable to get the current filename"))
	}

	dirname := filepath.Dir(filename)

	return dirname
}

func GetDirectories(path string) ([]os.DirEntry, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	return entries, err
}

func FindByExtension(root, ext string) []string {
	var files []string
	err := filepath.WalkDir(root, func(s string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(d.Name()) == ext {
			files = append(files, d.Name())
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	return files
}
