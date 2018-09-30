package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	// 连接语句格式
	// user@unix(/path/to/socket)/dbname?charset=utf8
	// user:password@tcp(localhost:5555)/dbname?charset=utf8
	// user:password@/dbname
	// user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/go_test?charset=utf8")
	// 是否连接成功
	CheckErr(err)

	// 插入数据
	//db.Prepare() 函数用来返回准备要执行的 sql 操作，然后返回准备完毕的执行状态
	stmt, err := db.Prepare("INSERT userinfo SET username=?,department=?,created=?")
	CheckErr(err)

	// stmt.Exec() 函数用来执行 stmt 准备好的SQL语句
	res, err := stmt.Exec("tomo", "研发部", "2018-9-30")
	CheckErr(err)

	id, err := res.LastInsertId()
	CheckErr(err)

	fmt.Println("insert success, uid：", id)

	// 更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	CheckErr(err)

	res,err = stmt.Exec("tomoUP", id)
	CheckErr(err)

	affect, err := res.RowsAffected()
	CheckErr(err)

	fmt.Println("updated success, affect: ", affect)

	// 查询数据
	// db.Query() 函数用来直接执行 Sql 返回 Rows 结果
	rows, err := db.Query("select * from userinfo")
	CheckErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		CheckErr(err)
		fmt.Println("select info: ")
		fmt.Println("uid: ", uid)
		fmt.Println("username: ", username)
		fmt.Println("department: ", department)
		fmt.Println("created: ", created)
	}

	// 删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	CheckErr(err)

	res, err = stmt.Exec(id)
	CheckErr(err)

	affect, err = res.RowsAffected()
	CheckErr(err)

	fmt.Println("delete success, affect: ", affect)

	db.Close()
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
