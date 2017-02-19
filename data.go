package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/speps/go-hashids"
	"os"
	"strings"
)

var HashSalt = "H3ll0:H3ll0!@@#$"

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

func get_url_data(id int64) urls {
	urls_row := urls{}

	CON, err := connect_db()
	if err != nil {
		return urls_row
	}

	err = CON.QueryRow("select id,full_url from urls where id=?", id).Scan(&urls_row.Id, &urls_row.Full_url)

	// switch {
	// case err == sql.ErrNoRows:
	// 	res, err := CON.Exec("insert into urls set full_url=?", url)
	// 	if err != nil {
	// 		return 0
	// 	} else {
	// 		id, err := res.LastInsertId()
	// 		if err != nil {
	// 			return 0
	// 		}
	// 		short_id = id
	// 	}
	// case err != nil:
	// 	return 0
	// default:
	// 	short_id = urls_row.Id
	// }

	return urls_row
}

func save_new_url(url string) int64 {
	CON, err := connect_db()
	if err != nil {
		return 0
	}

	url = strings.Trim(url, " ")
	var urls_row urls
	var short_id int64
	err = CON.QueryRow("select id,full_url from urls where full_url=?", url).Scan(&urls_row.Id, &urls_row.Full_url)
	short_id = 0
	switch {
	case err == sql.ErrNoRows:
		res, err := CON.Exec("insert into urls set full_url=?", url)
		if err != nil {
			return 0
		} else {
			id, err := res.LastInsertId()
			if err != nil {
				return 0
			}
			short_id = id
		}
	case err != nil:
		return 0
	default:
		short_id = urls_row.Id
	}

	return short_id
}

func get_short_url(url string) string {
	url_id := save_new_url(url)

	if url_id <= 0 {
		return ""
	}

	hd := hashids.NewData()
	hd.Salt = HashSalt
	hd.MinLength = 3

	h := hashids.NewWithData(hd)
	id, err := h.EncodeInt64([]int64{url_id})
	if err != nil {
		return ""
	}
	return id
}

func get_full_url(short_id string) string {
	hd := hashids.NewData()
	hd.Salt = HashSalt
	hd.MinLength = 3

	h := hashids.NewWithData(hd)
	ids, err := h.DecodeInt64WithError(short_id)
	if err != nil || len(ids) <= 0 {
		return ""
	}
	id := ids[0]
	url_data := get_url_data(id)
	if url_data.Full_url != "" {
		return url_data.Full_url
	}
	return ""
}
