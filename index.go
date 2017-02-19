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

func redirect_handler(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path[1:]

	//----------------Validate-------------------
	if path == "" || is_alpha_numeric(path) != true {
		error_func(http.StatusNotFound, res, req)
		return
	}

	//------------Processing part-----------------
	url := get_full_url(path)
	if url == "" {
		error_func(http.StatusNotFound, res, req)
		return
	}

	//------------Redirect------------------------
	http.Redirect(res, req, url, http.StatusMovedPermanently)
}

func shorten_handler(res http.ResponseWriter, req *http.Request) {

	var data_req data_req_struct
	base_url := get_base_url(req)

	//----------------Validate-------------------
	err := json.NewDecoder(req.Body).Decode(&data_req)
	if err != nil || data_req.Url == "" {
		error_func(http.StatusBadRequest, res, req)
		return
	}

	//------------Processing part-----------------
	short_id := get_short_url(data_req.Url)
	short_url := base_url + "/" + short_id

	//------------Return result-------------------
	data_res := data_res_struct{Short: short_url}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(data_res)
}

func main() {
	http.HandleFunc("/shorten/", shorten_handler)
	http.HandleFunc("/", redirect_handler)
	http.ListenAndServe(":9090", nil)
}
