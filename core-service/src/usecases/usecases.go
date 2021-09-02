package usecases

import (
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/repositories/mysql"
)

// Usecases ...
type Usecases struct {
	CategoryUcase    domain.CategoryUsecase
	NewsUcase        domain.NewsUsecase
	InformationUcase domain.InformationUsecase
	UnitUcase        domain.UnitUsecase
	EventUcase       domain.EventUsecase
}

// NewUcase will create an object that represent all usecases interface
func NewUcase(r *mysql.Repositories, timeoutContext time.Duration) *Usecases {
	return &Usecases{
		NewsUcase:        NewNewsUsecase(r.NewsRepo, r.CategoryRepo, timeoutContext),
		InformationUcase: NewInformationUsecase(r.InformationRepo, r.CategoryRepo, timeoutContext),
		UnitUcase:        NewUnitUsecase(r.UnitRepo, timeoutContext),
		EventUcase:       NewEventUsecase(r.EventRepo, r.CategoryRepo, timeoutContext),
	}
}
