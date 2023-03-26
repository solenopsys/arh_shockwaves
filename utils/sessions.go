package utils

import (
	os "os"
	"path/filepath"
)

const pattern = "xs-session-"

func WriteSessionToTempFile(data []byte) (string, error) {
	dir := os.TempDir()

	tempFile, err := os.CreateTemp(dir, pattern)
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	_, err = tempFile.Write(data)
	if err != nil {
		return "", err
	}

	return tempFile.Name(), err
}

func findYongestFile(files []string) (string, error) {
	var youngestFile string
	var youngestFileTime int64

	for _, file := range files {
		fileInfo, err := os.Stat(file)
		if err != nil {
			return "", err
		}
		if fileInfo.ModTime().Unix() > youngestFileTime {
			youngestFile = file
			youngestFileTime = fileInfo.ModTime().Unix()
		}
	}

	return youngestFile, nil
}

func ReadSessionFromTempFile() ([]byte, error) {
	dir := os.TempDir()
	filePattern := filepath.Join(dir, pattern+"*")
	files, err := filepath.Glob(filePattern)
	if err != nil {
		return nil, err
	}
	fileName, err := findYongestFile(files)

	if err != nil {
		return nil, err
	}
	bytes, err := ReadFile(fileName)
	return bytes, err
}
