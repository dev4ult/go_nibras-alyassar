package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type contacts struct {
	id 			int
	full_name 	string
	email 		string
	created_at 	string
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("root:@tcp(127.0.0.1:3306)/%s", "basic_connect_golang"))

	if err != nil {
		return nil, err
	}

	return db, nil
}

func FetchAll() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return 
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM contacts")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer rows.Close()

	var result []contacts

	for rows.Next() {
        var each = contacts{}
        var err = rows.Scan(&each.id, &each.full_name, &each.email, &each.created_at)

        if err != nil {
            fmt.Println(err.Error())
            return
        }

        result = append(result, each)
    }

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, data := range result {
		fmt.Println(data.full_name)
	}
}