package server

import (
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/utils"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	_areaRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/area/repository/mysql"
	_awardRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/award/repository/mysql"
	_categoryRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/category/repository/mysql"
	_dataTagRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/data-tag/repository/mysql"
	_districtRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/district/repository/mysql"
	_documentArchiveRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/document-archive/repository/mysql"
	_eventRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/event/repository/mysql"
	_featuredProgramRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/featured-program/repository/mysql"
	_feedbackRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/feedback/repository/mysql"
	_informationRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/information/repository/mysql"
	_newsRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/news/repository/mysql"
	_regInvitationRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/registration-invitation/repository/mysql"
	_rolePermRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/role-permission/repository/mysql"
	_roleRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/role/repository/mysql"
	_searchRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/search/repository/elastic"
	_tagRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/tag/repository/mysql"
	_templateRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/template/repository/mysql"
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
	DataTagsRepo        domain.DataTagRepository
	TagRepo             domain.TagRepository
	SearchRepo          domain.SearchRepository
	RoleRepo            domain.RoleRepository
	RolePermRepo        domain.RolePermissionRepository
	TemplateRepo        domain.TemplateRepository
	RegInvitationRepo   domain.RegistrationInvitationRepository
	MailRepo            domain.MailRepository
	AwardRepo           domain.AwardRepository
	DistrictRepo        domain.DistrictRepository
	DocumentArchiveRepo domain.DocumentArchiveRepository
}

// NewRepository will create an object that represent all repos interface
func NewRepository(conn *utils.Conn) *Repository {
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
		DataTagsRepo:        _dataTagRepo.NewMysqlDataTagRepository(conn.Mysql),
		TagRepo:             _tagRepo.NewMysqlTagRepository(conn.Mysql),
		SearchRepo:          _searchRepo.NewElasticSearchRepository(conn.Elastic),
		RoleRepo:            _roleRepo.NewMysqlRoleRepository(conn.Mysql),
		RolePermRepo:        _rolePermRepo.NewMysqlRolePermissionRepository(conn.Mysql),
		TemplateRepo:        _templateRepo.NewMysqlMailTemplateRepository(conn.Mysql),
		RegInvitationRepo:   _regInvitationRepo.NewMysqlRegInvitationRepository(conn.Mysql),
		AwardRepo:           _awardRepo.NewMysqlAwardRepository(conn.Mysql),
		DistrictRepo:        _districtRepo.NewMysqlDistrictRepository(conn.Mysql),
		DocumentArchiveRepo: _documentArchiveRepo.NewMysqlDocumentArchiveRepository(conn.Mysql),
	}
}
