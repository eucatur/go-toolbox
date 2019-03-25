package validator

import "encoding/json"

// VError é usada como retorno da validação
type VError struct {
	Message string
	Details interface{}
}

func (vError *VError) Error() string {
	return vError.Message
}

// MarshalJSON define a maneira que gera o JSON
func (vError *VError) MarshalJSON() ([]byte, error) {
	return json.Marshal(funcMarshalJSON(vError))
}
