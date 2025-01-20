package noob

func lookupReplacement(r rune) rune {
  if val, ok := replacements[r]; ok {
    return val
  }
  return r
}

var replacements = map[rune]rune{
	'0': 'o',
	'1': 'i',
	'3': 'e',
	'4': 'a',
	'5': 's',
	'7': 'l',
	'$': 's',
	'!': 'i',
	'+': 't',
	'#': 'h',
	'@': 'a',
	'<': 'c',
  'Ł': 'l',
}
