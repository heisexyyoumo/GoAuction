package models

import (
	"github.com/astaxie/beego/orm"
)

// User 用户数据模型
type User struct{
	ID int `pk: "auto"`
	Name string `orm: "size(20)"`
	Password string `orm: "size(20)`
}

func init() {
	orm.RegisterModelWithPrefix("g_a_", new(User))
}