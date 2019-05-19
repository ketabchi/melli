package api

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/antzucaro/matchr"
	"github.com/ketabchi/util"
)

var client *http.Client

func GetBookURLByISBN(isbn string, args ...string) (string, error) {
	searchURL := fmt.Sprintf("http://opac.nlai.ir/opac-prod/search/bibliographicSimpleSearchProcess.do?simpleSearch.value=%s&bibliographicLimitQueryBuilder.biblioDocType=BF&simpleSearch.indexFieldId=221091&command=I&simpleSearch.tokenized=true&classType=0", isbn)
	doc, err := goquery.NewDocument(searchURL)
	if err != nil {
		return "", err
	}

	link, exists := doc.Find("#td2 > a").Attr("href")
	if len(args) > 0 {
		score := 0.0
		exists = false
		arg := util.Clean(args[0])
		doc.Find("#td2 > a").Each(func(i int, sel *goquery.Selection) {
			title := util.Clean(sel.Text())
			tmp := matchr.SmithWaterman(arg, title)
			tmp /= float64(len(arg))
			if tmp > score && (tmp > 0.1 || strings.Contains(arg, title)) {
				link, exists = sel.Attr("href")
				score = tmp
			}
		})
	}

	if !exists {
		return "", nil
	}

	u, err := url.Parse(link)
	if err != nil {
		return "", err
	}
	params, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return "", err
	}

	if id, exists := params["id"]; !exists {
		return "", fmt.Errorf("can't find book id in search page book link %s", link)
	} else {
		return fmt.Sprintf("http://opac.nlai.ir/opac-prod/bibliographic/%s", id[0]), nil
	}
}

func init() {
	client = &http.Client{}
}
