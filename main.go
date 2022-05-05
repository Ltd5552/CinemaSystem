package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := initDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("Init db failed,err:%v\n", err)
		return
	} else {
		fmt.Println("Init successfully!")
	}
UA:
	fmt.Println("-----请选择用户/管理员-----")
	fmt.Println("        1-用户")
	fmt.Println("        2-管理员")
	fmt.Println("        0-退出")
	var op string
	_, err = fmt.Scan(&op)
	if err != nil {
		fmt.Println("输入失败请重新尝试...")
		goto UA
	}
	switch op {
	case "0":
		break
	case "1":
		var u User
		for true {
		LR:
			fmt.Println("-----欢迎用户!-----")
			fmt.Println("     1-登录")
			fmt.Println("     2-注册")
			fmt.Println("     0-退出")
			_, err = fmt.Scan(&op)
			if err != nil {
				fmt.Println("输入失败请重新尝试...")
				continue
			}
			switch op {
			case "1":
				if login(&u) {
					fmt.Println("登录成功!")
					for true {
						fmt.Println("-----用户界面-----")
						fmt.Println("     1-查询个人信息")
						fmt.Println("     2-更新个人信息")
						fmt.Println("     3-买票")
						fmt.Println("     4-查看电影信息")
						fmt.Println("     5-查看放映厅信息")
						fmt.Println("     6-查看电影院信息")
						fmt.Println("     7-通过电影名查询场次")
						fmt.Println("     8-通过电影分数查询场次")
						fmt.Println("     9-查询所有场次")
						fmt.Println("     10-发布评论")
						fmt.Println("     0-退出")
						_, err = fmt.Scan(&op)
						if err != nil {
							return
						}
						switch op {
						case "0":
							goto LR
						case "1":
							u.Query()
						case "2":
							u.Update()
						case "3":
							fmt.Println("请输入需要购买的场次信息")
							var num string
							_, err := fmt.Scan(&num)
							if err != nil {
								fmt.Println("输入失败请重新尝试...")
								break
							}
							if !u.Buy(num) {
								fmt.Println("购买失败...")
							}
						case "4":
							fmt.Println("请输入需要查询的电影名称")
							var name string
							_, err := fmt.Scan(&name)
							if err != nil {
								fmt.Println("输入失败请重新尝试...")
								break
							}
							u.QueryMovie(name)
						case "5":
							fmt.Println("请输入需要查询的放映厅号")
							var num string
							_, err := fmt.Scan(&num)
							if err != nil {
								fmt.Println("输入失败请重新尝试...")
								break
							}
							u.QueryTheater(num)
						case "6":
							fmt.Println("请输入需要查询的电影院名")
							var name string
							_, err := fmt.Scan(&name)
							if err != nil {
								fmt.Println("输入失败请重新尝试...")
								break
							}
							u.QueryCinema(name)
						case "7":
							fmt.Println("请输入你想观看的电影名")
							var name string
							_, err := fmt.Scan(&name)
							if err != nil {
								fmt.Println("输入失败请重新尝试...")
								break
							}
							u.QueryScreeningByMovie(name)
						case "8":
							fmt.Println("请输入分数")
							var score float64
							_, err := fmt.Scan(&score)
							if err != nil {
								fmt.Println("输入失败请重新尝试...")
								break
							}
							u.QueryScreeningByScore(score)
						case "9":
							u.QueryScreeningAll()
						case "10":
							fmt.Println("请输入评价的电影编号,电影院编号以及对应分数")
							var film, cinema string
							var fs, cs int64
							_, err := fmt.Scan(&film, &fs, &cinema, &cs)
							if err != nil {
								fmt.Println("输入失败请重新尝试...")
								break
							}
							u.Release(cinema, cs, film, fs)
						}
					}
				} else {
					fmt.Println("登录失败,请重新尝试账户或密码")
					break
				}
			case "2":
				if register(&u) {
					fmt.Println("注册成功!")
				} else {
					fmt.Println("注册失败,请重新尝试...")
				}
				break
			case "0":
				goto UA
			default:
				fmt.Println("指令有误，请重新尝试输入")
			}
		}

	case "2":
		var a Administrator
		for true {
			fmt.Println("-----欢迎管理员!-----")
			fmt.Println("     1-查看电影信息")
			fmt.Println("     2-查看场次信息")
			fmt.Println("     3-查看取票机信息")
			fmt.Println("     4-查看用户评论信息")
			fmt.Println("     5-查看用户总评论数")
			fmt.Println("     6-增加电影信息")
			fmt.Println("     7-增加场次信息")
			fmt.Println("     8-删除电影信息")
			fmt.Println("     9-删除场次信息")
			fmt.Println("     0-退出")
			_, err = fmt.Scan(&op)
			if err != nil {
				fmt.Println("输入失败请重新尝试...")
				continue
			}
			switch op {
			case "0":
				goto UA
			case "1":
				fmt.Println("请输入需要查询的电影名称")
				var name string
				_, err := fmt.Scan(&name)
				if err != nil {
					fmt.Println("输入失败请重新尝试...")
					break
				}
				a.QueryMovie(name)
			case "2":
				a.QueryScreening()
			case "3":
				a.QueryLeftTicket()
			case "4":
				a.QueryAllEvaluation()
			case "5":
				a.QuerySumEvaluation()
			case "6":
				fmt.Println("请输入新增的电影编号、名称、上映日期、时长")
				var m Movie
				_, err := fmt.Scan(&m.movieNum, &m.movieTitle, &m.releaseDate, &m.duration)
				if err != nil {
					fmt.Println("输入失败请重新尝试...")
					break
				}
				a.InsertMovie(m.movieNum, m.movieTitle, m.releaseDate, m.duration)
			case "7":
				fmt.Println("请输入新增的场次编号、电影编号、放映厅号、放映时间、剩余座位")
				var s Screenings
				_, err := fmt.Scan(&s.screeningNum, &s.movieNum, &s.theaterNum, &s.showTime, &s.remainSeats)
				if err != nil {
					fmt.Println("输入失败请重新尝试...")
					break
				}
				a.InsertScreening(s.screeningNum, s.movieNum, s.theaterNum, s.showTime, s.remainSeats)
			case "8":
				fmt.Println("请输入需要删除的电影名")
				var name string
				_, err := fmt.Scan(&name)
				if err != nil {
					fmt.Println("输入失败请重新尝试...")
					break
				}
				a.DeleteMovie(name)
			case "9":
				fmt.Println("请输入需要删除的场次编号")
				var num string
				_, err := fmt.Scan(&num)
				if err != nil {
					fmt.Println("输入失败请重新尝试...")
					break
				}
				a.DeleteScreening(num)
			default:
				fmt.Println("指令有误，请重新尝试输入")
			}
		}

	default:
		fmt.Println("指令有误，请重新尝试输入")
	}
	//关闭数据库连接
	err = db.Close()
	if err != nil {
		fmt.Println("Close err:", err)
		return
	} else {
		fmt.Println("Close successfully!")
	}

}

func login(u *User) bool {
	fmt.Println("请输入uid和password")
	var uid, password string
	_, err := fmt.Scan(&uid, &password)
	if err != nil {
		return false
	}
	u.uid = uid
	u.password = password
	return u.Login()
}

func register(u *User) bool {
	fmt.Println("请输入注册的用户名和密码,或输入quit退出")
	var name, password string

	_, err := fmt.Scan(&name)
	if err != nil {
		return false
	}

	if name == "quit" {
		return false
	}

	_, err = fmt.Scan(&password)
	if err != nil {
		return false
	}
	u.name = name
	u.password = password
	return u.Register()
}
