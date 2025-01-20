package noob

import (
	"github.com/cloudflare/ahocorasick"
	"github.com/Hurka5/noob/words"
)

var stringMatcher *ahocorasick.Matcher

func init() {
  stringMatcher = ahocorasick.NewStringMatcher(words.AllWords)
}

func unmaskText(s string) string {
  runes := []rune(s)
  for i := range runes {
    runes[i] = lookupReplacement(runes[i])
  }
  return string(runes)
}

func IsProfane(s string) bool {
  unmasked := unmaskText(s)
  hits := stringMatcher.Match([]byte(unmasked))
  for _, hit := range hits {
    println(" ", hit)
  }
  return len(hits) != 0
}

func Censor(s string) string {
  return ""
}

func ExtractProfanity(s string) []string {
  return []string{}
}
