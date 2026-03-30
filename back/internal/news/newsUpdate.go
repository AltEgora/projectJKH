package news

/*
func UpdateNews() []domain.New {
	var mas []domain.New

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

			mas = append(mas, domain.New{
				Header:      header,
				URL:         href,
				DateUpdated: time.Now().Format("YYYY-MM-DD"),
			})
			//fmt.Printf("%s - %s\n", header, href)
		})

	return mas
}

func StartUpdate(repo domain.NewRepository) {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		fmt.Println("Updating")
		for _, new := range UpdateNews() {
			repo.Create(context.Background(), new)
		}
	}

}
*/
