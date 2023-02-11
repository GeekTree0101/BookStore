package model

import (
	"github.com/Geektree0101/bookstore/app/domain/entity"
)

type HomeModel struct {
	Title   string          `json:"title"`
	Books   []HomeBookModel `json:"books"`
	IsEmpty bool            `json:"is_empty"`
}

func RenderHomeModel(searchResult entity.SearchResult) HomeModel {
	return HomeModel{
		Books:   RenderHomeBookModels(searchResult.Books),
		IsEmpty: false,
	}
}
