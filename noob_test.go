package noob

import (
	"testing"
  _ "github.com/Hurka5/noob/langs/hu"
  _ "github.com/Hurka5/noob/langs/en"
)

func TestIsProfane(t *testing.T) {

	samples := map[string]bool{
		"france": false,
		"fasz":   true,
		"dick":   true,
		"cica":   false,
	}

	for word, isObscenity := range samples {
		got := IsProfane(word)
		if got != isObscenity {
			t.Errorf("IsProfane(%v) is %v; want %v", word, got, isObscenity)
		}
	}
}
