package file

import (
	"encoding/json"
	"os"
)

func ParseJson[T comparable](s []byte) (T, error) {

	var result T

	err := json.Unmarshal(s, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func ParseFileJson[T comparable](path string) (T, error) {
	jsonFile, _ := os.ReadFile(path)

	result, err := ParseJson[T](jsonFile)
	if err != nil {
		return result, err
	}

	return result, nil
}
