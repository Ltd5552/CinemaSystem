package main

import (
	"database/sql"
	"fmt"
)

//1.管理员可以删除场次信息
//2.管理员可以删除电影信息

func deleteScreenings(num string) {
	if !haveScreenings(num) {
		fmt.Printf("Screening %v is not exist,delete failed!\n", num)
		return
	}
	sqlStr := "DELETE FROM screenings WHERE screeningNum = ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Println("Close err", err)
		}
	}(stmt)
	_, err = stmt.Exec(num)
	if err != nil {
		fmt.Printf("Delete failed.\n%v\n", err)
		return
	}
	fmt.Printf("Delete %v successfully.\n", num)
}

func deleteMovies(name string) {
	if !haveMovies(name) {
		fmt.Printf("Movie %v is not exist,delete failed!\n", name)
		return
	}
	sqlStr := "DELETE FROM movie WHERE movieTitle = ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Println("Close err", err)
		}
	}(stmt)
	_, err = stmt.Exec(name)
	if err != nil {
		fmt.Printf("Delete failed.\n%v\n", err)
		return
	}
	fmt.Printf("Delete %v successfully.\n", name)
}
