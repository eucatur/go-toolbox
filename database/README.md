# database.v2 #

É um wrapper do [SQLx](https://github.com/jmoiron/sqlx) com o objetivo de entrar uma conexão com banco de dados (MySQL, Postgres ou SQLite) somente lhe indicando o arquivo env com os paramentros de conexão

** MAX_OPEN_CONNS **
Use este parametro no env para alterar o máximo de conexoes simultaneas possíveis


## Exemplo ##

```code
package main

import (
	"github.com/eucatur/go-toolbox/database.v2"
)

func main() {
	db, err := database.GetByFile("sqlite3-config.json")

	if err != nil {
		panic(err)
	}

	var users map[string]string

	err = db.Select(&users, `select * from users`)

	if err != nil {
		panic(err)
	}
}
```