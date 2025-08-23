# UUID

É uma lib para gerar informações baseando-se em UUID seguindo o padrão [RFC 9562](https://www.rfc-editor.org/rfc/rfc9562)

## UUID v7

Recursos para gerar UUID v7 seguindo o padrão conforme documentação [5.7 UUID Version 7](https://www.rfc-editor.org/rfc/rfc9562)

> Essa lib irá gerar UUID v7 conforme time informado ao instanciar a struct UUIDv7 ou utilizando apenas a função GenerateUUIDv7. \
> Em ambos os recursos para geração do UUID v7 todos eles irão aplicar o timezone UTC.

### Testes

Para rodar os testes e validar a cobertura, base rodar os seguintes comandos:

```bash
go clean -testcache && \
go test -timeout 30s `go list ./... | grep -v -E 'mocks|interfaces'` -coverprofile=cover.out -covermode=atomic -coverpkg=./... ./... && \
echo "Cobertura: " $(go tool cover -func cover.out | grep total: | awk '{print $3}') && \
rm -rf cover.out;
```

### GenerateUUIDv7

```go
GenerateUUIDv7() (string, error) {}
```

Irá gerar uma informação UUID v7 instanciando a struct UUIDv7 e com base na data e hora atual será gerado a informação UUID

Caso ocorra erro no fluxo de geração irá retornar uma interface error

Utilização

```go
package main

import (
 "fmt"

 "github.com/eucatur/go-toolbox/uuid"
)

func main() {

 uuidv7, err := uuid.GenerateUUIDv7()

 if err != nil {
  fmt.Println("Error generating UUIDv7:", err)
  return
 }

 fmt.Println("Generated UUIDv7:", uuidv7)

}
```

### GetTimeUUIDv7

```go
GetTimeUUIDv7(uuid string) (time.Time, error)
```

Irá retornar um time a partir de um UUID v7 válido informado.

Caso não consiga extrair a informação ou o UUID v7 seja inválido irá retornar uma interface error

Utilização

```go
package main

import (
 "fmt"

 "github.com/eucatur/go-toolbox/uuid"
)

func main() {

 infoTime, err := uuid.GetTimeUUIDv7("0198d3ff-746c-74fb-9589-1686ec05eede")

 if err != nil {
  fmt.Println("Error to get time from UUID v7:", err)
  return
 }

 fmt.Println("Time from UUID v7:", infoTime)

}
```

## NewUUIDv7

```go
NewUUIDv7(t time.Time) (u UUIDv7)
```

Ao instanciar o recurso para gerar UUID v7 conforme o time informado terá acesso aos segunites métodos

```go
UUIDv7.Generate()
```

Irá gerar um UUID v7 conforme a instancia criada com a declaração time no parâmetro

```go
UUIDv7.String()
```

Irá retornar a informação em string do UUID v7 gerado

Utilização

```go
package main

import (
 "fmt"
 "time"

 "github.com/eucatur/go-toolbox/uuid"
)

func main() {

 timePerformed, err := time.Parse(time.RFC3339, "2025-08-22T19:07:38Z")

 if err != nil {
  fmt.Println("Error parsing time:", err)
  return
 }

 fmt.Println("Time performed:", timePerformed)

 uuidV7 := uuid.NewUUIDv7(timePerformed)

 uuidInfo, err := uuidV7.Generate()

 if err != nil {
  fmt.Println("Error generating UUIDv7:", err)
  return
 }

 fmt.Println("Generated UUIDv7:", uuidInfo)

 strUUIDv7 := uuidV7.String()
 fmt.Println("UUIDv7 as string:", strUUIDv7)

 infoTime, err := uuid.GetTimeUUIDv7(strUUIDv7)
 
 if err != nil {
  fmt.Println("Error getting time from UUIDv7:", err)
  return
 }
 
 fmt.Println("Time from UUIDv7:", infoTime)

}
```

### Referências

As referências utilizadas para desenvolvimento desta lib:

> RFC 9562 acessado em 2025-08-22 20:03:19, disponível em [https://www.rfc-editor.org/rfc/rfc9562#name-uuid-version-7](https://www.rfc-editor.org/rfc/rfc9562#name-uuid-version-7) \
UUID7.com acessado em 2025-08-22 20:03:33, disponível em [https://uuid7.com/](https://uuid7.com/) \
Repositório uuidv7 do usuário [Son Nguyen - @sonitx](https://github.com/sonitx) acessado em 2025-08-22 20:05:52, disponível em [https://github.com/sonitx/uuidv7](https://github.com/sonitx/uuidv7) \
Validador UUID v7 acessado em 2025-08-22 20:09:46, disponível em [https://ray.run/tools/uuid-validator](https://ray.run/tools/uuid-validator)
