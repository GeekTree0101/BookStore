package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Geektree0101/bookstore/app/domain/usecase"
	"github.com/Geektree0101/bookstore/app/presentation/model"
	"github.com/gin-gonic/gin"
)

type HomeController struct {
	SearchUseCase usecase.BookSearchUseCase
}

func (controller *HomeController) Index(c *gin.Context) {
	var query string
	var page int
	var err error

	query = c.Query("query")

	if len(query) == 0 {
		c.HTML(200, "home.html", model.HomeModel{IsEmpty: true})
		return
	}

	page, err = strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	result, err := controller.SearchUseCase.Execute(query, int64(page))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "home.html", model.RenderHomeModel(result))
}

func (controller *HomeController) Search(c *gin.Context) {
	query := c.Query("query")
	c.Redirect(http.StatusPermanentRedirect, fmt.Sprintf("index.html?query=%s&page=1", query))
}
