package mysql_study

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func F1() {
	fmt.Println("mysql study")
}

// 使用mysql驱动
//func Open(driverName, dataSourceName string) (*sql.DB, error)
// sql.DB维护了与mysql的所有连接信息 底层维护了一个连接池
// 方便查询 先创建一个表示表的结构体

type People struct {
	id     int
	name   string
	gender string
}

func open() (*sql.DB, error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("数据库连接失败")
		return nil, err
	}
	fmt.Println("数据库连接成功")
	return db, nil
}

// Query 查询单行
func Query() {
	db, err := open()
	if db == nil {
		return
	}
	sqlString := "select id, name, gender from people where id = ?"
	var p People
	fmt.Println("开始执行QueryRow()")
	err = db.QueryRow(sqlString, 1).Scan(&p.id, &p.name, &p.gender)
	if err != nil {
		fmt.Println("查询错误")
		fmt.Printf("查询错误: %s", err.Error())
		return
	}
	fmt.Println(p)
}

func QueryAll() {
	db, err := open()
	if db == nil {
		return
	}
	sqlString := "select * from people"
	rows, err := db.Query(sqlString)
	if err != nil {
		fmt.Println("查询出现问题", err.Error())
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("row关闭失败")
		}
	}(rows)
	if rows == nil {
		return
	}
	for rows.Next() {
		var p People
		err := rows.Scan(&p.id, &p.name, &p.gender)
		if err != nil {
			return
		}
		fmt.Println(p)
	}
}

// 增删改均使用Exec方法
// sql预编译

// Go中的事务
// 1. Begin
// 2. Commit
// 3.Rollback
