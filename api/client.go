package api

import "net/http"

var (
	client *http.Client
)

func getBookURLByISBN(isbn string) string {
}

func init() {
	client = &http.Client{}
}
