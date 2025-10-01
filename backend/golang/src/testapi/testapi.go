package testapi

import (
	"database/sql"
	"fmt"
)

type TestObject struct {
	id    int
	title string
}

func ReadAll(db *sql.DB) {
	var objects []TestObject
	rows, err := db.Query("select * from test;")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		myObject := TestObject{}
		err = rows.Scan(&myObject.id, &myObject.title)
		if err != nil {
			panic(err)
		}
		objects = append(objects, myObject)
	}
	rows.Close()

	fmt.Println(objects)
}
