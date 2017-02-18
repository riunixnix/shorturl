package main

import (
	"fmt"
	"net/http"
	"strings"
)

func get_base_url(req *http.Request) string {
	isHTTPS := req.TLS != nil
	scheme := ""
	if isHTTPS == true {
		scheme = "https://"
	} else {
		scheme = "http://"
	}

	return scheme + req.Host
}

func is_alpha_numeric(str string) bool {
	f := func(r rune) bool {
		return (r < 'A' || r > 'Z') && (r < 'a' || r > 'z') && (r < '0' || r > '9')
	}

	return strings.IndexFunc(str, f) == -1
}

func error_func(error_type int, res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(error_type)
	txt_status := ""
	switch error_type {
	case http.StatusNotFound:
		txt_status = "Page not found"
	case http.StatusBadRequest:
		txt_status = "Bad Request"
	default:
		txt_status = "Unknown Error"
	}
	fmt.Fprint(res, txt_status)
}
