package main

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func FileExists(f string) bool {

	if _, err := os.Stat(f); os.IsNotExist(err) {

		return false
	}

	return true
}

func HashFile(p string) (string, error) {

	h := sha256.New()

	f, err := os.Open(p)
	if err != nil {

		return "", err
	}
	defer f.Close()

	if _, err := io.Copy(h, f); err != nil {

		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func FetchRemoteFile(u string) ([]byte, error) {

	resp, err := http.Get(u)
	if err != nil {

		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {

		return nil, errors.New(fmt.Sprintf("Remote server returned status code >= 400: %s", resp.Status))
	}

	return ioutil.ReadAll(resp.Body)
}

func GetRemoteFileSize(u string) (int64, error) {

	resp, err := http.Get(u)
	if err != nil {

		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {

		return 0, errors.New(fmt.Sprintf("Remote server returned status code >= 400: %s", resp.Status))
	}

	return resp.ContentLength, nil
}

func GetFileSize(p string) (int64, error) {

	f, err := os.Stat(filepath.Clean(p))
	if err != nil {

		return 0, err
	}

	return f.Size(), nil
}
