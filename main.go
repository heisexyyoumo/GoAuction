package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	_ "GoAuction/routers"
	"GoAuction/models"
)

hello world
func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)

	orm.RegisterDataBase("default", "postgres", "user=alan password=helloalan host=localhost dbname=Alan sslmode=disable")

	orm.RunSyncdb("default", false, false)
}

func main() {
	o := orm.NewOrm()
	o.Using("default")
	user := new(models.User)
	user.ID = 0
	user.Name = "alice"
	user.Password = "helloAlice"

	fmt.Println(o.Insert(user))
}

