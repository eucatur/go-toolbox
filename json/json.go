// Package json is a library that performs operations for json
package json

import (
	"encoding/json"
	"io/ioutil"
)

// UnmarshalFile does unmarshal from the path of a file
func UnmarshalFile(filePath string, v interface{}) error {
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(fileBytes, v)
}
