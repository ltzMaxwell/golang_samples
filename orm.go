package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type User struct {
	Username string
	Nickname string
}

func init() {
	// register model
	// orm.RegisterModel(new(Options))

	// set default database
	orm.RegisterDataBase("default", "mysql", "root:4281018@/zhenzu?charset=utf8", 30)
}

func main() {
	o := orm.NewOrm()
	//query row
	var user User
	err := o.Raw("SELECT username, nickname FROM users where id = ?", 103).QueryRow(&user)
	if err != nil {
		panic(err)
	}
	fmt.Println("username is :", user.Username)
	fmt.Println("nickname is :", user.Nickname)
	//query rows
	var users []User
	num, err := o.Raw("SELECT username ,nickname FROM users").QueryRows(&users)
	fmt.Println("num is :", num)
	fmt.Println("users length is :", len(users))
	for i := 0; i < len(users); i++ {
		fmt.Println("results are :", users[i].Username)
		fmt.Println("results are :", users[i].Nickname)
	}
}
