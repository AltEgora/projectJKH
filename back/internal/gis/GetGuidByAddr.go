package gis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func GetGuidByAddr(addr string) (string, string, error) {
	fmt.Println("Starting to send GUID...")

	client := &http.Client{}
	data := []byte(fmt.Sprintf(`{"appName":"ГИС ЖКХ","reqId":"ГИС ЖКХ","rmFedCities":true,"woFlat":true,"addr":"%s"}`, addr))
	req, err := http.NewRequest(
		"POST", "https://portal.dom.gosuslugi.ru/gispa/suggest/api/v4_6", bytes.NewBuffer(data),
	)

	fmt.Println(req.Body)

	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	// добавляем заголовки
	//req.Header.Add("Session-Guid", "79eb087b-35b5-4207-ace0-3ac786df3b6a")
	//req.Header.Add("Request-Guid", "2ce9e7ed-fe05-4155-b05f-19656af79d7c")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(data)))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	defer resp.Body.Close()

	fmt.Println("Respons status: " + resp.Status)

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		return "", "", err
	}

	var result map[string]interface{}
	json.Unmarshal(data, &result)

	var ans string
	var ansReg string

	if len(result["addr"].([]interface{})) > 0 {
		ans = result["addr"].([]interface{})[0].(map[string]interface{})["addressGuid"].(string)
		ansReg = result["addr"].([]interface{})[0].(map[string]interface{})["elements"].([]interface{})[1].(map[string]interface{})["guid"].(string)
	} else {
		return "", "", fmt.Errorf("This address doesnot exist")
	}

	fmt.Println("Requested GUID: " + ans)
	fmt.Println("Requested Region GUID: " + ansReg)

	return ans, ansReg, nil
}
