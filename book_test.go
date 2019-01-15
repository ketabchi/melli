package melli

import "testing"

func TestName(t *testing.T) {
	tests := []struct {
		isbn    string
		expName string
	}{
		{"9789643113445", "سمفونی مردگان"},
		{"9789643114794", "تسلی بخشی‌های فلسفه"},
		{"9789646235793", "ش‍غ‍ل‌ م‍ن‍اس‍ب‌ ش‍م‍ا: ب‍ا ت‍وج‍ه‌ ب‍ه‌ وی‍ژگ‍ی‌ه‍ای‌ ش‍خ‍ص‍ی‍ت‍ی‌ خ‍ود ک‍ارت‍ان‌ را ان‍ت‍خ‍اب‌ ک‍ن‍ی‍د ..."},
	}

	for i, test := range tests {
		book, err := NewBookByISBN(test.isbn)
		if err != nil {
			t.Errorf("Test %d: Error on creating book from isbn: %s",
				i, err)
		}
		if name := book.Name(); name != test.expName {
			t.Errorf("Test %d: Expected book name %s, but got %s",
				i, test.expName, name)
		}
	}
}
