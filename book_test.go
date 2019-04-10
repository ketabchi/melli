package melli

import (
	"testing"

	"github.com/ketabchi/util"
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
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/5285471",
			"هوپا",
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/3766613",
			"پینه‌دوز",
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
		url    string
		faName string
		enName string
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
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/5120771",
			"ویکتوریا ترنبول",
			"Victoria Turnbull",
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/5373371",
			"برایان تریسی",
			"Brian Tracy",
		},
	}

	for i, test := range tests {
		book, err := NewBook(test.url)
		if err != nil {
			t.Errorf("Test %d: Error on creating book from %s: %s",
				i, test.url, err)
		}

		faName, enName := book.Author()
		if faName != test.faName {
			t.Logf("\n%q\n%q", test.faName, faName)
			t.Errorf("Test %d: Expected author faName '%s', but got '%s'",
				i, test.faName, faName)
		}
		if enName != test.enName {
			t.Logf("\n%q\n%q", test.enName, enName)
			t.Errorf("Test %d: Expected author enName '%s', but got '%s'",
				i, test.enName, enName)
		}
	}
}

func TestOriginalName(t *testing.T) {
	tests := []struct {
		url string
		exp string
	}{
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/5265395",
			"The paradox of choice: why more is less",
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/5363581",
			"Lying",
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/2072242",
			"Diary of a wimpy kid: Greg Heffley’s journal",
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/1092979",
			"A thousand splendid suns",
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/4630184",
			"Fahrenheit 451",
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/5481844",
			"Becoming",
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/1929190",
			"",
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/5171490",
			"Carta al General Franco",
		},
	}

	for i, test := range tests {
		book, err := NewBook(test.url)
		if err != nil {
			t.Errorf("Test %d: Error on creating book from %s: %s",
				i, test.url, err)
		}
		if name := book.OriginalName(); name != test.exp {
			t.Logf("\n%q\n%q", test.exp, name)
			t.Errorf("Test %d: Expected original name '%s', but got '%s'",
				i, test.exp, name)
		}
	}
}

func TestTranslator(t *testing.T) {
	tests := []struct {
		url string
		exp []string
	}{
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/5483716",
			[]string{"ارسلان فصیحی"},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/3125961",
			[]string{"محمدرضا طبیب‌زاده"},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/5174229",
			[]string{},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/4315430",
			[]string{"محمد عباس‌آبادی"},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/3608346",
			[]string{"آتوسا صالحی"},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/5112855",
			[]string{"هدا نژادحسینیان"},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/4336045",
			[]string{"پریسا صیادی", "سرور صیادی"},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/3997499",
			[]string{"فهیمه سیدناصری"},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/3973224",
			[]string{"عادل فردوسی‌پور", "بهزاد توکلی", "علی شهروز"},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/3393537",
			[]string{"مسعود رایگان"},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/5259291",
			[]string{"محمدامین رضایی", "فواد صبورنیا"},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/4392445",
			[]string{"امیرحسین میرزائیان", "عبدالرضا شهبازی"},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/5187277",
			[]string{"آذر متين", "نيلوفر اعظمی"},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/4427806",
			[]string{"محمدعلی جعفری"},
		},
	}

	for i, test := range tests {
		book, err := NewBook(test.url)
		if err != nil {
			t.Errorf("Test %d: Error on creating book from %s: %s",
				i, test.url, err)
		}
		if translators := book.Translators(); !util.CheckSliceEq(translators, test.exp) {
			t.Errorf("Test %d: Expected translators %q, but got %q",
				i, test.exp, translators)
		}
	}
}
