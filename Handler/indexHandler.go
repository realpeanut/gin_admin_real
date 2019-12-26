package Handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Index(context *gin.Context)  {
	context.HTML(http.StatusOK,"tes",gin.H{
		"title": "hello gin " + strings.ToLower(context.Request.Method) + " method",
	})
}
