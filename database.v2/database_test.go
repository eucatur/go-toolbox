package database

import "testing"

func TestGetByFile(t *testing.T) {
	_, err := GetByFile("mysql-example.json")
	if err != nil {
		t.Error(err)
	}

	_, err = GetByFile("postgres-example.json")
	if err != nil {
		t.Error(err)
	}

	_, err = GetByFile("sqlite3-example.json")
	if err != nil {
		t.Error(err)
	}

	_, err = GetByFile("no-file-example.json")
	if err == nil {
		t.Error("The GetByFile function should not find the file.")
	}
}
