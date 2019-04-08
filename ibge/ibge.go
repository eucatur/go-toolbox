package ibge

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/eucatur/go-toolbox/viacep"
)

func FindIBGEByZipCode(zipcode int) (result IBGEResult, err error) {
	ibgeStr, err := viacep.FindByZipCode(zipcode)
	if err != nil {
		return
	}

	ibge, err := strconv.Atoi(ibgeStr.Ibge)
	if err != nil {
		return
	}

	result, err = FindIBGEByCode(ibge)

	return
}

func FindIBGEByCode(codigo int) (result IBGEResult, err error) {
	resp, err := http.Get("http://servicodados.ibge.gov.br/api/v1/localidades/municipios/" + strconv.Itoa(int(codigo)))
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
