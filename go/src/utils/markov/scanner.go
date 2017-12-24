package markov

import (
	"unicode/utf8"
)

// DefaultWordEquivalentChars is the set of characters that are
// "word equivalent", meaning that they should be included in the markov
// chain model as words
//
var DefaultWordEquivalentChars = map[rune]bool{
	',':  true,
	':':  true,
	'.':  true,
	';':  true,
	'(':  true,
	')':  true,
	'`':  true,
	'\'': true,
	'-':  true,
	'[':  true,
	']':  true,
}

// DefaultExcludedChars is the set of characters that will be excluded
// when scanning input. By default it excludes all unicode "space" characters.
var DefaultExcludedChars = map[rune]bool {
	' ':true,
	'\t':true,
	'\n':true,
	'\v':true,
	'\f':true,
	'\r':true,
	'\u0085':true,
	'\u00A0':true,
	'\u1680':true,
	'\u2028':true,
	'\u2029':true,
	'\u202f':true,
	'\u205f':true,
	'\u3000':true,
}


// NewWordScanner confirms to the bufio.Scanner type.
// It is for use when building the model from text.
func NewWordScanner(wordEq, excluded  map[rune]bool,) func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		start := 0
		for width := 0; start < len(data); start += width {
			var r rune
			r, width = utf8.DecodeRune(data[start:])
			if !excluded[r] {
				break
			}
		}

		for width, i := 0, start; i < len(data); i += width {
			var r rune
			r, width = utf8.DecodeRune(data[i:])
			if excluded[r] {
				return i + width, data[start:i], nil
			}

			if wordEq[r] && i != 0 {
				return i, data[start:i], nil
			}

			if wordEq[r] {
				return i + width, data[start : i+width], nil
			}
		}
		// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
		if atEOF && len(data) > start {
			return len(data), data[start:], nil
		}
		// Request more data.
		return start, nil, nil
	}
}