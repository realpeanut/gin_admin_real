package Handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UseSave(context *gin.Context)  {
	name := context.Param("name")
	context.String(http.StatusOK, name + "保存成功")
}

func UserQuerySave(context * gin.Context)  {
	name := context.DefaultQuery("name","default_value")
	context.String(http.StatusOK,name + "保存成功")
}