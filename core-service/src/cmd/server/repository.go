package server

import (
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/database"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	_areaRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/area/repository/mysql"
	_categoryRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/category/repository/mysql"
	_dataTagsRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/data-tags/repository/mysql"
	_eventRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/event/repository/mysql"
	_featuredProgramRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/featured-program/repository/mysql"
	_feedbackRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/feedback/repository/mysql"
	_informationRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/information/repository/mysql"
	_newsRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/news/repository/mysql"
	_searchRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/search/repository/elastic"
	_unitRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/unit/repository/mysql"
	_userRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/user/repository/mysql"
)

// Repository ...
type Repository struct {
	CategoryRepo        domain.CategoryRepository
	NewsRepo            domain.NewsRepository
	InformationRepo     domain.InformationRepository
	UnitRepo            domain.UnitRepository
	AreaRepo            domain.AreaRepository
	EventRepo           domain.EventRepository
	FeedbackRepo        domain.FeedbackRepository
	FeaturedProgramRepo domain.FeaturedProgramRepository
	UserRepo            domain.UserRepository
	DataTagsRepo        domain.DataTagsRepository
	SearchRepo          domain.SearchRepository
}

// NewRepository will create an object that represent all repos interface
func NewRepository(conn *database.DBConn) *Repository {
	return &Repository{
		CategoryRepo:        _categoryRepo.NewMysqlCategoryRepository(conn.Mysql),
		NewsRepo:            _newsRepo.NewMysqlNewsRepository(conn.Mysql),
		InformationRepo:     _informationRepo.NewMysqlInformationRepository(conn.Mysql),
		UnitRepo:            _unitRepo.NewMysqlUnitRepository(conn.Mysql),
		AreaRepo:            _areaRepo.NewMysqlAreaRepository(conn.Mysql),
		EventRepo:           _eventRepo.NewMysqlEventRepository(conn.Mysql),
		FeedbackRepo:        _feedbackRepo.NewMysqlFeedbackRepository(conn.Mysql),
		FeaturedProgramRepo: _featuredProgramRepo.NewMysqlFeaturedProgramRepository(conn.Mysql),
		UserRepo:            _userRepo.NewMysqlUserRepository(conn.Mysql),
		DataTagsRepo:        _dataTagsRepo.NewMysqlDataTagsRepository(conn.Mysql),
		SearchRepo:          _searchRepo.NewElasticSearchRepository(conn.Elastic),
	}
}
