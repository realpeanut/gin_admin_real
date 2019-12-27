package model

import (
	"gin_admin_real/initDb"
	"log"
)

type UserModel struct {
	Email         string `form:"email" binding:"email"`
	Password      string `form:"password"`
}

func (user *UserModel) Save() int64 {
	result, e := initDb.Db.Exec("insert into gin_admin_real.user (email, password) values (?,?);", user.Email, user.Password)
	if e != nil {
		log.Panicln("user insert error", e.Error())
	}
	id, err := result.LastInsertId();
	if err != nil {
		log.Panicln("user insert id error", err.Error())
	}
	return id
}