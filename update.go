package main

import (
	"fmt"
)

// 更新座位数
func updateSeats(num string) {
	sqlStr := "update screenings set remainSeats = ? where screeningNum = ?"
	_, err := db.Exec(sqlStr, haveSeat(num)-1, num)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	} else {
		fmt.Println("Seats update success")
	}
}

//更新个人信息
func updatePersonal(uid string) {
	sqlStr := "update user set name=?,sex=?,birthday=?,location=?,phone=? where uid = ?"
	var u User
	u.uid = uid
	fmt.Println("请依次输入昵称 性别 生日 地址 电话")
	_, err := fmt.Scan(&u.name, &u.sex, &u.birthday, &u.location, &u.phone)
	if err != nil {
		return
	}

	_, err = db.Exec(sqlStr, u.name, u.sex, u.birthday, u.location, u.phone, u.uid)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	} else {
		fmt.Println("update success")
	}
}

// updateFilmScore 更新电影分数
func updateFilmScore() {
	sqlStr := `SELECT movieNum,AVG(fileScore)
				FROM evaluation
				GROUP BY movieNum`
	query, err := db.Query(sqlStr)
	if err != nil {
		fmt.Printf("Get ave failed, err:%v\n", err)
		return
	}
	var aveFilmScore float64
	var num string
	for query.Next() {
		err = query.Scan(&num, &aveFilmScore)
		if err != nil {
			return
		}
		FilmScore(num, aveFilmScore)
	}

}

func FilmScore(num string, score float64) {
	sqlStr := "update movie set aveFilmScore = ? where movieNum = ?"
	_, err := db.Exec(sqlStr, score, num)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	} else {
		fmt.Println("filmScore update success")
	}
}

// updateCinemaScore 更新电影院分数
func updateCinemaScore() {
	sqlStr := `SELECT cinemaNum,AVG(cinemaScore)
				FROM evaluation
				GROUP BY cinemaNum`
	query, err := db.Query(sqlStr)
	if err != nil {
		fmt.Printf("Get ave failed, err:%v\n", err)
		return
	}
	var aveCinemaScore float64
	var num string
	for query.Next() {
		err = query.Scan(&num, &aveCinemaScore)
		if err != nil {
			return
		}
		CinemaScore(num, aveCinemaScore)
	}

}

func CinemaScore(num string, score float64) {
	sqlStr := "update cinema set aveCinemaScore = ? where cinemaNum = ?"
	_, err := db.Exec(sqlStr, score, num)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	} else {
		fmt.Println("cinemaScore update success")
	}
}
