package lib

import (
	"os"
)

func IsFileExists(p string) bool {
	f, err := os.Stat(p)
	if err != nil {
		if os.IsExist(err) {
			if f.IsDir() {
				return false
			} else {
				return true
			}
		}
		return false
	}
	if f.IsDir() {
		return false
	} else {
		return true
	}
}

func IsPathExists(p string) bool {
	f, err := os.Stat(p)
	if err != nil {
		if os.IsExist(err) {
			if f.IsDir() {
				return true
			} else {
				return false
			}
		}
		return false
	}
	if f.IsDir() {
		return true
	} else {
		return false
	}
}

func CreatePath(filePath string) error {
	if !IsPathExists(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		return err
	}
	return nil
}

func CreateFile(file string) error {
	f, err := os.Create(file)
	defer func() {
		_ = f.Close()
	}()

	if err != nil {
		return err
	}
	return nil
}
