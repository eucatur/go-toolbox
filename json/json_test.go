package json

import "testing"

func Test(t *testing.T) {

	myJSON := struct {
		IP string `json:"ip"`
	}{}

	err := UnmarshalFile("test.json", &myJSON)

	if err != nil {
		t.Fatal(err)
	}

	if myJSON.IP != "192.168.1.1" {
		t.Error("The UnmarshalFile function did not return expected")
	}
}
