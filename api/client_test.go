package api

import "testing"

func TestGetBookURLByISBN(t *testing.T) {
	tests := []struct {
		isbn string
		exp  string
	}{
		{
			"9789646235793",
			"http://opac.nlai.ir/opac-prod/bibliographic/3399286",
		},
		{
			"9786008237631",
			"http://opac.nlai.ir/opac-prod/bibliographic/5134460",
		},
	}

	for i, test := range tests {
		url, err := GetBookURLByISBN(test.isbn)
		if err != nil {
			t.Errorf("Test %d: Error on getting book url by %s isbn: %s",
				i, test.isbn, err)
		}
		if url != test.exp {
			t.Errorf("Test %d: Expected %s but got %s for %s usbn.",
				i, test.exp, url, test.isbn)
		}
	}
}
