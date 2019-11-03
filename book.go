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
	translatorRe       = regexp.MustCompile(`(?:(?:[\[\(])?(?:\s)?(?:ترجمه(?:(?:\x{200c})+ی)?|مترجم(?:ان|ین)?)(?: و (?:[\[\(])?(?:تنظیم|گردآوری|گردآورنده|سرپرستی|تدوین|تالیف|انطباق فرهنگی|ویرایش|بومی‌سازی|ترانه‌سرا|ترانه سرا|شعرهای|انتخاب|نگارش|ویراستار|بازآفرینی|بررسی)(?:[\]\)])?)?(?:\s)?(?:[\]\)])?)(.+?)(?:؛|\.|\]|$)`)
	cleanPubDateRe     = regexp.MustCompile(`(\[.*\]|[,.]\s?c?\[?\d{4}\]?.?$)`)
	cleanDoubleColonRe = regexp.MustCompile(`:[\s\x{200f}\x{202b}]+:`)

	NoBookErr = errors.New("no book with this isbn")
)

func NewBookByISBN(isbn string, args ...string) (*Book, error) {
	url, err := api.GetBookURLByISBN(isbn, args...)
	if err != nil {
		return nil, err
	}
	if url == "" {
		return nil, NoBookErr
	}

	return NewBook(url)
}

func NewBook(url string) (*Book, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}

	return &Book{url: url, doc: doc}, nil
}

func (b *Book) Name() (name string) {
	b.doc.Find("td").EachWithBreak(func(i int, sel *goquery.Selection) bool {
		if sel.Text() == "‏عنوان و نام پديدآور" {
			text := sel.Next().Next().Text()
			name = b.nameFromField(text)
			return false
		}
		return true
	})

	return
}

func (b *Book) nameFromField(text string) string {
	splited := strings.Split(text, "/")
	name := strings.Replace(splited[0], "[کتاب]", "", 1)
	name = util.Clean(name)

	return name
}

func (b *Book) Publisher() (publisher string) {
	b.doc.Find("td").EachWithBreak(func(i int, sel *goquery.Selection) bool {
		if sel.Text() == "‏مشخصات نشر" {
			text := sel.Next().Next().Text()
			publisher = b.publisherFromField(text)
			return false
		}
		return true
	})

	return
}

func (b *Book) publisherFromField(text string) string {
	text = strings.ReplaceAll(text, "٬", "،")
	text = strings.ReplaceAll(text, "؛", "،")
	text = cleanDoubleColonRe.ReplaceAllString(text, ":")
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
	b.doc.Find("td").EachWithBreak(func(i int, sel *goquery.Selection) bool {
		if sel.Text() == "‏سرشناسه" {
			text := sel.Next().Next().Text()
			splited := strings.Split(text, "\n")

			faName = b.authorFromField(splited[0])
			if len(splited) > 1 {
				enName = b.authorEnFromField(splited[1])
			}
			return false
		}
		return true
	})

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
	ln := util.Clean(splited[0])
	name := fmt.Sprintf("%s %s", fn, ln)
	name = strings.TrimSpace(name)

	return name
}

func (b *Book) OriginalName() (name string) {
	b.doc.Find("td").EachWithBreak(func(i int, sel *goquery.Selection) bool {
		if sel.Text() == "‏يادداشت" {
			text := sel.Next().Next().Text()
			if !strings.Contains(text, "عنوان اصلی:") {
				return true
			}

			text = strings.ReplaceAll(text, "عنوان اصلی:", "")
			text = util.Clean(text)
			text = strings.ReplaceAll(text, "\u202d", "")
			text = strings.ReplaceAll(text, "\u200e", "")
			text = cleanPubDateRe.ReplaceAllString(text, "")
			text = strings.ReplaceAll(text, ",", "")

			name = strings.Trim(text, ".[] ")

			return false
		}
		return true
	})

	return
}

func (b *Book) Translators() []string {
	translators := make([]string, 0)
	b.doc.Find("td").EachWithBreak(func(i int, sel *goquery.Selection) bool {
		if sel.Text() == "‏عنوان و نام پديدآور" {
			text := sel.Next().Next().Text()
			translators = b.translatorsFromField(text)
			return false
		}
		return true
	})

	return translators
}

// TODO: samples we can't currently parse:
// http://opac.nlai.ir/opac-prod/bibliographic/3015099
// http://opac.nlai.ir/opac-prod/bibliographic/4312607
func (b *Book) translatorsFromField(text string) []string {
	ss := strings.Split(text, "/")
	text = ss[len(ss)-1]
	text = util.Clean(text)

	translators := make([]string, 0)
	ss = translatorRe.FindStringSubmatch(text)
	if len(ss) < 2 {
		return translators
	}

	text = strings.ReplaceAll(ss[1], " و ", "،")
	text = strings.ReplaceAll(text, "٬", "،")
	text = strings.ReplaceAll(text, "؛", "،")

	ss = strings.Split(text, "،")

	for _, s := range ss {
		s = util.Clean(s)
		if len(s) != 0 {
			translators = append(translators, s)
		}
	}

	return translators
}

func (b *Book) ISBN() (isbn string) {
	b.doc.Find("td").EachWithBreak(func(i int, sel *goquery.Selection) bool {
		if sel.Text() == "‏‏شابک" {
			text := sel.Next().Next().Text()
			isbn = b.isbnFromField(text)
			return false
		}
		return true
	})

	return
}

func (b *Book) isbnFromField(text string) string {
	// TODO: needs more parsing text, fails on:
	// http://opac.nlai.ir/opac-prod/bibliographic/2072242
	return strings.ReplaceAll(text, "-", "")
}
