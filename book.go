package melli

import (
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
	translatorRe = regexp.MustCompile(`(?:(?:[\[\(])?(?:ترجمه(?:‌ی| و تنظیم| و (?:[\[\(])?گردآوری(?:[\]\)])?| و سرپرستی)?|مترجم(?:ان)?)(?:[\]\)])?)(.+?)(?:؛|\.|$)`)

	cleanPubDateRe     = regexp.MustCompile("(\\[.*\\]|[,.]\\s?c?\\d{4}.?$)")
	cleanDoubleColonRe = regexp.MustCompile(":[\\s\\x{200f}\\x{202b}]+:")
)

func NewBookByISBN(isbn string) (*Book, error) {
	url, err := api.GetBookURLByISBN(isbn)
	if err != nil {
		return nil, err
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
	name := util.Clean(splited[0])

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
	text = strings.Replace(text, "٬", "،", -1)
	text = strings.Replace(text, "؛", "،", -1)
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
	text = strings.Replace(text, "٬", "،", -1)
	text = strings.Replace(text, "؛", "،", -1)
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

	return name
}

func (b *Book) OriginalName() (name string) {
	b.doc.Find("td").EachWithBreak(func(i int, sel *goquery.Selection) bool {
		if sel.Text() == "‏يادداشت" {
			text := sel.Next().Next().Text()
			if !strings.Contains(text, "عنوان اصلی:") {
				return true
			}

			text = strings.Replace(text, "عنوان اصلی:", "", -1)
			text = util.Clean(text)
			text = strings.Replace(text, "\u202d", "", -1)
			text = strings.Replace(text, "\u200e", "", -1)
			text = cleanPubDateRe.ReplaceAllString(text, "")
			text = strings.Replace(text, ",", "", -1)

			name = strings.Trim(text, ".")

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
	translators := make([]string, 0)
	ss := translatorRe.FindStringSubmatch(text)
	if len(ss) < 2 {
		return translators
	}

	text = ss[1]
	if strings.Contains(text, " و ") {
		ss = strings.Split(ss[1], " و ")
	} else {
		ss = strings.Split(ss[1], "،")
	}

	for _, s := range ss {
		s = util.Clean(s)
		if len(s) != 0 {
			translators = append(translators, s)
		}
	}

	return translators
}
