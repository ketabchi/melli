package melli

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	tests := []struct {
		url     string
		expName string
	}{
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/636958",
			"سمفونی مردگان",
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/3399286",
			"شغل مناسب شما: با توجه به ویژگیهای شخصیتی خود کارتان را انتخاب کنید ...",
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/5030326",
			"مدیریت اجرایی (For Dummies (MBA",
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/5481844",
			"شدن",
		},
	}

	for i, test := range tests {
		book, err := NewBook(test.url)
		if err != nil {
			t.Errorf("Test %d: Error on creating book from %s: %s",
				i, test.url, err)
		}
		if name := book.Name(); name != test.expName {
			fmt.Println(test.expName, []rune(test.expName))
			fmt.Println(name, []rune(name))
			t.Errorf("Test %d: Expected book name '%s', but got '%s'",
				i, test.expName, name)
		}
	}
}
