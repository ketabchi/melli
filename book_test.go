package melli

import (
	"testing"
)

func TestName(t *testing.T) {
	tests := []struct {
		url string
		exp string
	}{
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/636958",
			"سمفونی مردگان",
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/3399286",
			"شغل مناسب شما: با توجه به ویژگی‌های شخصیتی خود کارتان را انتخاب کنید...",
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/5030326",
			"مدیریت اجرایی (For Dummies (MBA",
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/5481844",
			"شدن",
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/3049599",
			"طلبه زیستن: پژوهشی مقدماتی در سنخ‌شناسی جامعه‌شناختی زیست‌طلبگی",
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/1070294",
			"ارتباط رو در رو: کلید موفقیت برای مدیریت موثر و کارا مجموعه مقالاتی از دانشگاه هاروارد...",
		},
	}

	for i, test := range tests {
		book, err := NewBook(test.url)
		if err != nil {
			t.Errorf("Test %d: Error on creating book from %s: %s",
				i, test.url, err)
		}
		if name := book.Name(); name != test.exp {
			t.Errorf("Test %d: Expected book name '%s', but got '%s'",
				i, test.exp, name)
		}
	}
}

func TestPublisher(t *testing.T) {
	tests := []struct {
		url string
		exp string
	}{
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/636958",
			"ققنوس",
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/3399286",
			"نقش و نگار",
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/5030326",
			"آوند دانش",
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/5481844",
			"مهر اندیش",
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/5355759",
			"ثالث",
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/3382881",
			"زعفران",
		},
	}

	for i, test := range tests {
		book, err := NewBook(test.url)
		if err != nil {
			t.Errorf("Test %d: Error on creating book from %s: %s",
				i, test.url, err)
		}
		if name := book.Publisher(); name != test.exp {
			t.Errorf("Test %d: Expected publisher name '%s', but got '%s'",
				i, test.exp, name)
		}
	}
}
