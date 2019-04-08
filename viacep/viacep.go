package viacep

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

// FindByZipCode Search the information of the locality informed through the zipcode
func FindByZipCode(zipcode int) (result ZipCode, err error) {
	resp, err := http.Get("https://viacep.com.br/ws/" + strconv.Itoa(int(zipcode)) + "/json/")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return
	}

	return
}
