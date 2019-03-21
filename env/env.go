// Package env is a lib for managing environment variables, it
// makes it easy to add and fetch values
package env

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
)

// Set adds a value in the environment variables
func Set(key, value string) error {
	return os.Setenv(key, value)
}

// MustSet adds a value in the environment variables
// and generates a panic in case of error
func MustSet(key, value string) {
	err := Set(key, value)
	if err != nil {
		panic(err)
	}
}

// SetByJSONFile sets environment variables via json file
func SetByJSONFile(filePath string) error {
	var (
		fileBytes []byte
		values    map[string]string
		err       error
	)

	if fileBytes, err = ioutil.ReadFile(filePath); err != nil {
		return err
	}

	if err = json.Unmarshal(fileBytes, &values); err != nil {
		return err
	}

	for key, value := range values {
		if err = Set(key, value); err != nil {
			return err
		}
	}

	return nil
}

// MustSetByJSONFile sets environment variables via json file
// and generates a panic in case of error
func MustSetByJSONFile(filePath string) {
	if err := SetByJSONFile(filePath); err != nil {
		panic(err)
	}
}

// Get search for a value between the environment variables
func Get(key string) string {
	return os.Getenv(key)
}

// GetInt search for a value between the environment
// variables and convert to int
func GetInt(key string) (int, error) {
	return strconv.Atoi(Get(key))
}

// MustGetInt search for a value between the environment
// variables, convert to int and generates a panic in case of error
func MustGetInt(key string) int {
	value, err := GetInt(key)
	if err != nil {
		panic(err)
	}
	return value
}

// GetInt64 search for a value between the environment variables
// and convert to int64
func GetInt64(key string) (int64, error) {
	return strconv.ParseInt(Get(key), 10, 64)
}

// MustGetInt64 search for a value between the environment
// variables, convert to int64 and generates a panic in case of error
func MustGetInt64(key string) int64 {
	value, err := GetInt64(key)
	if err != nil {
		panic(err)
	}
	return value
}

// GetFloat64 earch for a value between the environment variables
// and convert to float64
func GetFloat64(key string) (float64, error) {
	return strconv.ParseFloat(Get(key), 64)
}

// MustGetFloat64 search for a value between the environment
// variables, convert to float64 and generates a panic in case of error
func MustGetFloat64(key string) float64 {
	value, err := GetFloat64(key)
	if err != nil {
		panic(err)
	}
	return value
}
