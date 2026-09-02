package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// CreateZipArchive archives multiple files into a single zip file
func CreateZipArchive(files []string, zipOutputPath string) error {
	zipFile, err := os.Create(zipOutputPath)
	if err != nil {
		return fmt.Errorf("failed to create zip file: %w", err)
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, file := range files {
		fileInfo, err := os.Stat(file)
		if err != nil {
			continue
		}

		header, err := zip.FileInfoHeader(fileInfo)
		if err != nil {
			continue
		}

		header.Name = filepath.Base(file)
		header.Method = zip.Deflate

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return fmt.Errorf("failed to write zip header for %s: %w", file, err)
		}

		fileToZip, err := os.Open(file)
		if err != nil {
			return fmt.Errorf("failed to open file %s for zip: %w", file, err)
		}

		_, err = io.Copy(writer, fileToZip)
		fileToZip.Close()
		if err != nil {
			return fmt.Errorf("failed to copy content of %s to zip: %w", file, err)
		}
	}

	return nil
}
