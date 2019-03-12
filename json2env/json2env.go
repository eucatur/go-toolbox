package json2env

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

func LoadFile(env_json_file_path string) (err error) {
	fileBytes, err := ioutil.ReadFile(env_json_file_path)

	if err != nil {
		return errors.New("File: " + env_json_file_path + " not found")
	}

	var key_values map[string]string

	if err = json.Unmarshal(fileBytes, &key_values); err != nil {
		log.Panic(err)
	}

	for key, value := range key_values {
		if err := os.Setenv(key, value); err != nil {
			log.Panic(err)
		}
	}

	return
}
