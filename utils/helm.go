package utils

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
)

func pushDir(archive []byte, name string) {

}

func ArchiveDir(dirName string) []byte {
	bufferWrite := new(bytes.Buffer)
	// Create a gzip writer
	gzipWriter := gzip.NewWriter(bufferWrite)
	defer gzipWriter.Close()

	// Create a tar writer
	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	// Walk through the directory and add each file to the archive
	filepath.Walk(dirName, func(path string, info os.FileInfo, err error) error {
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
		header.Name = relPath

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

	// Close the tar writer to flush any remaining data to the gzip writer
	if err := tarWriter.Close(); err != nil {
		panic(err)
	}

	return bufferWrite.Bytes()
}
