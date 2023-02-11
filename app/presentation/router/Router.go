package router

import (
	"github.com/Geektree0101/bookstore/app/dependency"
	"github.com/gin-gonic/gin"
)

func Setup(serv *gin.Engine, container *dependency.DependencyContainer) {
	homeController := container.ResolveHomeController()
	serv.GET("index.html", homeController.Index)
	serv.GET("search", homeController.Search)
}
