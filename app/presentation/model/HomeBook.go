package model

import "github.com/Geektree0101/bookstore/app/domain/entity"

type HomeBookModel struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Isbn13   string `json:"isbn13"`
	Price    string `json:"price"`
	Image    string `json:"image"`
	Url      string `json:"url"`
}

func RenderHomeBookModel(book entity.Book) HomeBookModel {
	return HomeBookModel{
		Title:    book.Title,
		Subtitle: book.Subtitle,
		Isbn13:   book.Isbn13,
		Price:    book.Price,
		Image:    book.Image,
		Url:      book.Url,
	}
}

func RenderHomeBookModels(books []entity.Book) []HomeBookModel {
	models := []HomeBookModel{}
	for _, b := range books {
		models = append(models, RenderHomeBookModel(b))
	}
	return models
}
