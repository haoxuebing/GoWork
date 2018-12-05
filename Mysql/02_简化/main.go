package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, _ := sql.Open("mysql", "root:Admin!1234@tcp(10.99.2.153:3306)/xes_nrcp")
	defer db.Close()

	rows, _ := db.Query("select * from xes_rw_words;")
	defer rows.Close()

	cloumns, _ := rows.Columns()
	values := make([]sql.RawBytes, len(cloumns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		rows.Scan(scanArgs...)
		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(cloumns[i], ": ", value)
		}
		fmt.Println("------------------")
	}

}
