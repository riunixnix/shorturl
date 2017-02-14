package main

import (
	"encoding/json"
	//"errors"
	"fmt"
	//"io/ioutil"
	"net/http"
	//"regexp"
)

type url struct {
	Id  int
	Url string
}

type data_req_struct struct {
	Url string
}

func process_handler(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var data_req data_req_struct
	err := decoder.Decode(&data_req)
	if is_error(res, req, err) {
		fmt.Println("aaa=" + err.Error())
		return
	}

	fmt.Println("hello" + data_req.Url)

}

func is_error(res http.ResponseWriter, req *http.Request, err error) bool {
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return true
	}
	return false
}

func main() {
	http.HandleFunc("/shorten/", process_handler)
	http.ListenAndServe(":9090", nil)
}
