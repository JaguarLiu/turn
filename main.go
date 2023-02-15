package main

import (
	"github.com/JaguarLiu/turn/app/api"
	"github.com/JaguarLiu/turn/app/web"
	"github.com/JaguarLiu/turn/service"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.Static("/assets", "./app/template/assets")
	fileSrv := service.NewFileSrv()
	api.NewFileRouter(route, fileSrv)
	web.NewWebRouter(route)
	route.Run()
}
