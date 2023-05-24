package server

import (
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/utils"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	_areaUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/area/usecase"
	_authUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/auth/usecase"
	_awardUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/award/usecase"
	_districtUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/district/usecase"
	_documentArchiveUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/document-archive/usecase"
	_eventUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/event/usecase"
	_featuredProgramUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/featured-program/usecase"
	_feedbackUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/feedback/usecase"
	_governmentAffairUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/government-affair/usecase"
	_informationUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/information/usecase"
	_masterDataPublicationUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/master-data-publication/usecase"
	_masterDataServiceUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/master-data-service/usecase"
	_mediaUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/media/usecase"
	_newsUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/news/usecase"
	_popUpBannerUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/pop-up-banner/usecase"
	_publicServiceUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/public-service/usecase"
	_regInvitationUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/registration-invitation/usecase"
	_searchUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/search/usecase"
	_servicePublicUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/service-public/usecase"
	_spbeRalsUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/spbe-rals/usecase"
	_tagUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/tag/usecase"
	_templateUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/template/usecase"
	_unitUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/unit/usecase"
	_uptdCabdinUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/uptd-cabdin/usecase"
	_userUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/user/usecase"
	_visitorUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/visitor/usecase"
)

// Usecases ...
type Usecases struct {
	AreaUcase                    domain.AreaUsecase
	CategoryUcase                domain.CategoryUsecase
	NewsUcase                    domain.NewsUsecase
	InformationUcase             domain.InformationUsecase
	UnitUcase                    domain.UnitUsecase
	EventUcase                   domain.EventUsecase
	FeedbackUcase                domain.FeedbackUsecase
	FeaturedProgramUcase         domain.FeaturedProgramUsecase
	AuthUcase                    domain.AuthUsecase
	SearchUcase                  domain.SearchUsecase
	ServicePublicUcase           domain.ServicePublicUsecase
	UserUsecase                  domain.UserUsecase
	MediaUsecase                 domain.MediaUsecase
	TagUsecase                   domain.TagUsecase
	TemplateUsecase              domain.TemplateUsecase
	RegInvitationUsecase         domain.RegistrationInvitationUsecase
	AwardUsecase                 domain.AwardUsecase
	DistrictUsecase              domain.DistrictUsecase
	DocumentArchiveUsecase       domain.DocumentArchiveUsecase
	PublicServiceUsecase         domain.PublicServiceUsecase
	VisitorUsecase               domain.VisitorUsecase
	PopUpBannerUsecase           domain.PopUpBannerUsecase
	GovernmentAffairUsecase      domain.GovernmentAffairUsecase
	SpbeRalsUsecase              domain.SpbeRalsUsecase
	UptdCabdinUsecase            domain.UptdCabdinUsecase
	MasterDataServiceUsecase     domain.MasterDataServiceUsecase
	MasterDataPublicationUsecase domain.MasterDataPublicationUsecase
}

// NewUcase will create an object that represent all usecases interface
func NewUcase(cfg *config.Config, conn *utils.Conn, r *Repository, timeoutContext time.Duration) *Usecases {
	return &Usecases{
		AreaUcase:               _areaUcase.NewAreaUsecase(r.AreaRepo, timeoutContext),
		NewsUcase:               _newsUcase.NewNewsUsecase(r.NewsRepo, r.CategoryRepo, r.UserRepo, r.TagRepo, r.DataTagsRepo, r.AreaRepo, r.SearchRepo, cfg, timeoutContext),
		InformationUcase:        _informationUcase.NewInformationUsecase(r.InformationRepo, r.CategoryRepo, timeoutContext),
		UnitUcase:               _unitUcase.NewUnitUsecase(r.UnitRepo, timeoutContext),
		EventUcase:              _eventUcase.NewEventUsecase(r.EventRepo, r.CategoryRepo, r.TagRepo, r.DataTagsRepo, r.UserRepo, timeoutContext),
		FeedbackUcase:           _feedbackUcase.NewFeedbackUsecase(r.FeedbackRepo, timeoutContext),
		FeaturedProgramUcase:    _featuredProgramUcase.NewFeaturedProgramUsecase(r.FeaturedProgramRepo, timeoutContext),
		AuthUcase:               _authUcase.NewAuthUsecase(cfg, r.UserRepo, r.UnitRepo, r.RoleRepo, r.RolePermRepo, timeoutContext),
		SearchUcase:             _searchUcase.NewSearchUsecase(r.SearchRepo, cfg, timeoutContext),
		ServicePublicUcase:      _servicePublicUcase.NewServicePublicUsecase(r.ServicePublicRepo, r.GeneralInformationRepo, r.UserRepo, r.SearchRepo, cfg, timeoutContext),
		UserUsecase:             _userUcase.NewUserUsecase(r.UserRepo, r.UnitRepo, r.RoleRepo, r.TemplateRepo, r.RegInvitationRepo, timeoutContext),
		MediaUsecase:            _mediaUcase.NewMediaUsecase(cfg, conn, timeoutContext),
		TagUsecase:              _tagUcase.NewTagUsecase(r.TagRepo, timeoutContext),
		TemplateUsecase:         _templateUcase.NewTemplateUsecase(r.TemplateRepo, r.UserRepo, timeoutContext),
		RegInvitationUsecase:    _regInvitationUcase.NewRegInvitationUsecase(r.RegInvitationRepo, r.UserRepo, r.MailRepo, r.TemplateRepo, timeoutContext),
		AwardUsecase:            _awardUcase.NewAwardUsecase(r.AwardRepo, timeoutContext),
		DistrictUsecase:         _districtUcase.NewDistrictUsecase(r.DistrictRepo, timeoutContext),
		DocumentArchiveUsecase:  _documentArchiveUcase.NewDocumentArchiveUsecase(r.DocumentArchiveRepo, r.UserRepo, cfg, timeoutContext),
		PublicServiceUsecase:    _publicServiceUcase.NewPublicServiceUsecase(r.PublicServiceRepo, r.UserRepo, r.SearchRepo, cfg, timeoutContext),
		VisitorUsecase:          _visitorUcase.NewVisitorUsecase(r.ExternalVisitorRepo, conn, timeoutContext),
		PopUpBannerUsecase:      _popUpBannerUcase.NewPopUpBannerUsecase(r.PopUpBannerRepo, cfg, timeoutContext),
		GovernmentAffairUsecase: _governmentAffairUcase.NewGovernmentAffairUsecase(r.GovernmentAffairRepo, cfg, timeoutContext),
		SpbeRalsUsecase:         _spbeRalsUcase.NewSpbeRalsUsecase(r.SpbeRalsRepo, cfg, timeoutContext),
		UptdCabdinUsecase:       _uptdCabdinUcase.NewUptdCabdinUsecase(r.UptdCabdinRepo, cfg, timeoutContext),
		MasterDataServiceUsecase: _masterDataServiceUcase.NewMasterDataServiceUsecase(domain.MasterDataServiceUsecaseArgs{
			MdsRepo:        r.MasterDataServiceRepo,
			MsRepo:         r.MainServiceRepo,
			ApRepo:         r.ApplicationRepo,
			AiRepo:         r.AdditionalInfRepo,
			UserRepo:       r.UserRepo,
			Cfg:            cfg,
			ContextTimeout: timeoutContext,
		}),
		MasterDataPublicationUsecase: _masterDataPublicationUcase.NewMasterDataPublicationUsecase(domain.MasterDataPublicationUsecaseArgs{
			PubRepo:        r.MasterDataPublicationRepo,
			MdsRepo:        r.MasterDataServiceRepo,
			MsRepo:         r.MainServiceRepo,
			ApRepo:         r.ApplicationRepo,
			UserRepo:       r.UserRepo,
			Cfg:            cfg,
			ContextTimeout: timeoutContext,
		}),
	}
}
