package melli

import (
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
	name := splited[0]

	// TODO what to do with nimfasele های
	filter := func(r rune) rune {
		switch r {
		case 8205:
			return -1
		case 8204:
			return -1
		}
		return r
	}
	name = strings.Map(filter, name)

	return strings.Trim(name, " ")
}
