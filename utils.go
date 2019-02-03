package melli

import "strings"

func filter(r rune) rune {
	// TODO what to do with nimfasele های
	switch r {
	case 8205, 8207, 8235, 8236, 8238:
		return -1
	}
	return r
}

func clean(s string) string {
	s = strings.Map(filter, s)
	s = strings.Replace(s, "\u200C ", " ", -1)
	s = strings.Replace(s, " : ", ": ", -1)
	s = strings.Replace(s, " ...", "...", -1)
	s = strings.TrimSuffix(s, string('\u200C'))

	return strings.TrimSpace(s)
}
