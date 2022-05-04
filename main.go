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

	//关闭数据库连接
	err = db.Close()
	if err != nil {
		fmt.Println("Close err:", err)
		return
	} else {
		fmt.Println("Close successfully!")
	}

}
