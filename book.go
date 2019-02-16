package melli

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/ketabchi/melli/api"
)

type Book struct {
	url string
	doc *goquery.Document
}

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
	name := clean(splited[0])

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
	splited := strings.Split(text, ":")
	if len(splited) < 2 {
		return ""
	}
	splited = strings.Split(splited[1], "،")
	if len(splited) == 0 {
		return ""
	}
	name := clean(splited[0])
	name = strings.TrimPrefix(name, "نشر ")
	name = strings.TrimPrefix(name, "انتشارات ")

	return name
}

func (b *Book) Author() (name string, eName string) {
	b.doc.Find("td").EachWithBreak(func(i int, sel *goquery.Selection) bool {
		if sel.Text() == "‏سرشناسه" {
			text := sel.Next().Next().Text()
			splited := strings.Split(text, "\n")

			name = b.authorFromField(splited[0])
			if len(splited) > 1 {
				eName = b.authorEnFromField(splited[1])
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
	fn := clean(splited[1])
	ln := clean(splited[0])
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

			text = strings.Replace(text, "‏‫عنوان اصلی:", "", -1)
			text = clean(text)
			name = strings.Trim(text, ".")

			return false
		}
		return true
	})

	return
}
