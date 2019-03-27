# EUCATUR Go-ToolBox #
Repositório com o objetivo de compartilhar a "caixa de ferramentas" que utilizamos na empresa EUCATUR, para agilizar o nosso processo de desenvolvimento de soluções tecnológicas

## api ##

[api](https://github.com/eucatur/go-toolbox/tree/master/api) é um wrapper do [Echo](https://github.com/labstack/echo) com a configurações básicas para criar uma API REST em poucas linhas

## cache ##

[cache](https://github.com/eucatur/go-toolbox/tree/master/cache) É um wrapper do [go-cache](https://github.com/patrickmn/go-cache) uma lib de cache em memória com tempo de expiração, básicamente tem somente o metodo Set e Get

## cookie ##

[cookie](https://github.com/eucatur/go-toolbox/tree/master/cookie) É um lib para adicionar e deletar cookie no framework [Echo](https://github.com/labstack/echo)

## database ##

[database](https://github.com/eucatur/go-toolbox/tree/master/database) É um wrapper do [SQLx](https://github.com/jmoiron/sqlx) com o objetivo de entrar uma conexão com banco de dados (MySQL, Postgres ou SQLite) somente lhe indicando o arquivo env com os paramentros de conexão

## format ##

[format](https://github.com/eucatur/go-toolbox/tree/master/format) É um lib com funções de formatação para diversos tipos

## handler ##

[handler](https://github.com/eucatur/go-toolbox/tree/master/handler) É um lib para criar funções utilizadas em diversos handlers no framework [Echo](https://github.com/labstack/echo) como a BindAndValidade para fazer o bind na struct e validar ela

## json2env ##

[json2env](https://github.com/eucatur/go-toolbox/tree/master/json2env) é uma lib que le um arquivo json e coloca os valores no enviroment

## jwt ##

[jwt](https://github.com/eucatur/go-toolbox/tree/master/jwt) É um wrapper do [jwt-go](https://github.com/dgrijalva/jwt-go) para facilitar a utilização de jwt nos projetos

## log ##

[log](https://github.com/eucatur/go-toolbox/tree/master/log) é uma lib para lidar com log, para log em arquivo ou no terminal com a linha do arquivo com o erro 