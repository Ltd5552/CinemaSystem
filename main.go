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
	fmt.Println("请选择用户/管理员")
	fmt.Println("1-用户")
	fmt.Println("2-管理员")
	fmt.Println("0-退出")
	var op int64
	_, err = fmt.Scan(&op)
	if err != nil {
		return
	}
	switch op {
	case 0:
		break
	case 1:
		var u User
		for true {
		LR:
			fmt.Println("-----欢迎用户!-----")
			fmt.Println("     1-登录")
			fmt.Println("     2-注册")
			fmt.Println("     0-退出")
			_, err = fmt.Scan(&op)
			if err != nil {
				return
			}
			switch op {
			case 1:
				if login(&u) {
					fmt.Println("登录成功!")
					for true {
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
						case 0:
							goto LR
						case 1:
							u.Query()
						case 2:
							u.Update()
						case 3:
							fmt.Println("请输入需要购买的场次信息")
							var num string
							fmt.Scan(&num)
							u.Buy(num)
						case 4:
							fmt.Println("请输入需要查询的电影名称")
							var name string
							fmt.Scan(&name)
							u.QueryMovie(name)
						case 5:
							fmt.Println("请输入需要查询的放映厅号")
							var num string
							fmt.Scan(&num)
							u.QueryTheater(num)
						case 6:
							fmt.Println("请输入需要查询的电影院名")
							var name string
							fmt.Scan(&name)
							u.QueryCinema(name)
						case 7:
							fmt.Println("请输入你想观看的电影名")
							var name string
							fmt.Scan(&name)
							u.QueryScreeningByMovie(name)
						case 8:
							fmt.Println("请输入分数")
							var score float64
							fmt.Scan(&score)
							u.QueryScreeningByScore(score)
						case 9:
							u.QueryScreeningAll()
						case 10:
							fmt.Println("请输入评价的电影编号,电影院编号以及对应分数")
							var film, cinema string
							var fs, cs int64
							fmt.Scan(&film, &fs, &cinema, &cs)
							u.Release(cinema, cs, film, fs)
						}
						break
					}
				} else {
					fmt.Println("登录失败,请重新尝试账户或密码")
					break
				}
			case 2:
				if register(&u) {
					fmt.Println("注册成功!")
				} else {
					fmt.Println("注册失败,请重新尝试...")
				}
				break
			case 0:
				goto UA
			}
		}

	case 2:
		//var a Administrator
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
	fmt.Println("请输入注册的用户名和密码")
	var name, password string
	_, err := fmt.Scan(&name, &password)
	if err != nil {
		return false
	}
	u.name = name
	u.password = password
	return u.Register()
}
