package main

import (
	_ "gin_admin_real/docs"
	"gin_admin_real/initRouter"
)

// @title Gin swagger
// @version 1.0
// @description Gin swagger 示例项目
// @contact.name realpeanut
// @contact.url www.kmf.com
// @contact.email 1743902627@qq.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:80
func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()
}
