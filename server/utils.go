package server

import (
	"io/ioutil"
	"net/http"
	"os"
)

// IsExist checks whether a file or directory exists.
// It returns false when the file or directory does not exist.
func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

// Get child dirs
func walkDir(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var result []string
	for _, file := range files {
		if file.IsDir() {
			if file.Name() != ".git" { //忽略项目中的.git文件夹
				result = append(result, file.Name())
			}
		}
	}
	return result, nil
}

func ReadFile(path string) (string, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func GetWebPage(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return string(body), err
	}

	return string(body), nil
}
