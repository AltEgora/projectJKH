package gis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type UKContacts struct {
	Phones []string
	Emails []string
}

func GetUKContacts(id string) (*UK, error) {
	fmt.Println("Starting to get UK contacts")

	client := &http.Client{}
	data := []byte(fmt.Sprintf(`{"id":1,"method":"houseController.getHouseManagementOrganization","params":{"id":"%s"},"jsonrpc":"2.0"}`, id))
	req, err := http.NewRequest(
		"POST", "https://portal.dom.gosuslugi.ru/api", bytes.NewBuffer(data),
	)

	fmt.Println(req.Body)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// добавляем заголовки
	//req.Header.Add("Session-Guid", "79eb087b-35b5-4207-ace0-3ac786df3b6a")
	//req.Header.Add("Request-Guid", "2ce9e7ed-fe05-4155-b05f-19656af79d7c")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(data)))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	fmt.Println("Respons status: " + resp.Status)

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		return nil, err
	}

	var result map[string]interface{}
	json.Unmarshal(data, &result)

	contacts := result["result"].(map[string]interface{})["managementOrganization"].(map[string]interface{})["contacts"].(map[string]interface{})
	//fmt.Println(contacts["phones"].([]interface{})[0])
	ans := UK{}

	for _, i := range contacts["phones"].([]interface{}) {
		ans.Phones = append(ans.Phones, i.(map[string]interface{})["value"].(string))
	}

	for _, i := range contacts["emails"].([]interface{}) {
		ans.Emails = append(ans.Emails, i.(map[string]interface{})["value"].(string))
	}

	ans.Addr = result["result"].(map[string]interface{})["managementOrganization"].(map[string]interface{})["orgAddress"].(string)
	ans.Name = result["result"].(map[string]interface{})["managementOrganization"].(map[string]interface{})["shortName"].(string)

	fmt.Println("Results:")
	fmt.Print(ans)

	return &ans, nil
}
