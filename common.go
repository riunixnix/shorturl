package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"os"
	"strings"
)

type dbConfig struct {
	Host string
	User string
	Pass string
	Db   string
}

func load_db_conf(path string) dbConfig {
	file, err := os.Open(path + "/db.json")
	if err != nil {
		panic(err.Error())
	}
	var conf dbConfig
	err = json.NewDecoder(file).Decode(&conf)
	if err != nil {
		panic(err.Error())
	}
	return conf
}

func connect_db() (*sql.DB, error) {
	conf_path := os.Getenv("conf_path")
	conf := load_db_conf(conf_path)

	str_connect := conf.User + ":" + conf.Pass + "@tcp(" + conf.Host + ")/" + conf.Db
	fmt.Println(str_connect)
	db, err := sql.Open("mysql", str_connect)
	if err != nil {
		fmt.Println("err=" + err.Error())
		return db, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Db is not connected")
		return db, err
	}
	return db, nil
}

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
