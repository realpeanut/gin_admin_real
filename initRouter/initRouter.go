package initRouter

import (
	"gin_admin_real/Handler"
	"gin_admin_real/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func SetupRouter() *gin.Engine {

	router := gin.New()
	router.Use(middleware.Logger(),gin.Recovery())
	router.LoadHTMLGlob("templates/**/*")
	//router.LoadHTMLFiles("templates/test.tmpl","templates/ab.tmpl")

	router.Static("/statics","./statics")


	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// 添加 Get 请求路由
	router.GET("/", middleware.SetCookie(), middleware.Auth(), Handler.Index)

	// 添加 Post 请求路由
	router.POST("/post", fmtResponse)

	// 添加 Put 请求路由
	router.PUT("/", fmtResponse)

	// 添加 Delete 请求路由
	router.DELETE("/", fmtResponse)

	// 添加 Patch 请求路由
	router.PATCH("/", fmtResponse)

	// 添加 Head 请求路由
	router.HEAD("/", fmtResponse)

	// 添加 Options 请求路由
	router.OPTIONS("/", fmtResponse)



	//user路由组
	user := router.Group("/user")
	{
		user.GET("/save/:name",Handler.UseSave)
		user.GET("/saveQuery",Handler.UserQuerySave)
		user.POST("/register",Handler.UserRegister)
	}

	return router
}

func fmtResponse(context *gin.Context){
	context.String(http.StatusOK, context.Request.Method)
}
