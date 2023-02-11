package usecase

import (
	"github.com/Geektree0101/bookstore/app/data/repository"
	"github.com/Geektree0101/bookstore/app/domain/entity"
)

type BookSearchUseCase struct {
	searchRepo *repository.SearchRepository
}

func DefaultBookSearchUseCase(searchRepo *repository.SearchRepository) BookSearchUseCase {
	return BookSearchUseCase{
		searchRepo: searchRepo,
	}
}

func (useCase *BookSearchUseCase) Execute(query string, page int64) (entity.SearchResult, error) {
	searchRes, err := useCase.searchRepo.Search(query, page)

	if err != nil {
		return entity.SearchResult{}, nil
	}

	return searchRes.ToDomain(), nil
}
