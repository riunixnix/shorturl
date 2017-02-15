package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"reflect"
	"strings"
)

type dbConfig struct {
	Host string
	User string
	Pass string
	Db   string
}

type urls struct {
	Id       int64
	Full_url string
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
	fmt.Println(reflect.TypeOf(db))
	if err != nil {
		fmt.Println("err=" + err.Error())
		return db, err
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Db is not connected")
		return db, err
	}
	return db, nil
}

func save_new_url(url string) string {
	CON, err := connect_db()
	if err != nil {
		return ""
	}

	url = strings.Trim(url, " ")
	var urls_row urls
	var room_id int64
	err = CON.QueryRow("select id,full_url from urls where full_url=?", url).Scan(&urls_row.Full_url)
	room_id = 0
	switch {
	case err == sql.ErrNoRows:
		res, err := CON.Exec("insert into urls set full_url=?", url)
		if err != nil {

		} else {
			id, err := res.LastInsertId()
			if err != nil {
			}
			room_id = id
		}
	case err != nil:
		panic(err.Error())
	default:
		room_id = urls_row.Id
	}
	return string(room_id)
}
