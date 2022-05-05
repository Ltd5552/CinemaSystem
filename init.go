package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"
)

// 定义一个全局对象db
var db *sql.DB

type User struct {
	uid      string
	name     string
	sex      string
	birthday string
	location string
	phone    string
	password string
}

type Ticket struct {
	ticketNum    string
	price        float64
	screeningNum string
}

type Movie struct {
	movieNum     string
	movieTitle   string
	releaseDate  string
	duration     float64
	aveFilmScore float64
}

type TicketMachine struct {
	ticketMachineNum string
	remainingTickets int64
}

type Theater struct {
	theaterNum string
	cinemaNum  string
	seatsNum   int64
}

type Cinema struct {
	cinemaNum      string
	cinemaName     string
	city           string
	contact        string
	aveCinemaScore float64
}

type Evaluation struct {
	evaluationId string
	cinemaNum    string
	cinemaScore  int64
	movieNum     string
	filmScore    int64
}

type Screenings struct {
	screeningNum string
	movieNum     string
	theaterNum   string
	showTime     string
	remainSeats  int64
}

type Administrator struct {
}

// 初始化数据库的函数
func initDB() (err error) {
	// DSN:Data Source Name
	dsn := "root:lht12138@tcp(127.0.0.1:3306)/cinema?charset=utf8mb4&parseTime=True"
	// 这里只是赋值，不会校验账号密码是否正确
	// 这里不要使用:=由于是给全局变量db赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（这里会校验dsn内容是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

// RandStr 生成随机字符串,用于电影票编号
func RandStr() string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < 10; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

//判断是否还有座位,返回座位数
func haveSeat(num string) int64 {
	sqlStr := "SELECT remainSeats FROM screenings WHERE screeningNum = ?"
	query, err := db.Query(sqlStr, num)
	if err != nil {
		fmt.Println("Prepare failed,err:", err)
		return 0
	}
	var seats int64
	for query.Next() {
		err := query.Scan(&seats)
		if err != nil {
			return 0
		}
		if seats != 0 {
			return seats
		}
	}
	return 0
}

//判断是否存在用户
func haveUser(uid string) bool {
	sqlStr := "SELECT * FROM user WHERE uid = ?"
	query, err := db.Query(sqlStr, uid)
	if err != nil {
		fmt.Println("Prepare failed,err:", err)
		return false
	}
	if query.Next() {
		return true
	}
	return false
}

//判断是否存在电影
func haveMovies(name string) bool {
	sqlStr := "SELECT * FROM movie WHERE movieTitle = ?"
	//进行预处理，先将sql发送给mysql服务端
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("Prepare failed,err:", err)
		return false
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			return
		}
	}(stmt)

	//发送请求参数
	query, err := stmt.Query(name)
	if err != nil {
		fmt.Println("Query failed,err:", err)
		return false
	}
	if query.Next() {
		return true
	}
	return false
}

//判断是否存在场次
func haveScreenings(num string) bool {
	sqlStr := "SELECT * FROM screenings WHERE screeningNum = ?"
	//进行预处理，先将sql发送给mysql服务端
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("Prepare failed,err:", err)
		return false
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			return
		}
	}(stmt)

	//发送请求参数
	query, err := stmt.Query(num)
	if err != nil {
		fmt.Println("Query failed,err:", err)
		return false
	}
	if query.Next() {
		return true
	}
	return false
}
