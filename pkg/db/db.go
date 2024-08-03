package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "./virtbro.db")
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS hosts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		uri TEXT NOT NULL,
		uuid TEXT NOT NULL UNIQUE
	);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
	}
}

func AddHost(name, uri, uuid string) error {
	stmt, err := db.Prepare("INSERT INTO hosts(name, uri, uuid) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(name, uri, uuid)
	if err != nil {
		return err
	}
	return nil
}

func GetHostByID(id int) (string, error) {
	var uri string
	err := db.QueryRow("SELECT uri FROM hosts WHERE id = ?", id).Scan(&uri)
	if err != nil {
		return "", err
	}
	return uri, nil
}

func ListHosts() ([]map[string]string, error) {
	rows, err := db.Query("SELECT id, name, uri, uuid FROM hosts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hosts []map[string]string
	for rows.Next() {
		var id int
		var name, uri, uuid string
		err = rows.Scan(&id, &name, &uri, &uuid)
		if err != nil {
			return nil, err
		}
		hosts = append(hosts, map[string]string{
			"id":   fmt.Sprintf("%d", id),
			"name": name,
			"uri":  uri,
			"uuid": uuid,
		})
	}
	return hosts, nil
}

func CloseDB() {
	db.Close()
}
