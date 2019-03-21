# JSON2Env #

Is a library that performs operations for json


## Exemplo ##

```code

//test.json
{
  "ip":"192.168.1.1"
}

//exemple.go
package main

import (
	"log"

	"github.com/eucatur/go-toolbox/json"
)

func main() {

	myJSON := struct {
		IP string `json:"ip"`
	}{}

	if err := json.UnmarshalFile("test.json", &myJSON); err != nil {
		panic(err)
	}

	log.Println(myJSON)
}
```