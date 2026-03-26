package news

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	UpdateNews()
}

type New struct {
	Header string
	Href   string
}

func UpdateNews() []New {
	var mas []New

	resp, err := http.Get("https://tverigrad.ru/category/obshhestvo/main-030/")

	if err != nil {
		fmt.Println("Error while connecting to news site")
	}

	defer resp.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(resp.Body)

	doc.Find("[class^='row block pad1em cat_item item']").Each(
		func(i int, s *goquery.Selection) {
			a := s.Find("h3 a")

			header := a.Text()
			href, _ := a.Attr("href")

			mas = append(mas, New{
				Header: header,
				Href:   href,
			})
			//fmt.Printf("%s - %s\n", header, href)
		})

	return mas
}
