package download

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// Download file (src) and write it to dst
func Download(src, dst string) error {
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	// create a new request
	client := &http.Client{}
	req, _ := http.NewRequest("GET", src, nil)
	req.Header.Set("User-Agent", "slashquery-download")
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	_, err = io.Copy(out, res.Body)
	if err != nil {
		return err
	}
	return nil
}

// Unzip extract contents of a zip file
func Unzip(zipfile, target string) error {
	reader, err := zip.OpenReader(zipfile)
	if err != nil {
		return err
	}
	defer reader.Close()

	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		fmt.Printf("path = %+v\n", path)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}
	return nil
}
