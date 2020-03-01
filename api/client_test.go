package api

import "testing"

func TestGetBookURLByISBN(t *testing.T) {
	tests := []struct {
		isbn  string
		title []string
		exp   string
	}{
		{
			"9789646235793",
			[]string{},
			"http://opac.nlai.ir/opac-prod/bibliographic/5800683",
		},
		{
			"9786008237631",
			[]string{},
			"http://opac.nlai.ir/opac-prod/bibliographic/5134460",
		},
		{
			"9786006860152",
			[]string{"کودک باهوش(4سالگی)مهارت‌نوشتن(کتاب‌‌پرنده) #"},
			"http://opac.nlai.ir/opac-prod/bibliographic/4634555",
		},
		{
			"9789642616763",
			[]string{"ویتامین‎های موفقیت(نقش‎نگین)"},
			"http://opac.nlai.ir/opac-prod/bibliographic/2055747",
		},
		{
			"9786009639267",
			[]string{"فرسنگ‌های نزدیک(کاویان‌کتاب)"},
			"",
		},
		{
			"9786009686162",
			[]string{"این من هستم(4ج،همراه‌کیف)لوپه‌تو #"},
			"",
		},
	}

	for i, test := range tests {
		url, err := GetBookURLByISBN(test.isbn, test.title...)
		if err != nil {
			t.Errorf("Test %d: Error on getting book url by %s isbn: %s",
				i, test.isbn, err)
		}
		if url != test.exp {
			t.Errorf("Test %d: Expected %s but got %s for %s isbn.",
				i, test.exp, url, test.isbn)
		}
	}
}
