package dependency

import (
	"github.com/Geektree0101/bookstore/app/data/repository"
	"github.com/Geektree0101/bookstore/app/domain/usecase"
	"github.com/Geektree0101/bookstore/app/presentation/controller"
)

/// Toooo ugly..., i will use an IoC library :(
type DependencyContainer struct {
	searchRepo *repository.SearchRepository
}

func DefaultDependencyContainer() *DependencyContainer {
	networking := repository.DefaultNetworking()
	return &DependencyContainer{
		searchRepo: repository.DefaultSearchRepository(networking),
	}
}

func (container *DependencyContainer) ResolveHomeController() *controller.HomeController {
	return &controller.HomeController{
		SearchUseCase: usecase.DefaultBookSearchUseCase(
			container.searchRepo,
		),
	}
}
