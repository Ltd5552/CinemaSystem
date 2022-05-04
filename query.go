package main

import (
	"database/sql"
	"fmt"
)

//查询函数
//要求包含 1.查询个人信息(用户)
//       2.查询场次信息(用户、管理员)
//       3.查询电影信息(用户、管理员)
// 		 4.查询电影院信息(用户、管理员)
//       5.查询放映厅信息(用户、管理员)
//		 6.查询取票机剩余可印电影票数(管理员)
//       7.查询评论信息(管理员)

//查询个人信息，用户的uid作为传入参数，使用了预处理
func queryPersonal(UID string) {
	sqlStr := "SELECT * FROM detailuser WHERE uid = ?"
	//进行预处理，先将sql发送给mysql服务端
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("Prepare failed,err:", err)
		return
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			return
		}
	}(stmt)

	//发送请求参数
	query, err := stmt.Query(UID)
	if err != nil {
		fmt.Println("Query failed,err:", err)
		return
	}
	var u User
	for query.Next() {
		err := query.Scan(&u.uid, &u.name, &u.sex, &u.birthday, &u.location, &u.phone, &u.password)
		if err != nil {
			fmt.Println("Scan failed,err:", err)
			return
		}
		fmt.Printf("uid:%v name:%v sex:%v birthday:%v location:%v phone:%v password:%v\n", u.uid, u.name, u.sex, u.birthday, u.location, u.phone, u.password)
	}
}

//查询所有场次信息，直接查询
func queryScreenings() {
	sqlStr := "SELECT * FROM screenings"
	query, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println("Prepare failed,err:", err)
		return
	}
	var s Screenings
	for query.Next() {
		err := query.Scan(&s.screeningNum, &s.movieNum, &s.theaterNum, &s.showTime, &s.remainSeats)
		if err != nil {
			fmt.Println("Scan failed,err:", err)
			return
		}
		fmt.Printf("screeningNum:%v cinemaNum:%v theaterNum:%v showTime:%v remainSeats:%v \n", s.screeningNum, s.movieNum, s.theaterNum, s.showTime, s.remainSeats)
	}
}

//查询电影信息，电影名作为传入参数，使用了预处理
func queryMovies(name string) {
	sqlStr := "SELECT * FROM movie WHERE movieTitle = ?"
	//进行预处理，先将sql发送给mysql服务端
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("Prepare failed,err:", err)
		return
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
		return
	}
	var m Movie
	for query.Next() {
		err := query.Scan(&m.movieNum, &m.movieTitle, &m.releaseDate, &m.duration, &m.aveFilmScore)
		if err != nil {
			fmt.Println("Scan failed,err:", err)
			return
		}
		fmt.Printf("movieNum:%v movieTitle:%v releaseDate:%v duration:%v aveFilmScore:%v \n", m.movieNum, m.movieTitle, m.releaseDate, m.duration, m.aveFilmScore)
	}
}

//查询电影院信息，电影院名作为传入参数，使用了预处理
func queryCinema(name string) {
	sqlStr := "SELECT * FROM cinema WHERE cinemaName = ?"
	//进行预处理，先将sql发送给mysql服务端
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("Prepare failed,err:", err)
		return
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
		return
	}
	var c Cinema
	for query.Next() {
		err := query.Scan(&c.cinemaNum, &c.cinemaName, &c.city, &c.contact, &c.aveCinemaScore)
		if err != nil {
			fmt.Println("Scan failed,err:", err)
			return
		}
		fmt.Printf("cinemaNum:%v cinemaName:%v city:%v contact:%v aveCinemaScore:%v\n", c.cinemaNum, c.cinemaName, c.city, c.contact, c.aveCinemaScore)
	}
}

//查询放映厅信息，放映厅编号名作为传入参数，使用了预处理
func queryTheater(num string) {
	sqlStr := "SELECT * FROM theater WHERE theaterNum = ?"
	//进行预处理，先将sql发送给mysql服务端
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("Prepare failed,err:", err)
		return
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
		return
	}
	var t Theater
	for query.Next() {
		err := query.Scan(&t.theaterNum, &t.cinemaNum, &t.seatsNum)
		if err != nil {
			fmt.Println("Scan failed,err:", err)
			return
		}
		fmt.Printf("theaterNum:%v cinemaNum:%v seatsNum:%v\n", t.theaterNum, t.cinemaNum, t.seatsNum)
	}
}

//查询所有取票机剩余可印票数，使用了预处理
func queryTicketMachine() {
	sqlStr := "SELECT * FROM ticketmachine"
	//进行预处理，先将sql发送给mysql服务端
	query, err := db.Query(sqlStr)
	if err != nil {
		return
	}
	var t TicketMachine
	for query.Next() {
		err := query.Scan(&t.ticketMachineNum, &t.remainingTickets)
		if err != nil {
			fmt.Println("Scan failed,err:", err)
			return
		}
		fmt.Printf("ticketMachineNum:%v remainingTickets:%v\n", t.ticketMachineNum, t.remainingTickets)
	}
}

//查询所有评论信息，直接查询
func queryEvaluation() {
	sqlStr := "SELECT * FROM evaluation"
	query, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println("Prepare failed,err:", err)
		return
	}
	var e Evaluation
	for query.Next() {
		err := query.Scan(&e.evaluationId, &e.cinemaScore, &e.filmScore)
		if err != nil {
			fmt.Println("Scan failed,err:", err)
			return
		}
		fmt.Printf("evaluationId:%v cinemaScore:%v filmScore:%v\n", e.evaluationId, e.cinemaScore, e.filmScore)
	}
}

//通过电影名查询场次信息，等值连接查询
func queryScreeningsByMovie(name string) {
	sqlStr := `SELECT screeningNum, screenings.movieNum, theaterNum, showTime, remainSeats
		        FROM screenings,movie
		        WHERE screenings.movieNum = movie.movieNum AND 
		              movieTitle = ?`
	//进行预处理，先将sql发送给mysql服务端
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("Prepare failed,err:", err)
		return
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
		return
	}
	var s Screenings
	for query.Next() {
		err := query.Scan(&s.screeningNum, &s.movieNum, &s.theaterNum, &s.showTime, &s.remainSeats)
		if err != nil {
			fmt.Println("Scan failed,err:", err)
			return
		}
		fmt.Printf("screeningNum:%v cinemaNum:%v theaterNum:%v showTime:%v remainSeats:%v \n", s.screeningNum, s.movieNum, s.theaterNum, s.showTime, s.remainSeats)
	}
}

//统计所有用户总评论数，分组查询
func querySumEvaluationByUid() {
	sqlStr := `SELECT uid,name,COUNT(*)
			   FROM user NATURAL JOIN evaluation
			   GROUP BY uid`
	query, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println("Prepare failed,err:", err)
		return
	}
	var uid string
	var name string
	var count int
	for query.Next() {
		err := query.Scan(&uid, &name, &count)
		if err != nil {
			fmt.Println("Scan failed,err:", err)
			return
		}
		fmt.Printf("uid:%v name:%v count:%v\n", &uid, &name, &count)
	}
}

//查询高于指定分数的电影的场次信息，嵌套查询
func queryScreeningsByScore(score float64) {
	sqlStr := `SELECT *
				FROM screenings
				WHERE movieNum IN(
				    SELECT movieNum
				    FROM movie
				    WHERE aveFilmScore >= ?
				)`
	//进行预处理，先将sql发送给mysql服务端
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("Prepare failed,err:", err)
		return
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			return
		}
	}(stmt)

	//发送请求参数
	query, err := stmt.Query(score)
	if err != nil {
		fmt.Println("Query failed,err:", err)
		return
	}
	var s Screenings
	for query.Next() {
		err := query.Scan(&s.screeningNum, &s.movieNum, &s.theaterNum, &s.showTime, &s.remainSeats)
		if err != nil {
			fmt.Println("Scan failed,err:", err)
			return
		}
		fmt.Printf("screeningNum:%v cinemaNum:%v theaterNum:%v showTime:%v remainSeats:%v \n", s.screeningNum, s.movieNum, s.theaterNum, s.showTime, s.remainSeats)
	}
}
