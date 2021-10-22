package server

import (
	"database/sql"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	_areaRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/area/repository/mysql"
	_categoryRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/category/repository/mysql"
	_eventRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/event/repository/mysql"
	_featuredProgramRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/featured-program/repository/mysql"
	_feedbackRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/feedback/repository/mysql"
	_informationRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/information/repository/mysql"
	_newsRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/news/repository/mysql"
	_unitRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/unit/repository/mysql"
	_userRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/user/repository/mysql"
)

// Repositories ...
type Repositories struct {
	CategoryRepo        domain.CategoryRepository
	NewsRepo            domain.NewsRepository
	InformationRepo     domain.InformationRepository
	UnitRepo            domain.UnitRepository
	AreaRepo            domain.AreaRepository
	EventRepo           domain.EventRepository
	FeedbackRepo        domain.FeedbackRepository
	FeaturedProgramRepo domain.FeaturedProgramRepository
	UserRepo            domain.UserRepository
}

// NewMysqlRepositories will create an object that represent all repos interface
func NewMysqlRepositories(Conn *sql.DB) *Repositories {
	return &Repositories{
		CategoryRepo:        _categoryRepo.NewMysqlCategoryRepository(Conn),
		NewsRepo:            _newsRepo.NewMysqlNewsRepository(Conn),
		InformationRepo:     _informationRepo.NewMysqlInformationRepository(Conn),
		UnitRepo:            _unitRepo.NewMysqlUnitRepository(Conn),
		AreaRepo:            _areaRepo.NewMysqlAreaRepository(Conn),
		EventRepo:           _eventRepo.NewMysqlEventRepository(Conn),
		FeedbackRepo:        _feedbackRepo.NewMysqlFeedbackRepository(Conn),
		FeaturedProgramRepo: _featuredProgramRepo.NewMysqlFeaturedProgramRepository(Conn),
		UserRepo:            _userRepo.NewMysqlUserRepository(Conn),
	}
}
