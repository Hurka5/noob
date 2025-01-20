package noob

import (
	_ "github.com/Hurka5/noob/soft/en" // Soft english   language filter
	_ "github.com/Hurka5/noob/soft/hu" // Soft hungarian language filter
	"testing"
)

func TestIsProfane(t *testing.T) {

	samples := map[string]bool{
		"france is not a bad word": false,
		"fasz is :(":               true,
		"d!ck":                     true,
		"cica":                     false,
		"dickface":                 true,
	}

	for word, isObscenity := range samples {
		got := IsProfane(word)
		if got != isObscenity {
			t.Errorf("IsProfane(%v) is %v; want %v", word, got, isObscenity)
		}
	}
}

func TestCensor(t *testing.T) {

	samples := map[string]string{
		"france is not a bad word": "france is not a bad word",
		"fasz is :(":               "**** is :(",
		"d!ck":                     "****",
		"cica":                     "cica",
		"dickface":                 "****face",
	}

	for word, answer := range samples {
		got := Censor(word)
		if got != answer {
			t.Errorf("Censor(\"%v\") is \"%v\"; want \"%v\"", word, got, answer)
		}
	}
}
