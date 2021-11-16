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
			"شغل مناسب شما: با توجه به ویژگی\u200cهای شخصیتی خود کارتان را انتخاب کنید...",
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
			"طلبه زیستن: پژوهشی مقدماتی در سنخ\u200cشناسی جامعه\u200cشناختی زیست\u200cطلبگی",
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/1070294",
			"ارتباط رو در رو: کلید موفقیت برای مدیریت موثر و کارا مجموعه مقالاتی از دانشگاه هاروارد...",
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/3553118",
			"دریدا و فلسفه",
		},
	}

	for i, test := range tests {
		book, err := NewBook(test.url)
		if err != nil {
			t.Errorf("Test %d: Error on creating book from %s: %s",
				i, test.url, err)
		}
		if name := book.Name(); name != test.exp {
			t.Logf("\n%q\n%q", test.exp, name)
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
			"پینه\u200cدوز",
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
			"عفت\u200cالسادات مرقاتی خویی",
			"",
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/4834116",
			"لیز پیشون",
			"Liz Pichon",
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/3735689",
			"زین\u200cالدین\u200cبن علی شهیدثانی",
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
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/7356042",
			"ملکم",
			"",
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
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/4722298",
			"Whispers in the walls",
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
			[]string{"محمدرضا طبیب\u200cزاده"},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/5174229",
			[]string{},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/4315430",
			[]string{"محمد عباس\u200cآبادی"},
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
			[]string{"عادل فردوسی\u200cپور", "بهزاد توکلی", "علی شهروز"},
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
			"http://opac.nlai.ir/opac-prod/bibliographic/4427806",
			[]string{"محمدعلی جعفری"},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/5028775",
			[]string{"فریبا شریفی"},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/969350",
			[]string{"محمد عالمی"},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/1557799",
			[]string{"بهرام قاسمی\u200cنژاد"},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/5293382",
			[]string{"فرزام کریمی"},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/5451160",
			[]string{"ایمان گنجی", "محدثه زارع"},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/5631247",
			[]string{"احمد اخوت"},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/1126271",
			[]string{"تارا سالک"},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/4788927",
			[]string{"سارا طاهری", "علیرضا کوشکی\u200cجهرمی"},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/4235490",
			[]string{"مهدی شفقتی"},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/5009256",
			[]string{"لیلا کاشانی وحید"},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/5081800",
			[]string{"فاطمه صادقیان"},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/4912007",
			[]string{"لیلا کاشانی\u200cوحید"},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/4912007",
			[]string{"لیلا کاشانی\u200cوحید"},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/1983690",
			[]string{"مصطفی رحماندوست"},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/1983689",
			[]string{"مصطفی رحماندوست"},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/1983687",
			[]string{"مصطفی رحماندوست"},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/2900920",
			[]string{"مجتبی مقصودی", "الهه علوی", "مسعود جوادیان"},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/760159",
			[]string{"نجف دریابندری"},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/6239468",
			[]string{"انشاء\u200cالله رحمتی"},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/4165246",
			[]string{},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/1127515",
			[]string{"محمدعلی فروغی"},
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

func TestSerie(t *testing.T) {
	tests := []struct {
		url string
		exp []string
	}{
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/5438339",
			[]string{},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/4914535",
			[]string{"پرسی جکسون"},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/4573407",
			[]string{"رمان نوجوان", "قهرمانان المپ"},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/2854139",
			[]string{"پرسی جکسون و فرمانروایان آلپ"},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/2893901",
			[]string{"سی و نه سرنخ", "مجموعه کارآگاهی نشر ویدا"},
		},
		{
			"http://opac.nlai.ir/opac-prod/bibliographic/5030326",
			[]string{"کتاب\u200cهای دامیز٬ کاربردی و سودمند"},
		},
	}

	for i, test := range tests {
		book, err := NewBook(test.url)
		if err != nil {
			t.Errorf("Test %d: Error on creating book from %s: %s",
				i, test.url, err)
		}
		if series := book.Series(); !util.CheckSliceEq(series, test.exp) {
			t.Errorf("Test %d: Expected series %q, but got %q",
				i, test.exp, series)
		}
	}
}
