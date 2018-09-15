package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const imagesDir = "image_script/images/"
const inputFilenameBase = "input"

// DownloadImage downloads an image into the images directory
func DownloadImage(url string) (string, error)  {
	filename := getFilename(url)
	out, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func getFilename(url string) string {
	filename := fmt.Sprintf("%s%s", inputFilenameBase, filepath.Ext(url))
	
	return filepath.Join(imagesDir, filename)
}