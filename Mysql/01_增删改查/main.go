package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:Admin!1234@tcp(10.99.2.153:3306)/xes_nrcp")
	if err != nil {
		fmt.Println("db err=", err)
	}
	defer db.Close()
	Get(db)

}
func Get(db *sql.DB) {
	rows, err := db.Query("select * from xes_rw_words;")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	cloumns, err := rows.Columns()
	if err != nil {
		fmt.Println(err)
	}
	// for rows.Next() {
	//  err := rows.Scan(&cloumns[0], &cloumns[1], &cloumns[2])
	//  if err != nil {
	//      fmt.Println(err)
	//  }
	//  fmt.Println(cloumns[0], cloumns[1], cloumns[2])
	// }
	values := make([]sql.RawBytes, len(cloumns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			fmt.Println(err)
		}
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
	if err = rows.Err(); err != nil {
		fmt.Println(err)
	}
}

// 插入数据
func Insert(db *sql.DB) {
	stmt, err := db.Prepare("INSERT INTO user(name, age) VALUES(?, ?);")
	if err != nil {
		fmt.Println(err)
	}
	res, err := stmt.Exec("python", 19)
	if err != nil {
		fmt.Println(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("ID=%d, affected=%d\n", lastId, rowCnt)
}

// 删除数据
func Delete(db *sql.DB) {
	stmt, err := db.Prepare("DELETE FROM user WHERE name='python'")
	if err != nil {
		fmt.Println(err)
	}
	res, err := stmt.Exec()
	lastId, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("ID=%d, affected=%d\n", lastId, rowCnt)
}

// 更新数据
func Update(db *sql.DB) {
	stmt, err := db.Prepare("UPDATE user SET age=27 WHERE name='python'")
	if err != nil {
		fmt.Println(err)
	}
	res, err := stmt.Exec()
	lastId, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("ID=%d, affected=%d\n", lastId, rowCnt)
}
