package Handler

import (
	"gin_admin_real/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func UseSave(context *gin.Context) {
	name := context.Param("name")
	context.String(http.StatusOK, name+"保存成功")
}

func UserQuerySave(context *gin.Context) {
	name := context.DefaultQuery("name", "default_value")
	context.String(http.StatusOK, name+"保存成功")
}

func UserRegister(ctx *gin.Context) {

	var user model.UserModel
	if err := ctx.ShouldBind(&user);err != nil {
		log.Println("error->",err.Error())
		ctx.String(http.StatusBadRequest,"数据不合法")
	} else {
		id := user.Save();
		log.Printf("user_add",id)
	}



}
