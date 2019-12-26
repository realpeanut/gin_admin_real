package main

import (
	"gin_do/initRouter"
)

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()
}
