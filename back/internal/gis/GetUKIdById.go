package gis

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

type UKInfo struct {
	Id                 string `json:"id"`
	OrganizationRootId string `json:"organizationRole"`
	ShortName          string `json:"shrtName"`
	OrganizationRole   string `json:"organizationRootId"`
	Phone              string `json:"phone"`
}

func GetUKIdById(id string) (*UKInfo, error) {

	client := &http.Client{}
	data := []byte(`{"id":1,"method":"searchController.houseSearch","params":{"region_fias_id":"61723327-1c20-42fe-8dfa-402638d9b396","house_fias_id":"22b62592-3cad-4953-bafb-6232cdbf7d72"},"jsonrpc":"2.0"}`)
	req, err := http.NewRequest(
		"GET", "https://portal.dom.gosuslugi.ru/home/"+id, nil,
	)

	fmt.Println("Starting to get UKID on " + req.URL.Path)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// добавляем заголовки
	//req.Header.Add("Session-Guid", "79eb087b-35b5-4207-ace0-3ac786df3b6a")
	//req.Header.Add("Request-Guid", "2ce9e7ed-fe05-4155-b05f-19656af79d7c")
	//req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	//req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:149.0) Gecko/20100101 Firefox/149.0")
	//req.Header.Add("Pragma", "no-cache")
	//req.Header.Add("Accept-Encoding", "gzip, deflate, br, zstd")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("Request done, code: " + resp.Status)

	defer resp.Body.Close()

	fmt.Println("Respons status: " + resp.Status)

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		return nil, err
	}

	html := string(data)

	pattern := regexp.MustCompile(`"managementOrganization":{[^}]*}`)

	matches := strings.Split(pattern.FindString(html), "{")

	if len(matches) <= 1 {
		fmt.Println("Wrong response while GetUKIdById")
		for _, m := range matches {
			fmt.Println(m)
		}
		return nil, fmt.Errorf("Wrong response while GetUKIdById")
	}

	match := "{" + matches[1]

	ans := UKInfo{}
	json.Unmarshal([]byte(match), &ans)

	fmt.Println("Result:")
	fmt.Println(ans)

	return &ans, nil
}
