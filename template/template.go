package template

import (
	"bytes"
	"strings"
	"text/template"
)

//ExecFile ...
func ExecuteFile(pathFile string, str interface{}) (res string, err error) {
	tempBuffer := new(bytes.Buffer)
	t := template.Must(template.ParseGlob(pathFile))
	err = t.Execute(tempBuffer, str)

	if err != nil {
		return
	}

	res = strings.TrimSpace(tempBuffer.String())
	return
}
