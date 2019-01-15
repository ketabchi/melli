package melli

import "github.com/PuerkitoBio/goquery"

type Book struct {
	url string
	doc *goquery.Document
}

func NewBookByISBN(isbn string) (*Book, error) {

}

func (b *Book) Name() string {

}
