package structcheck

import (
	"reflect"

	"emperror.dev/errors"
)

// CheckResourceInterfacesSetted
// Atente-se ao utilizar essa função:
// se a struct estiver como ponteiro informe-a diretamente
// se a struct não estiver como ponteiro informe-a por referência: Ex.: &struct
func CheckResourceInterfacesSetted(resource interface{}) (err error) {

	defer func() {
		errRecover := recover()

		if errRecover != nil {
			err = errors.Errorf("Não foi possível verificar os itens do recurso: %#v", resource)
		}
	}()

	if resource == nil {
		return errors.New("Recurso não definido para verificação das implementações")
	}

	rsc := reflect.ValueOf(resource).Elem()
	rscName := reflect.TypeOf(resource).Elem().Name()
	rscTp := rsc.Type()

	for i := 0; i < rsc.NumField(); i++ {
		field := rsc.Field(i)
		nameField := rscTp.Field(i).Name

		if field.Type().Kind() == reflect.Interface && field.Interface() == nil {

			err = errors.Append(err, errors.Errorf("A struct %s, possui o campo %s não definido", rscName, nameField))

		}
	}

	return
}

// CheckResourceStructsSetted
// Atente-se ao utilizar essa função:
// se a struct estiver como ponteiro informe-a diretamente
// se a struct não estiver como ponteiro informe-a por referência: Ex.: &struct
func CheckResourceStructsSetted(resource interface{}) (err error) {

	defer func() {
		errRecover := recover()

		if errRecover != nil {
			err = errors.Errorf("Não foi possível verificar os itens do recurso: %#v", resource)
		}
	}()

	if resource == nil {
		return errors.New("Recurso não definido para verificação das implementações")
	}

	rsc := reflect.ValueOf(resource).Elem()
	rscName := reflect.TypeOf(resource).Elem().Name()
	rscTp := rsc.Type()

	for i := 0; i < rsc.NumField(); i++ {
		field := rsc.Field(i)
		nameField := rscTp.Field(i).Name

		if field.Type().Kind() == reflect.Struct && field.IsZero() || (field.Type().Kind() == reflect.Pointer && field.IsNil()) {

			err = errors.Append(err, errors.Errorf("A struct %s, possui o campo %s não definido", rscName, nameField))

		}
	}

	return
}

// CheckStructFilled
// irá utiliza a verificação de interface e campos de uma struct
func CheckStructFilled(resources interface{}) error {

	return errors.Append(
		CheckResourceInterfacesSetted(resources),
		CheckResourceStructsSetted(resources),
	)

}
