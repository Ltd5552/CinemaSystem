package main

import (
	"fmt"
	"time"
)

// Register 注册函数
func (u *User) Register() bool {
	u.uid = RandStr()
	insertPersonal(u.uid, u.name, u.password)
	fmt.Println("注册成功,uid为", u.uid)
	return true
}

// Login 登录函数
func (u *User) Login() bool {
	if !haveUser(u.uid) {
		fmt.Println("User is not exist,please try again")
		return false
	}

	sqlStr := "SELECT password FROM account WHERE uid = ?"

	query, err := db.Query(sqlStr, u.uid)
	if err != nil {
		fmt.Println("Query failed,err:", err)
		return false
	}
	var m User
	for query.Next() {
		err := query.Scan(&m.password)
		if err != nil {
			fmt.Println("Scan failed,err:", err)
			return false
		}
		if u.password == m.password {
			fmt.Println("Login successfully!")
			return true
		}
	}
	return false
}

// Query 查询个人信息函数
func (u *User) Query() {
	queryPersonal(u.uid)
}

// Update 更新个人信息函数
func (u *User) Update() {
	updatePersonal(u.uid)
}

// Buy 用户购买函数，传入购买的场次编号
func (u *User) Buy(num string) bool {
	begin, err := db.Begin()
	if err != nil {
		return false
	}
	if haveSeat(num) == 0 {
		fmt.Println("Failed,no seat left!")
		return false
	}

	if !haveScreenings(num) {
		fmt.Println("Failed,no such screening!")
		return false
	}
	sqlStr := "INSERT INTO ticket(ticketNum, price, screeningNum) VALUES (?,?,?)"
	var t Ticket
	t.ticketNum = RandStr()
	t.price = 49.9
	t.screeningNum = num
	_, err = db.Exec(sqlStr, t.ticketNum, t.price, t.screeningNum)
	if err != nil {
		fmt.Println("Insert failed.", err)
		err := begin.Rollback()
		if err != nil {
			return false
		}
		return false
	}
	buyTime := time.Now().Format("2006-01-02 15:04:05") //当前时间的字符串，2006-01-02 15:04:05是固定写法

	sqlStr = "INSERT INTO buy(uid, ticketNum, buyTime) VALUES (?,?,?)"
	_, err = db.Exec(sqlStr, u.uid, t.ticketNum, buyTime)
	if err != nil {
		fmt.Println("Insert failed.", err)
		err := begin.Rollback()
		if err != nil {
			return false
		}
		return false
	}
	fmt.Println("购买成功")
	if !updateSeats(num) {
		err := begin.Rollback()
		if err != nil {
			return false
		}
	}
	err = begin.Commit()
	if err != nil {
		return false
	}
	return true
}

// QueryMovie 查询电影信息函数，传入电影名
func (u *User) QueryMovie(name string) {
	if !queryMovies(name) {
		fmt.Println("查询失败,没有该电影!")
	}
}

// QueryTheater 查询放映厅信息,传入放映厅编号
func (u *User) QueryTheater(num string) {
	queryTheater(num)
}

// QueryCinema 查询电影院信息,传入电影院名
func (u *User) QueryCinema(name string) {
	queryCinema(name)
}

// QueryScreeningByMovie 查询场次信息,参数为电影名
func (u *User) QueryScreeningByMovie(name string) {
	if !haveMovies(name) {
		fmt.Println("查询失败，没有该电影...")
		return
	}
	queryScreeningsByMovie(name)
}

// QueryScreeningByScore 查询场次信息,参数为电影分数
func (u *User) QueryScreeningByScore(score float64) {
	if score > 10 || score < 0 {
		fmt.Println("请输入0-10以内的小数或整数...")
		return
	}
	queryScreeningsByScore(score)
}

// QueryScreeningAll 查询所有场次信息
func (u *User) QueryScreeningAll() {
	queryScreenings()
}

// Release 发布评论的函数,参数为电影院和电影评分,均为十以内的整数,以及电影编号和电影院编号
func (u *User) Release(cinemaNum string, cinema int64, movieNum string, film int64) {
	Eid := RandStr()
	begin, err := db.Begin()
	if err != nil {
		return
	}
	sqlStr := "INSERT INTO evaluation(evaluationId, cinemaScore, fileScore,movieNum,cinemaNum) VALUES (?,?,?,?,?)"
	_, err = db.Exec(sqlStr, Eid, cinema, film, movieNum, cinemaNum)
	if err != nil {
		fmt.Println("Insert failed.", err)
		err := begin.Rollback()
		if err != nil {
			return
		}
		return
	}

	releaseTime := time.Now().Format("2006-01-02 15:04:05") //当前时间的字符串，2006-01-02 15:04:05是固定写法

	sqlStr = "INSERT INTO releases(uid, evaluationId, releaseTime) VALUES (?,?,?)"
	_, err = db.Exec(sqlStr, u.uid, Eid, releaseTime)
	if err != nil {
		fmt.Println("Insert failed.", err)
		err := begin.Rollback()
		if err != nil {
			return
		}
		return
	}
	//更新电影分数和电影院分数
	updateFilmScore()
	updateCinemaScore()
	err = begin.Commit()
	if err != nil {
		return
	}
}
