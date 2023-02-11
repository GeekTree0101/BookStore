package repository

import (
	"fmt"
	"net/http"

	"github.com/Geektree0101/bookstore/app/data/dto"
)

type SearchRepository struct {
	Networking
}

func DefaultSearchRepository(networking Networking) *SearchRepository {
	return &SearchRepository{
		Networking: networking,
	}
}

func (repo *SearchRepository) Search(query string, page int64) (*dto.SearchResponse, error) {
	url := fmt.Sprintf("https://api.itbook.store/1.0/search/%s/%d", query, page)
	req, err := http.NewRequest("GET", url, nil)
	searchRes := &dto.SearchResponse{}
	err = repo.fetch(req, searchRes)
	return searchRes, err
}
