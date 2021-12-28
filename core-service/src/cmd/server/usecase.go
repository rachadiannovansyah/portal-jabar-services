package server

import (
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	_authUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/auth/usecase"
	_eventUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/event/usecase"
	_featuredProgramUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/featured-program/usecase"
	_feedbackUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/feedback/usecase"
	_informationUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/information/usecase"
	_newsUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/news/usecase"
	_searchUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/search/usecase"
	_unitUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/unit/usecase"
	_userUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/user/usecase"
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
	AuthUcase            domain.AuthUsecase
	SearchUcase          domain.SearchUsecase
	UserUsecase          domain.UserUsecase
}

// NewUcase will create an object that represent all usecases interface
func NewUcase(cfg *config.Config, r *Repository, timeoutContext time.Duration) *Usecases {
	return &Usecases{
		NewsUcase:            _newsUcase.NewNewsUsecase(r.NewsRepo, r.CategoryRepo, r.UserRepo, r.DataTagsRepo, timeoutContext),
		InformationUcase:     _informationUcase.NewInformationUsecase(r.InformationRepo, r.CategoryRepo, timeoutContext),
		UnitUcase:            _unitUcase.NewUnitUsecase(r.UnitRepo, timeoutContext),
		EventUcase:           _eventUcase.NewEventUsecase(r.EventRepo, r.CategoryRepo, r.DataTagsRepo, timeoutContext),
		FeedbackUcase:        _feedbackUcase.NewFeedbackUsecase(r.FeedbackRepo, timeoutContext),
		FeaturedProgramUcase: _featuredProgramUcase.NewFeaturedProgramUsecase(r.FeaturedProgramRepo, timeoutContext),
		AuthUcase:            _authUcase.NewAuthUsecase(cfg, r.UserRepo, timeoutContext),
		SearchUcase:          _searchUcase.NewSearchUsecase(r.SearchRepo, timeoutContext),
		UserUsecase:          _userUcase.NewUserkUsecase(r.UserRepo, timeoutContext),
	}
}
