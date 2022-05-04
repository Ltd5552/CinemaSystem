package main

import "fmt"

// Register 注册函数
func (u *User) Register() {
	u.uid = RandStr()
	insertPersonal(u.uid, u.name, u.password)
}

// Login 登录函数
func (u *User) Login() bool {
	if !haveUser(u.uid) {
		fmt.Println("User is not exist,please try again")
		return false
	}

	sqlStr := "SELECT password FROM user WHERE uid = ?"

	query, err := db.Query(sqlStr, u.uid)
	if err != nil {
		fmt.Println("Query failed,err:", err)
		return false
	}
	var m User
	for query.Next() {
		err := query.Scan(&m.uid, &m.name, &m.sex, &m.birthday, &m.location, &m.phone, &m.password)
		if err != nil {
			fmt.Println("Scan failed,err:", err)
			return false
		}
	}
	if u.password == m.password {
		fmt.Println("Login successfully!")
		return true
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
func (u *User) Buy(num string) {
	if haveSeat(num) == 0 {
		fmt.Println("Failed,no seat left!")
		return
	}
	sqlStr := "INSERT INTO ticket(ticketNum, price, screeningNum) VALUES (?,?,?)"
	var t Ticket
	t.ticketNum = RandStr()
	t.price = 49.9
	t.screeningNum = num
	_, err := db.Exec(sqlStr, t.ticketNum, t.price, t.screeningNum)
	if err != nil {
		fmt.Println("Insert failed.", err)
		return
	}
	updateSeats(num)
}

// QueryMovie 查询电影信息函数，传入电影名
func (u *User) QueryMovie(name string) {
	queryMovies(name)
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
	queryScreeningsByMovie(name)
}

// QueryScreeningByScore 查询场次信息,参数为电影分数
func (u *User) QueryScreeningByScore(score float64) {
	queryScreeningsByScore(score)
}

// QueryScreeningAll 查询所有场次信息
func (u *User) QueryScreeningAll() {
	queryScreenings()
}

// Release 发布评论的函数,参数为电影院和电影评分,均为十以内的整数,以及电影编号和电影院编号
func (u *User) Release(cinemaNum string, cinema int64, movieNum string, film int64) {
	Eid := RandStr()
	sqlStr := "INSERT INTO evaluation(evaluationId, cinemaScore, fileScore,movieNum,cinemaNum) VALUES (?,?,?,?,?)"
	_, err := db.Exec(sqlStr, Eid, cinema, film, movieNum, cinemaNum)
	if err != nil {
		fmt.Println("Insert failed.", err)
		return
	}
	//更新电影分数和电影院分数
	updateFilmScore()
	updateCinemaScore()
}
