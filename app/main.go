package app

import (
	"github.com/Geektree0101/bookstore/app/dependency"
	"github.com/Geektree0101/bookstore/app/presentation/router"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	serv := gin.Default()
	container := dependency.DefaultDependencyContainer()
	router.Setup(serv, container)
	serv.LoadHTMLGlob("public/html/*.html")
	serv.Run(":8080")
}
