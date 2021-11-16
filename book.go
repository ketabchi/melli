package melli

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/ketabchi/melli/api"
	"github.com/ketabchi/util"
)

type Book struct {
	url string
	doc *goquery.Document
}

var (
	reTranslators      = regexp.MustCompile(`(?:(?:[\[\(])?(?:\s)?(?:ترجمه(?:(?:\x{200c})+ی)?|مترجم(?:ان|ین)?)(?: \[?و (?:[\[\(])?(?:تنظیم|گردآوری|گردآورنده|سرپرستی|تدوین|تالیف|انطباق فرهنگی|ویرایش|بومی\x{200c}سازی|ترانه\x{200c}سرا|ترانه سرا|شعرهای|انتخاب|نگارش|ویراستار|بازآفرینی|بررسی|تحقیق|شرح)(?:[\]\)])?)?(?:\s)?(?:[\]\)])?)(.+?)(?:؛|\.|\]|$)`)
	reCleanPubDate     = regexp.MustCompile(`(\[.*\]|[,.]\s?c?\[?\d{4}\]?.?$)`)
	reCleanDoubleColon = regexp.MustCompile(`:[\s\x{200f}\x{202b}]+:`)
	reSerie            = regexp.MustCompile(`[^\.]+؛[۰-۹\s]+`)
	reNumber           = regexp.MustCompile(`[0-9۰-۹]`)

	ErrNoBook = errors.New("no book with this isbn")
)

func NewBookByISBN(isbn string, args ...string) (*Book, error) {
	url, err := api.GetBookURLByISBN(isbn, args...)
	if err != nil {
		return nil, err
	}
	if url == "" {
		return nil, ErrNoBook
	}

	return NewBook(url)
}

func NewBook(url string) (*Book, error) {
	res, err := api.Client.Get(url)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		return nil, err
	}

	return &Book{url: url, doc: doc}, nil
}

func (b *Book) Name() (name string) {
	if text := b.getField("\u200fعنوان و نام پديدآور"); text != "" {
		return b.nameFromField(text)
	}

	return ""
}

func (b *Book) nameFromField(text string) string {
	splited := strings.Split(text, "/")
	name := strings.Replace(splited[0], "[کتاب]", "", 1)
	name = util.Clean(name)

	return name
}

func (b *Book) Publisher() (publisher string) {
	if text := b.getField("\u200fمشخصات نشر"); text != "" {
		return b.publisherFromField(text)
	}

	return ""
}

func (b *Book) publisherFromField(text string) string {
	text = strings.ReplaceAll(text, "٬", "،")
	text = strings.ReplaceAll(text, "؛", "،")
	text = reCleanDoubleColon.ReplaceAllString(text, ":")
	splited := strings.Split(text, ":")
	if len(splited) < 2 {
		return ""
	}
	splited = strings.Split(splited[1], "،")
	if len(splited) == 0 {
		return ""
	}
	name := util.Clean(splited[0])
	name = strings.TrimPrefix(name, "نشر ")
	name = strings.TrimPrefix(name, "انتشارات ")

	return name
}

func (b *Book) Author() (faName string, enName string) {
	if text := b.getField("\u200fسرشناسه"); text != "" {
		splited := strings.Split(text, "\n")

		faName = b.authorFromField(splited[0])
		if len(splited) > 1 {
			enName = b.authorEnFromField(splited[1])
		}
	}

	return
}

func (b *Book) authorFromField(text string) string {
	text = strings.ReplaceAll(text, "٬", "،")
	text = strings.ReplaceAll(text, "؛", "،")
	splited := strings.Split(text, "،")

	return b.authorFullName(splited)
}

func (b *Book) authorEnFromField(text string) string {
	splited := strings.Split(text, ",")

	return b.authorFullName(splited)
}

func (b *Book) authorFullName(splited []string) string {
	if len(splited) < 2 {
		return ""
	}

	fn := util.Clean(splited[1])
	if reNumber.MatchString(fn) {
		fn = ""
	}

	ln := util.Clean(splited[0])
	if reNumber.MatchString(ln) {
		ln = ""
	}

	name := fmt.Sprintf("%s %s", fn, ln)
	name = strings.TrimSpace(name)

	return name
}

func (b *Book) OriginalName() (name string) {
	if text := b.getField("\u200fيادداشت"); text != "" {
		if !strings.Contains(text, "عنوان اصلی:") {
			return ""
		}

		text = strings.ReplaceAll(text, "عنوان اصلی:", "")
		text = util.Clean(text)
		text = strings.ReplaceAll(text, "\u202d", "")
		text = strings.ReplaceAll(text, "\u200e", "")
		text = reCleanPubDate.ReplaceAllString(text, "")
		text = strings.ReplaceAll(text, ",", "")

		name = strings.Trim(text, ".[] ")
	}

	return
}

func (b *Book) Translators() []string {
	if text := b.getField("\u200fعنوان و نام پديدآور"); text != "" {
		return b.translatorsFromField(text)
	}

	return []string{}
}

// TODO: samples we can't currently parse:
// http://opac.nlai.ir/opac-prod/bibliographic/3015099
// http://opac.nlai.ir/opac-prod/bibliographic/4312607
// http://opac.nlai.ir/opac-prod/bibliographic/4758204
// http://opac.nlai.ir/opac-prod/bibliographic/1595981
func (b *Book) translatorsFromField(text string) []string {
	ss := strings.Split(text, "/")
	text = ss[len(ss)-1]
	text = util.Clean(text)

	translators := make([]string, 0)
	ss = reTranslators.FindStringSubmatch(text)
	if len(ss) < 2 {
		return translators
	}

	text = strings.ReplaceAll(ss[1], " و ", "،")
	text = strings.ReplaceAll(text, "٬", "،")
	text = strings.ReplaceAll(text, "؛", "،")

	ss = strings.Split(text, "،")

	for _, s := range ss {
		s = util.Clean(s)
		if len(s) != 0 && !strings.Contains(s, "[") && !strings.Contains(s, "]") {
			translators = append(translators, s)
		}
	}

	return translators
}

func (b *Book) ISBN() (isbn string) {
	if text := b.getField("\u200f‏شابک"); text != "" {
		return b.isbnFromField(text)
	}

	return ""
}

func (b *Book) isbnFromField(text string) string {
	// TODO: needs more parsing text, fails on:
	// http://opac.nlai.ir/opac-prod/bibliographic/2072242
	return strings.ReplaceAll(text, "-", "")
}

func (b *Book) Link() string {
	return b.url
}

func (b *Book) Series() (ss []string) {
	if text := b.getField("\u200fفروست"); text != "" {
		return b.seriesFromField(text)
	}

	return []string{}
}

func (b *Book) seriesFromField(text string) []string {
	series := make([]string, 0)

	ss := reSerie.FindAllString(text, -1)
	if len(ss) == 0 {
		ss = append(ss, text)
	}
	for _, s := range ss {
		ss2 := strings.Split(s, "؛")
		s = ss2[0]
		s = strings.TrimSpace(s)
		s = strings.Replace(s, "\n", " ", -1)
		s = util.Clean(s)

		series = append(series, strings.TrimSuffix(s, "."))
	}

	return series
}

func (b *Book) getField(key string) (ret string) {
	b.doc.Find("td").EachWithBreak(func(i int, sel *goquery.Selection) bool {
		if sel.Text() == key {
			ret = sel.Next().Next().Text()
			return false
		}
		return true
	})

	return
}
