package utils

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	xmlpath "gopkg.in/xmlpath.v2"
	"mvdan.cc/xurls"
)

type TabelogResult struct {
	Address string
	URL     string
	Name    string
	Genre   string
}

func RegexTabelogURL(post string) string {
	URL := xurls.Relaxed().FindString(post)
	return URL
}

func GetAddressFromTabelogURL(URL string) TabelogResult {
	body, _ := HttpGet(URL)

	address := xpath(body, `//*[@id="contents-rstdata"]/div[2]/table[1]/tbody/tr[5]/td/p`)
	name := xpath(body, `//*[@id="rstdtl-head"]/div[1]/div[1]/div[1]/div[1]/div/h2/a/span`)
	genre := xpath(body, `//*[@id="contents-rstdata"]/div[2]/table[1]/tbody/tr[2]/td/span`)

	return TabelogResult{
		Address: address,
		URL:     URL,
		Name:    name,
		Genre:   genre,
	}
}

func GetInfoFromTabelogURL(URL string) TabelogResult {
	doc, err := goquery.NewDocument(URL)
	if err != nil {
		fmt.Println("error happend while requesting to %s\n")
	}

	address := findClassText(doc, ".rstinfo-table__address")
	name := findClassTDText(doc, ".rstinfo-table__table th", "店名")
	genre := findClassTDText(doc, ".rstinfo-table__table th", "ジャンル")

	return TabelogResult{
		Address: address,
		URL:     URL,
		Name:    name,
		Genre:   genre,
	}
}

func findClassText(doc *goquery.Document, class string) string {
	var val = "<No Data>"
	doc.Find(class).Each(func(i int, s *goquery.Selection) {
		val = s.Text()
		val = strings.TrimSpace(val)
	})
	return val
}

func findClassTDText(doc *goquery.Document, selector string, th string) string {
	var td = "<No Data>"

	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		val := s.Text()
		if val == th {
			td = s.Next().Text()
			td = strings.TrimSpace(td)
		}
	})

	return td
}

func xpath(body []byte, xpath string) string {
	var val = "<No Data>"

	path := xmlpath.MustCompile(xpath)
	root, err := xmlpath.ParseHTML(bytes.NewReader(body))
	if err != nil {
		return val
	}

	iter := path.Iter(root)
	for iter.Next() {
		n := iter.Node()
		val = n.String()
		val = strings.TrimSpace(val)
	}

	return val
}
