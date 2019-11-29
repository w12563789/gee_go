package main

import (
	"geeoffice/route"
)

func main() {

	router := route.RegisterRoutes()
	// 绑定端口是8080
	router.Run(":8080")
}