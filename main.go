package main

import (
	"fmt"
	"go-web1/config"
	"go-web1/router"
)

func main() {
	config.InitConfig()
	r := router.SetupRouter()
	port := config.Appconfig.App.Port
	if port == "" {
		port = ":8080"
	}
	fmt.Println(port)
	r.Run(port)
	//fmt.Println(config.Appconfig.App.Port)
	//type Info struct {
	//	Message string
	//}
	//InfoTest := Info{
	//	Message: "123123",
	//}
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, InfoTest)
	//})
}
