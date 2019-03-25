package validator

var funcMarshalJSON = func(vError *VError) interface{} {
	return map[string]interface{}{
		"message": vError.Message,
		"details": vError.Details,
	}
}

// SetFuncMarshalJSON altera a forma de gerar o JSON
func SetFuncMarshalJSON(f func(vError *VError) interface{}) {
	funcMarshalJSON = f
}
