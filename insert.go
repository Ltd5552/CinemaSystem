package main

import (
	"database/sql"
	"fmt"
)

//插入函数
//1.系统可以插入个人信息
//2.管理员可以插入场次信息
//3.管理员可以插入电影信息

//当用户注册后执行
func insertPersonal(uid string, name string, password string) {
	sqlStr := "INSERT INTO user(uid,name,sex,birthday,location,phone, password) VALUES (?,?,?,?,?,?,?)"

	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Println("close err:", err)
		}
	}(stmt)
	_, err = stmt.Exec(uid, name, " ", " ", " ", " ", password)
	if err != nil {
		fmt.Printf("Insert failed.\n%v\n", err)
		return
	}
	fmt.Println("Insert success.")

}

func insertScreenings(screeningNum string, movieNum string, theaterNum string, showTime string, remainSeats int64) {

	sqlStr := "INSERT INTO screenings(screeningNum ,movieNum , theaterNum, showTime, remainSeats) VALUES (?,?,?,?,?)"

	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Println("close err:", err)
		}
	}(stmt)
	_, err = stmt.Exec(screeningNum, movieNum, theaterNum, showTime, remainSeats)
	if err != nil {
		fmt.Printf("Insert failed.\n%v\n", err)
		return
	}
	fmt.Println("Insert success.")

}

func insertMovies(movieNum string, movieTitle string, releaseDate string, duration float64) {

	sqlStr := "INSERT INTO movie(movieNum, movieTitle, releaseDate, duration,aveFilmScore ) VALUES (?,?,?,?,?)"

	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Println("close err:", err)
		}
	}(stmt)
	_, err = stmt.Exec(movieNum, movieTitle, releaseDate, duration, 0)
	if err != nil {
		fmt.Printf("Insert failed.\n%v\n", err)
		return
	}
	fmt.Println("Insert success.")
}
