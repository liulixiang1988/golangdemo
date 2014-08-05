package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/lunny/godbc"
)

type User struct {
	Id      int
	Name    string
	Profile *Profile `orm:"rel(one)"` // OneToOne relation
}

type Profile struct {
	Id   int
	Age  int16
	User *User `orm:"reverse(one)"` // 设置反向关系(可选)
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User), new(Profile))
	orm.RegisterDriver("mssql", 5)

	orm.RegisterDataBase("default", "mssql", "server=localhost;user id=sa;password=tlys.oaxmz.5860247;port=1433;database=test")

	// orm.RegisterDriver("odbc", 5)

	// orm.RegisterDataBase("default", "odbc", "driver={sql server};server=127.0.0.1;port=1433;uid=sa;pwd=tlys.oaxmz.5860247;database=test")
}

func main() {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	// 数据库别名
	name := "default"

	// drop table 后再建表
	force := false

	// 打印执行过程
	verbose := true

	// 遇到错误立即返回
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
	// profile := new(Profile)
	// profile.Age = 30

	// user := new(User)
	// user.Profile = profile
	// user.Name = "slene"

	// fmt.Println(o.Insert(profile))
	// fmt.Println(o.Insert(user))

}
