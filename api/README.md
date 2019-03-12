# api #

api é um wrapper do [Echo](https://github.com/labstack/echo) com a configurações básicas para criar uma API REST em poucas linhas, pode ser passado duas flags

**port**

Fala a porta em que o servidor da api irá rodar, o default é 9000

```code
--port=8080 //default 9000

```

**debug**

Informa se a API irá subir em modo debug, se for com true, irá loggar fazer um log de cada request feita

```code
--debug=true //default false
```

**Examplo para rodar em forma de teste**

```code
go run file.go --port=8080 --degub=true
```

## Exemplo ##

```code
package main

import (
	"net/http"

	"github.com/eucatur/go-toolbox/api"
	"github.com/labstack/echo"
)

func main() {
	server := api.Make()

	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})

	api.Run()
}
```