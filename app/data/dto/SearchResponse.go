package dto

import "github.com/Geektree0101/bookstore/app/domain/entity"

type SearchResponse struct {
	Total int64      `json:"total,string"`
	Page  int64      `json:"page,string"`
	Books []BookData `json:"books"`
}

func (data SearchResponse) ToDomain() entity.SearchResult {
	return entity.SearchResult{
		Page:  data.Page,
		Books: mapping(data.Books),
	}
}

func mapping(data []BookData) []entity.Book {
	books := []entity.Book{}
	for _, d := range data {
		books = append(books, d.ToDomain())
	}
	return books
}
