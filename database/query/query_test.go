package query

import "testing"

func TestNew(t *testing.T) {
	var name = "Henrique"

	q := New("SELECT * FROM Pessoas WHERE true")

	q.And(name == "Henrique", " AND Pessoa.Name = ? ", name)

	if q.String() != "SELECT * FROM Pessoas WHERE true AND Pessoa.Name = ? " {
		t.Error("query err")
	}

	if len(q.Args()) != 1 {
		t.Error("args err")
	}
}
