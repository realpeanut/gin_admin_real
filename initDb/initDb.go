package initDb

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *gorm.DB

func init()  {
	var err error
	Db, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/gin_admin_real")
	if err != nil {
		log.Panicln("err:", err.Error())
	}
}
