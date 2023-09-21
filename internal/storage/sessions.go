package storage

import (
	os "os"
	"path/filepath"
)

//const pattern = "xs-session-"

type TempStorage struct {
	FilePattern string
}

func NewSessionStorage(pattern string) *TempStorage {
	return &TempStorage{FilePattern: pattern}
}

func (ts *TempStorage) WriteSessionToTempFile(data []byte) (string, error) {
	dir := os.TempDir()

	tempFile, err := os.CreateTemp(dir, ts.FilePattern)
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

func (ts *TempStorage) findYongestFile(files []string) (string, error) {
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

func (ts *TempStorage) findFiles() ([]string, error) {
	dir := os.TempDir()
	filePattern := filepath.Join(dir, ts.FilePattern+"*")
	return filepath.Glob(filePattern)
}

func (ts *TempStorage) DeleteSessionTempFiles() error {
	files, err := ts.findFiles()
	if err != nil {
		return err
	}

	for _, file := range files {
		err = os.Remove(file)
		if err != nil {
			return err
		}
	}

	return nil

}

func (ts *TempStorage) ReadSessionFromTempFile() ([]byte, error) {
	files, err := ts.findFiles()
	if err != nil {
		return nil, err
	}
	if files == nil {
		return nil, nil
	}
	fileName, err := ts.findYongestFile(files)

	if err != nil {
		return nil, err
	}
	bytes, err := os.ReadFile(fileName)
	return bytes, err
}
