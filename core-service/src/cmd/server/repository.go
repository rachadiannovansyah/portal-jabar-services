package server

import (
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
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
	_generalInformationRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/general-information/repository/mysql"
	_governmentAffairRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/government-affair/repository/mysql"
	_informationRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/information/repository/mysql"
	_masterDataServiceRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/master-data-service/repository/mysql"
	_additionalInformationRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/mds-additional-information/repository/mysql"
	_applicationRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/mds-application/repository/mysql"
	_mainServiceRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/mds-main-service/repository/mysql"
	_newsRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/news/repository/mysql"
	_popUpBannerRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/pop-up-banner/repository/mysql"
	_publicServiceRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/public-service/repository/mysql"
	_regInvitationRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/registration-invitation/repository/mysql"
	_rolePermRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/role-permission/repository/mysql"
	_roleRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/role/repository/mysql"
	_searchRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/search/repository/elastic"
	_servicePublicRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/service-public/repository/mysql"
	_spbeRalsRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/spbe-rals/repository/mysql"
	_tagRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/tag/repository/mysql"
	_templateRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/template/repository/mysql"
	_unitRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/unit/repository/mysql"
	_uptdCabdinRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/uptd-cabdin/repository/mysql"
	_userRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/user/repository/mysql"
	_externalVisitorRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/visitor/repository/external"
)

// Repository ...
type Repository struct {
	CategoryRepo           domain.CategoryRepository
	NewsRepo               domain.NewsRepository
	InformationRepo        domain.InformationRepository
	UnitRepo               domain.UnitRepository
	AreaRepo               domain.AreaRepository
	EventRepo              domain.EventRepository
	FeedbackRepo           domain.FeedbackRepository
	FeaturedProgramRepo    domain.FeaturedProgramRepository
	UserRepo               domain.UserRepository
	DataTagsRepo           domain.DataTagRepository
	TagRepo                domain.TagRepository
	SearchRepo             domain.SearchRepository
	ServicePublicRepo      domain.ServicePublicRepository
	GeneralInformationRepo domain.GeneralInformationRepository
	RoleRepo               domain.RoleRepository
	RolePermRepo           domain.RolePermissionRepository
	TemplateRepo           domain.TemplateRepository
	RegInvitationRepo      domain.RegistrationInvitationRepository
	MailRepo               domain.MailRepository
	AwardRepo              domain.AwardRepository
	DistrictRepo           domain.DistrictRepository
	DocumentArchiveRepo    domain.DocumentArchiveRepository
	PublicServiceRepo      domain.PublicServiceRepository
	ExternalVisitorRepo    domain.ExternalVisitorRepository
	PopUpBannerRepo        domain.PopUpBannerRepository
	GovernmentAffairRepo   domain.GovernmentAffairRepository
	SpbeRalsRepo           domain.SpbeRalsRepository
	UptdCabdinRepo         domain.UptdCabdinRepository
	MasterDataServiceRepo  domain.MasterDataServiceRepository
	MainServiceRepo        domain.MainServiceRepository
	ApplicationRepo        domain.ApplicationRepository
	AdditionalInfRepo      domain.AdditionalInformationRepository
}

// NewRepository will create an object that represent all repos interface
func NewRepository(conn *utils.Conn, cfg *config.Config, logrus *utils.Logrus) *Repository {
	return &Repository{
		CategoryRepo:           _categoryRepo.NewMysqlCategoryRepository(conn.Mysql),
		NewsRepo:               _newsRepo.NewMysqlNewsRepository(conn.Mysql, logrus),
		InformationRepo:        _informationRepo.NewMysqlInformationRepository(conn.Mysql),
		UnitRepo:               _unitRepo.NewMysqlUnitRepository(conn.Mysql),
		AreaRepo:               _areaRepo.NewMysqlAreaRepository(conn.Mysql),
		EventRepo:              _eventRepo.NewMysqlEventRepository(conn.Mysql),
		FeedbackRepo:           _feedbackRepo.NewMysqlFeedbackRepository(conn.Mysql),
		FeaturedProgramRepo:    _featuredProgramRepo.NewMysqlFeaturedProgramRepository(conn.Mysql),
		UserRepo:               _userRepo.NewMysqlUserRepository(conn.Mysql),
		DataTagsRepo:           _dataTagRepo.NewMysqlDataTagRepository(conn.Mysql),
		TagRepo:                _tagRepo.NewMysqlTagRepository(conn.Mysql),
		SearchRepo:             _searchRepo.NewElasticSearchRepository(conn.Elastic),
		ServicePublicRepo:      _servicePublicRepo.NewMysqlServicePublicRepository(conn.Mysql),
		GeneralInformationRepo: _generalInformationRepo.NewMysqlGeneralInformationRepository(conn.Mysql),
		RoleRepo:               _roleRepo.NewMysqlRoleRepository(conn.Mysql),
		RolePermRepo:           _rolePermRepo.NewMysqlRolePermissionRepository(conn.Mysql),
		TemplateRepo:           _templateRepo.NewMysqlMailTemplateRepository(conn.Mysql),
		RegInvitationRepo:      _regInvitationRepo.NewMysqlRegInvitationRepository(conn.Mysql),
		AwardRepo:              _awardRepo.NewMysqlAwardRepository(conn.Mysql),
		DistrictRepo:           _districtRepo.NewMysqlDistrictRepository(conn.Mysql),
		DocumentArchiveRepo:    _documentArchiveRepo.NewMysqlDocumentArchiveRepository(conn.Mysql),
		PublicServiceRepo:      _publicServiceRepo.NewMysqlPublicServiceRepository(conn.Mysql),
		ExternalVisitorRepo:    _externalVisitorRepo.NewExternalVisitorRepository(cfg),
		PopUpBannerRepo:        _popUpBannerRepo.NewMysqlPopUpBannerRepository(conn.Mysql),
		GovernmentAffairRepo:   _governmentAffairRepo.NewMysqlGovernmentAffairRepository(conn.Mysql),
		SpbeRalsRepo:           _spbeRalsRepo.NewMysqlSpbeRalsRepository(conn.Mysql),
		UptdCabdinRepo:         _uptdCabdinRepo.NewMysqlUptdCabdinRepository(conn.Mysql),
		MasterDataServiceRepo:  _masterDataServiceRepo.NewMysqlMasterDataServiceRepository(conn.Mysql),
		MainServiceRepo:        _mainServiceRepo.NewMysqlMainServiceRepository(conn.Mysql),
		ApplicationRepo:        _applicationRepo.NewMysqlApplicationRepository(conn.Mysql),
		AdditionalInfRepo:      _additionalInformationRepo.NewMysqlAdditionalInformationRepository(conn.Mysql),
	}
}
