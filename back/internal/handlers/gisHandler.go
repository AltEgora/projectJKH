package handlers

import (
	"ConsultantBack/internal/gis"
	"encoding/json"
	"fmt"
	"net/http"
)

type gisReqType struct {
	Addr string `json:"address"`
}

func GisHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Got request on gis page")

	var reqData gisReqType
	decoder := json.NewDecoder(req.Body)

	if err := decoder.Decode(&reqData); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	defer req.Body.Close()

	uk, err := gis.GetUK(reqData.Addr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(uk)
}
