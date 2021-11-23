package server

import (
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	_eventUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/event/usecase"
	_featuredProgramUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/featured-program/usecase"
	_feedbackUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/feedback/usecase"
	_informationUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/information/usecase"
	_newsUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/news/usecase"
	_searchUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/search/usecase"
	_unitUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/unit/usecase"
)

// Usecases ...
type Usecases struct {
	CategoryUcase        domain.CategoryUsecase
	NewsUcase            domain.NewsUsecase
	InformationUcase     domain.InformationUsecase
	UnitUcase            domain.UnitUsecase
	EventUcase           domain.EventUsecase
	FeedbackUcase        domain.FeedbackUsecase
	FeaturedProgramUcase domain.FeaturedProgramUsecase
	SearchUcase          domain.SearchUsecase
}

// NewUcase will create an object that represent all usecases interface
func NewUcase(r *Repository, timeoutContext time.Duration) *Usecases {
	return &Usecases{
		NewsUcase:            _newsUcase.NewNewsUsecase(r.NewsRepo, r.CategoryRepo, r.UserRepo, timeoutContext),
		InformationUcase:     _informationUcase.NewInformationUsecase(r.InformationRepo, r.CategoryRepo, timeoutContext),
		UnitUcase:            _unitUcase.NewUnitUsecase(r.UnitRepo, timeoutContext),
		EventUcase:           _eventUcase.NewEventUsecase(r.EventRepo, r.CategoryRepo, timeoutContext),
		FeedbackUcase:        _feedbackUcase.NewFeedbackUsecase(r.FeedbackRepo, timeoutContext),
		FeaturedProgramUcase: _featuredProgramUcase.NewFeaturedProgramUsecase(r.FeaturedProgramRepo, timeoutContext),
		SearchUcase:          _searchUcase.NewSearchUsecase(r.SearchRepo, timeoutContext),
	}
}
