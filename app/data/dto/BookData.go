package dto

import "github.com/Geektree0101/bookstore/app/domain/entity"

type BookData struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Isbn13   string `json:"isbn13"`
	Price    string `json:"price"`
	Image    string `json:"image"`
	Url      string `json:"url"`
}

func (data BookData) ToDomain() entity.Book {
	return entity.Book{
		Title:    data.Title,
		Subtitle: data.Subtitle,
		Isbn13:   data.Isbn13,
		Price:    data.Price,
		Image:    data.Image,
		Url:      data.Url,
	}
}
