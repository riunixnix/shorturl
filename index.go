package main

import (
	"encoding/json"
	//"errors"
	//"fmt"
	//"io/ioutil"
	"net/http"
	//"regexp"
)

type url struct {
	Id  int
	Url string
}

type data_req_struct struct {
	Url string `json:"url"`
}

type data_res_struct struct {
	Short string `json:"short"`
}

func process_handler(res http.ResponseWriter, req *http.Request) {

	var data_req data_req_struct
	err := json.NewDecoder(req.Body).Decode(&data_req)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	//processing part
	short_url := data_req.Url

	//return result
	data_res := data_res_struct{Short: short_url}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(data_res)
}

func main() {
	http.HandleFunc("/shorten/", process_handler)
	http.ListenAndServe(":9090", nil)
}
