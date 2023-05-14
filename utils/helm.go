package utils

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func PushDir(archive []byte) {
	buffer := bytes.NewBuffer(archive)

	url := "http://helm.solenopsys.org/api/charts"
	resp, err := http.Post(url, "application/octet-stream", buffer)
	if err != nil {
		fmt.Println("Error sending request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		os.Exit(1)
	}

	// Print the response body
	fmt.Println(string(responseBody))
}

func ArchiveDir(dirName string, parentDir string) []byte {
	bufferWrite := bytes.NewBuffer([]byte{})

	// Create a gzip writer
	gzipWriter := gzip.NewWriter(bufferWrite)

	// Create a tar writer
	tarWriter := tar.NewWriter(gzipWriter)

	// Walk through the directory and add each file to the archive
	err := filepath.Walk(dirName, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Create a tar header from the file info
		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}

		// Set the name of the file within the archive
		relPath, err := filepath.Rel(dirName, path)
		if err != nil {
			return err
		}
		header.Name = filepath.Join(parentDir, relPath)

		// Write the tar header
		if err := tarWriter.WriteHeader(header); err != nil {
			return err
		}

		// If the file is not a directory, write its contents to the archive
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			if _, err := io.Copy(tarWriter, file); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	// Close the tar writer to flush any remaining data to the gzip writer
	if err := tarWriter.Close(); err != nil {
		panic(err)
	}

	gzipWriter.Close()
	tarWriter.Close()

	return bufferWrite.Bytes()
}