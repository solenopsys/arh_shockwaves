package tools

import (
	"crypto/md5"
	"encoding/hex"
	iio "io"
	"os"
	"path/filepath"
	"strings"
	"xs/pkg/io"
)

func computeHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := iio.Copy(hash, file); err != nil {
		return "", err
	}

	hashValue := hex.EncodeToString(hash.Sum(nil))
	return hashValue, nil
}

func hashDirectory(directoryPath string, excludedDirs []string) (string, error) {
	var concatenatedHash string

	err := filepath.Walk(directoryPath, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			dirName := filepath.Base(filePath)
			for _, excludedDir := range excludedDirs {
				if strings.EqualFold(dirName, excludedDir) {
					return filepath.SkipDir // Skip the directory and its contents
				}
			}
		} else {
			hash, err := computeHash(filePath)
			if err != nil {
				return err
			}
			concatenatedHash += hash
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	finalHash := computeFinalHash(concatenatedHash)
	return finalHash, nil
}

func computeFinalHash(concatenatedHash string) string {
	hash := md5.New()
	hash.Write([]byte(concatenatedHash))
	finalHash := hex.EncodeToString(hash.Sum(nil))
	return finalHash
}

func HashOfDir(path string, excludedDirs []string) (string, error) {

	finalHash, err := hashDirectory(path, excludedDirs)
	if err != nil {
		io.Println("Error:", err)
		return "", err
	} else {
		return finalHash, nil
	}
}
