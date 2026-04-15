package gis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GetIdByGuidStruct struct {
	Id     int `json:"id"`
	Result struct {
		Id string `json:"id"`
	} `json:"result"`
	Jsonrpc string `json:"jsonrpc"`
}

func GetIdByGuid(id string, regId string) (*GetIdByGuidStruct, error) {
	fmt.Println("Starting to send GUID...")

	client := &http.Client{}
	//61723327-1c20-42fe-8dfa-402638d9b396
	data := []byte(fmt.Sprintf(`{"id":1,"method":"searchController.houseSearch","params":{"region_fias_id":"%s","house_fias_id":"%s"},"jsonrpc":"2.0"}`, regId, id))
	req, err := http.NewRequest(
		"POST", "https://portal.dom.gosuslugi.ru/api", bytes.NewBuffer(data),
	)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// добавляем заголовки
	req.Header.Add("Session-Guid", "79eb087b-35b5-4207-ace0-3ac786df3b6a")
	req.Header.Add("Request-Guid", "2ce9e7ed-fe05-4155-b05f-19656af79d7c")
	req.Header.Add("Content-Type", "application/json")

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

	ans := GetIdByGuidStruct{}
	json.Unmarshal(data, &ans)

	fmt.Println("Requested ID: " + ans.Result.Id)

	return &ans, nil
}
