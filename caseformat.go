package caseformat

import (
	"unicode"
)

func ToUpperCamel(str string) string {
	s := ""
	for _, buf := range tokenize(str) {
		buf[0] = unicode.ToUpper(buf[0])
		s += string(buf)
	}
	return s
}

func ToLowerCamel(str string) string {
	s := ""
	for i, buf := range tokenize(str) {
		if i > 0 {
			buf[0] = unicode.ToUpper(buf[0])
		}
		s += string(buf)
	}
	return s
}

func ToLowerUnderscore(str string) string {
	s := ""
	for i, buf := range tokenize(str) {
		if i > 0 {
			s += "_"
		}
		s += string(buf)
	}
	return s
}

func ToUpperUnderscore(str string) string {
	s := ""
	for i, buf := range tokenize(str) {
		if i > 0 {
			s += "_"
		}
		for i, char := range buf {
			buf[i] = unicode.ToUpper(char)
		}
		s += string(buf)
	}
	return s
}

func tokenize(str string) [][]rune {

	parts := make([][]rune, 0)
	buf := []rune(str)
	bufLen := len(buf)
	current := make([]rune, 0)

	for i := 0; i < bufLen; i += 1 {
		canLookBehind, canLookAhead1, canLookAhead2 := i > 0, i < bufLen-1, i < bufLen-2
		aLower, bLower := unicode.IsLower(buf[i]), canLookAhead1 && unicode.IsLower(buf[i+1])
		aLetter, bLetter := unicode.IsLetter(buf[i]), canLookAhead1 && unicode.IsLetter(buf[i+1])
		behindLower, ahead2Lower := canLookBehind && unicode.IsLower(buf[i-1]), canLookAhead2 && unicode.IsLower(buf[i+2])
		split := !bLetter || (aLower && !bLower) || (!aLower && !bLower && (behindLower || ahead2Lower))
		if aLetter {
			current = append(current, unicode.ToLower(buf[i]))
			if split {
				parts = append(parts, current)
				current = make([]rune, 0)
			}
		}
	}

	if len(current) > 0 {
		parts = append(parts, current)
	}

	return parts
}
