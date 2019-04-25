// Package env is a lib for managing environment variables, it
// makes it easy to add and fetch values
package env

import (
	"encoding/json"
	"errors"
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

// String search for a value between the environment variables
func String(key string) string {
	return os.Getenv(key)
}

// Int search for a value between the environment
// variables and convert to int
func Int(key string) (int, error) {
	return strconv.Atoi(String(key))
}

// MustInt search for a value between the environment
// variables, convert to int and generates a panic in case of error
func MustInt(key string) int {
	value, err := Int(key)
	if err != nil {
		panic(err)
	}
	return value
}

// Int64 search for a value between the environment variables
// and convert to int64
func Int64(key string) (int64, error) {
	return strconv.ParseInt(String(key), 10, 64)
}

// MustInt64 search for a value between the environment
// variables, convert to int64 and generates a panic in case of error
func MustInt64(key string) int64 {
	value, err := Int64(key)
	if err != nil {
		panic(err)
	}
	return value
}

// Float64 earch for a value between the environment variables
// and convert to float64
func Float64(key string) (float64, error) {
	return strconv.ParseFloat(String(key), 64)
}

// MustFloat64 search for a value between the environment
// variables, convert to float64 and generates a panic in case of error
func MustFloat64(key string) float64 {
	value, err := Float64(key)
	if err != nil {
		panic(err)
	}
	return value
}

// MustString search for a value between the environment
// variables generates a panic in case of error
func MustString(key string) (value string) {
	value = os.Getenv(key)
	if len(value) == 0 {
		panic(errors.New("The environment variable [" + key + "] is not set"))
	}
	return
}

// Float64 earch for a value between the environment variables
// and convert to float64
func Bool(key string) (bool, error) {
	return strconv.ParseBool(key)
}

// MustFloat64 search for a value between the environment
// variables, convert to float64 and generates a panic in case of error
func MustBool(key string) bool {
	value, err := Bool(key)
	if err != nil {
		panic(err)
	}
	return value
}
