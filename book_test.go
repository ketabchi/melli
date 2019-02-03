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
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/5114665",
			"موسسه فرهنگی هنری شهرستان ادب",
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/3729525",
			"موسسه فرهنگی هنری شهرستان ادب",
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/4528181",
			"نیماژ",
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/511069",
			"افق",
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/2345835",
			"افق",
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/2891053",
			"پرشیا شمع و مه",
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/3388150",
			"شهر قلم",
		},
	}

	for i, test := range tests {
		book, err := NewBook(test.url)
		if err != nil {
			t.Errorf("Test %d: Error on creating book from %s: %s",
				i, test.url, err)
		}
		if name := book.Publisher(); name != test.exp {
			t.Logf("\n%q\n%q", test.exp, name)
			t.Errorf("Test %d: Expected publisher name '%s', but got '%s'",
				i, test.exp, name)
		}
	}
}

func TestAuthor(t *testing.T) {
	tests := []struct {
		url   string
		name  string
		eName string
	}{
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/5309538",
			"گری نورثفیلد",
			"Gary Northfield",
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/4929459",
			"ریک ریوردان",
			"Rick Riordan",
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/3649724",
			"عفت‌السادات مرقاتی خویی",
			"",
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/4834116",
			"لیز پیشون",
			"Liz Pichon",
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/3735689",
			"زین‌الدین‌بن علی شهیدثانی",
			"",
		},
	}

	for i, test := range tests {
		book, err := NewBook(test.url)
		if err != nil {
			t.Errorf("Test %d: Error on creating book from %s: %s",
				i, test.url, err)
		}

		name, eName := book.Author()
		if name != test.name {
			t.Logf("\n%q\n%q", test.name, name)
			t.Errorf("Test %d: Expected author name '%s', but got '%s'",
				i, test.name, name)
		}
		if eName != test.eName {
			t.Logf("\n%q\n%q", test.eName, eName)
			t.Errorf("Test %d: Expected author eName '%s', but got '%s'",
				i, test.eName, eName)
		}
	}
}
