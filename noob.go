package noob

import (
	"github.com/Hurka5/noob/words"
	"github.com/cloudflare/ahocorasick"
	"strings"
)

var stringMatcher *ahocorasick.Matcher

func init() {
	stringMatcher = ahocorasick.NewStringMatcher(words.AllWords)
}

func IsProfane(s string) bool {
	unmasked := unmaskText(strings.ToLower(s))
	hits := stringMatcher.MatchThreadSafe([]byte(unmasked))
	return len(hits) > 0
}

func Censor(s string) string {
	unmasked := unmaskText(strings.ToLower(s))
	hits := stringMatcher.MatchThreadSafe([]byte(unmasked))

	runeText := []rune(unmasked)

	for _, index := range hits {
		w := words.AllWords[index]
		start := strings.Index(unmasked, w)
		if start == -1 {
			continue
		}
		for i := start; i < start+len(w); i++ {
			runeText[i] = '*'
		}
	}

	return string(runeText)
}

func ExtractProfanity(s string) []string {
	// unmasked := unmaskText(strings.ToLower(s))
	return []string{}
}

// --- private ----------
func unmaskText(s string) string {
	return strings.Map(func(r rune) rune {
		nr := lookupReplacement(r)
		return nr
	}, s)
}
