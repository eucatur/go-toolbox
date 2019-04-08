package ibge

type IBGEResult struct {
	ID                  int `json:"id"`
	Nome                string `json:"nome"`
	MicroRegiao         MicroRegiao `json:"microrregiao"`
}

type MicroRegiao struct {
	ID                  int `json:"id"`
	Nome                string `json:"nome"`
	MesorRegiao         MesorRegiao `json:"mesorregiao"`
}

type MesorRegiao struct {
	ID                  int `json:"id"`
	Nome                string `json:"nome"`
	UF         UF `json:"uf"`
}

type UF struct {
	ID                  int `json:"id" db:"ID"`
	Sigla                string `json:"sigla" db:"Sigla"`
	Nome                string `json:"nome" db:"Nome"`
	Regiao         Regiao `json:"regiao"`
}

type Regiao struct {
	ID                  int `json:"id"`
	Sigla                string `json:"sigla"`
	Nome                string `json:"nome"`
}