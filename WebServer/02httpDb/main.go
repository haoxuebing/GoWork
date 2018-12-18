package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:Admin!1234@tcp(10.99.2.153:3306)/xes_nrcp?charset=utf8")
	defer db.Close()
	checkErr(err)

	//插入数据
	// insert(21, db)

	//删除数据
	// delete(14, db)

	//更新数据
	// update(15, "设计部门", db)

	//查询数据
	get(db)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func get(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM aaa_test")
	cloumns, _ := rows.Columns()

	//获取列字段
	for _, cloumn := range cloumns {
		fmt.Print(cloumn + " ")
	}
	fmt.Println()

	for rows.Next() {

		rows.Scan(&cloumns[0], &cloumns[1], &cloumns[2], &cloumns[3])
		fmt.Println(cloumns[0], cloumns[1], cloumns[2], cloumns[3])
		// var id int
		// var chineseWord string
		// rows.Scan(&id, &chineseWord)
		// fmt.Println(id)
		// fmt.Println(chineseWord)
	}
}

func insert(id int, db *sql.DB) {
	//插入数据
	// stmt, err := db.Prepare("INSERT aaa_test SET id=?,nameA=?,descA=?")
	// res, _ := stmt.Exec(17, "研发部门", "2012-12-09", time.Now())
	// id, _ := res.LastInsertId()
	// fmt.Println(id)

	// stmt, _ := db.Prepare("INSERT aaa_test (id,nameA,descA,createTime)values (?,?,?,?)")
	// res, _ := stmt.Exec(17, "研发部门", "2012-12-09", time.Now())
	// id, _ := res.LastInsertId()
	// fmt.Println(id)

	// res, _ := db.Exec("INSERT aaa_test (id,nameA,descA)values (?,?,?)", 18, "研发部门", "2012-12-09")
	// id, _ := res.LastInsertId()
	// fmt.Println(id)

	//最简洁插入
	// db.Exec("INSERT aaa_test (id,nameA,descA,createTime)values (?,?,?,?)", 14, "qwe", "2012-12-09", time.Now())
	//or
	db.Exec("INSERT aaa_test SET id=?,nameA=?,descA=?,createTime=?", id, "研发部门", "写写代码", time.Now())

}

func delete(id int, db *sql.DB) {
	stmt, err := db.Prepare("delete from aaa_test where id=?")
	checkErr(err)

	res, err := stmt.Exec(id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println("删除数据影响行数：" + strconv.FormatInt(affect, 10))

}

func update(id int, name string, db *sql.DB) {
	stmt, err := db.Prepare("update aaa_test set nameA=? where id=?")
	checkErr(err)

	res, err := stmt.Exec(name, id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}
