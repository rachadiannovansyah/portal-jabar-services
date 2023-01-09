package server

import (
	"log"
	"net/http"

	_goLogMiddl "github.com/jabardigitalservice/golog/http/middleware"
	_goLog "github.com/jabardigitalservice/golog/logger"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	_galleryHttpDelivery "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/media/delivery/http"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/utils"
	"github.com/spf13/viper"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"

	middl "github.com/jabardigitalservice/portal-jabar-services/core-service/src/middleware"

	_areaHttpDelivery "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/area/delivery/http"
	_authHttpDelivery "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/auth/delivery/http"
	_awardHttpDelivery "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/award/delivery/http"
	_districtHttpDelivery "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/district/delivery/http"
	_publicDocumentArchiveHttpDelivery "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/document-archive/delivery/http"
	_eventHttpDelivery "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/event/delivery/http"
	_featuredProgramHttpDelivery "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/featured-program/delivery/http"
	_feedbackHttpDelivery "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/feedback/delivery/http"
	_informationHttpDelivery "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/information/delivery/http"
	_newsHttpDelivery "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/news/delivery/http"
	_publicServiceHttpDelivery "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/public-service/delivery/http"
	_regInvitationHttpDelivery "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/registration-invitation/delivery/http"
	_searchHttpDelivery "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/search/delivery/http"
	_servicePublicHttpDelivery "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/service-public/delivery/http"
	_tagHttpDelivery "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/tag/delivery/http"
	_templateHttpDelivery "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/template/delivery/http"
	_unitHttpDelivery "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/unit/delivery/http"
	_userHttpDelivery "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/user/delivery/http"
	_visitorHttpDelivery "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/visitor/delivery/http"
)

func newAppHandler(e *echo.Echo) {
	cfg := config.NewConfig()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"app":     cfg.App.Name,
			"version": cfg.App.Version,
		})
	})
}

// NewHandler will create a new handler for the given usecase
func NewHandler(cfg *config.Config, apm *utils.Apm, u *Usecases, logger utils.Logrus) {

	e := echo.New()

	isProd := config.LoadAppConfig().Env == "production"
	e.HidePort = isProd
	e.HideBanner = isProd

	e.HTTPErrorHandler = ErrorHandler
	middL := middl.InitMiddleware(cfg, apm.NewRelic, logger)

	v1 := e.Group("/v1")
	r := v1.Group("")
	p := v1.Group("/public")

	goLog := _goLog.Init() // logging with golog package
	e.Use(echo.WrapMiddleware(_goLogMiddl.Logger(goLog, &_goLog.LoggerData{}, false)))

	r.Use(middL.JWT)
	e.Use(middL.NewRelic)
	e.Use(nrecho.Middleware(apm.NewRelic))
	e.Use(middleware.CORSWithConfig(cfg.Cors))

	newAppHandler(e)
	_areaHttpDelivery.NewAreaHandler(v1, r, u.AreaUcase)
	_newsHttpDelivery.NewNewsHandler(v1, r, u.NewsUcase)
	_newsHttpDelivery.NewPublicNewsHandler(p, u.NewsUcase)
	_informationHttpDelivery.NewInformationHandler(v1, r, u.InformationUcase)
	_unitHttpDelivery.NewUnitHandler(v1, r, u.UnitUcase)
	_eventHttpDelivery.NewEventHandler(v1, r, u.EventUcase)
	_eventHttpDelivery.NewPublicEventHandler(p, u.EventUcase)
	_feedbackHttpDelivery.NewFeedbackHandler(v1, r, u.FeedbackUcase)
	_featuredProgramHttpDelivery.NewFeaturedProgramHandler(v1, r, u.FeaturedProgramUcase)
	_authHttpDelivery.NewAuthHandler(v1, r, u.AuthUcase)
	_searchHttpDelivery.NewSearchHandler(v1, r, u.SearchUcase)
	_servicePublicHttpDelivery.NewServicePublicHandler(v1, p, r, u.ServicePublicUcase, apm)
	_userHttpDelivery.NewUserHandler(v1, r, u.UserUsecase)
	_galleryHttpDelivery.NewMediaHandler(v1, r, u.MediaUsecase)
	_tagHttpDelivery.NewTagHandler(v1, r, u.TagUsecase)
	_templateHttpDelivery.NewMailHandler(v1, r, u.TemplateUsecase)
	_regInvitationHttpDelivery.NewRegistrationInvitationHandler(v1, r, u.RegInvitationUsecase)
	_awardHttpDelivery.NewAwardHandler(v1, u.AwardUsecase)
	_districtHttpDelivery.NewDistrictHandler(v1, u.DistrictUsecase)
	_publicDocumentArchiveHttpDelivery.NewPublicDocumentArchiveHandler(p, u.DocumentArchiveUsecase)
	_publicServiceHttpDelivery.NewPublicServiceHandler(v1, p, u.PublicServiceUsecase)
	_visitorHttpDelivery.NewCounterVisitorHandler(p, u.VisitorUsecase)

	log.Fatal(e.Start(viper.GetString("APP_ADDRESS")))
}

// ErrorHandler ...
func ErrorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.Logger().Error(report)
	c.JSON(report.Code, report)
}
