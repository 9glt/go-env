package env

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	os.Setenv("hey", "value!")
	os.Setenv("float", "5.15")
	os.Setenv("num", "5")
	os.Setenv("bool", "true")

	if Get("float").Float() != 5.15 {
		t.Fatal()
	}

	if Get("num").Int() != 5 {
		t.Fatal()
	}

	if Get("hey").String() != "value!" {
		t.Fatal()
	}

	if Get("bool").Bool() != true {
		t.Fatal()
	}
}
