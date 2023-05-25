package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
)

type masterDataPublicationUsecase struct {
	mdpRepo        domain.MasterDataPublicationRepository
	mdsRepo        domain.MasterDataServiceRepository
	msRepo         domain.MainServiceRepository
	apRepo         domain.ApplicationRepository
	userRepo       domain.UserRepository
	cfg            *config.Config
	contextTimeout time.Duration
}

// NewMasterDataPublicationUsecase creates a new master-data-publication usecase
func NewMasterDataPublicationUsecase(pubArgs domain.MasterDataPublicationUsecaseArgs) domain.MasterDataPublicationUsecase {
	return &masterDataPublicationUsecase{
		mdpRepo:        pubArgs.PubRepo,
		mdsRepo:        pubArgs.MdsRepo,
		msRepo:         pubArgs.MsRepo,
		apRepo:         pubArgs.ApRepo,
		userRepo:       pubArgs.UserRepo,
		cfg:            pubArgs.Cfg,
		contextTimeout: pubArgs.ContextTimeout,
	}
}

func (n *masterDataPublicationUsecase) Store(ctx context.Context, body *domain.StoreMasterDataPublication) (err error) {
	// begin db transaction
	tx, err := n.mdpRepo.GetTx(ctx)
	if err != nil {
		return
	}

	// get id existing mds
	mds, err := n.mdsRepo.GetByID(ctx, body.DefaultInformation.MdsID)
	if err != nil {
		return errors.New("Your Master Data ID Is Not Exist.")
	}

	mdsBody := domain.StoreMasterDataService{}
	mdsBody.Services.Information.Benefits = body.DefaultInformation.Benefits
	mdsBody.Services.Information.Facilities = body.DefaultInformation.Facilities
	mdsBody.Services.ServiceDetail.TermsAndConditions = body.ServiceDescription.TermsAndConditions
	mdsBody.Services.ServiceDetail.ServiceProcedures = body.ServiceDescription.ServiceProcedures

	// completed existing main_service domain fields
	if err = n.msRepo.UpdateFromPublication(ctx, mds.MainService.ID, &mdsBody, tx); err != nil {
		return
	}

	// completed existing applications domain fields
	mdsBody.Application.Name = body.ServiceDescription.Application.Name
	mdsBody.Application.Status = body.ServiceDescription.Application.Status
	mdsBody.Application.Features = body.ServiceDescription.Application.Features
	mdsBody.Application.Title = body.ServiceDescription.Application.Title
	if err = n.apRepo.Update(ctx, mds.Application.ID, &mdsBody, tx); err != nil {
		return
	}

	// insert for new row mdp
	if err = n.mdpRepo.Store(ctx, body); err != nil {
		return
	}

	// transaction commit
	if err = tx.Commit(); err != nil {
		return
	}

	return
}

func (n *masterDataPublicationUsecase) Fetch(c context.Context, au *domain.JwtCustomClaims, params *domain.Request) (
	res []domain.MasterDataPublication, total int64, err error) {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	params = filterByRoleAcces(au, params)

	res, total, err = n.mdpRepo.Fetch(ctx, params)
	if err != nil {
		return nil, 0, err
	}

	return
}

func (u *masterDataPublicationUsecase) Delete(c context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	if err = u.mdpRepo.Delete(ctx, id); err != nil {
		return
	}

	return
}

func (u *masterDataPublicationUsecase) GetByID(c context.Context, id int64) (res domain.MasterDataPublication, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err = u.mdpRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	resCreatedBy, err := u.userRepo.GetByID(ctx, res.CreatedBy.ID)
	if err != nil {
		return
	}
	res.CreatedBy = resCreatedBy

	return
}

func filterByRoleAcces(au *domain.JwtCustomClaims, params *domain.Request) *domain.Request {

	if params.Filters == nil {
		params.Filters = map[string]interface{}{}
	}

	if au.Role.ID == domain.RoleContributor {
		params.Filters["created_by"] = au.ID
	} else if helpers.IsAdminOPD(au) {
		params.Filters["unit_id"] = au.Unit.ID
	}

	return params
}

func (n *masterDataPublicationUsecase) TabStatus(ctx context.Context, au *domain.JwtCustomClaims, params *domain.Request) (res []domain.TabStatusResponse, err error) {
	params = filterByRoleAcces(au, params)
	res, err = n.mdpRepo.TabStatus(ctx, params)
	if err != nil {
		return
	}
	return
}

func (n *masterDataPublicationUsecase) Update(ctx context.Context, body *domain.StoreMasterDataPublication, pubID int64) (err error) {
	// begin db transaction
	tx, err := n.mdpRepo.GetTx(ctx)
	if err != nil {
		return
	}

	// get id existing mds
	mds, err := n.mdsRepo.GetByID(ctx, body.DefaultInformation.MdsID)
	if err != nil {
		return errors.New("Your Master Data ID Is Not Exist.")
	}

	mdsBody := domain.StoreMasterDataService{}
	mdsBody.Services.Information.Benefits = body.DefaultInformation.Benefits
	mdsBody.Services.Information.Facilities = body.DefaultInformation.Facilities
	mdsBody.Services.ServiceDetail.TermsAndConditions = body.ServiceDescription.TermsAndConditions
	mdsBody.Services.ServiceDetail.ServiceProcedures = body.ServiceDescription.ServiceProcedures

	// completed existing main_service domain fields
	if err = n.msRepo.UpdateFromPublication(ctx, mds.MainService.ID, &mdsBody, tx); err != nil {
		return
	}

	// completed existing applications domain fields
	mdsBody.Application.Name = body.ServiceDescription.Application.Name
	mdsBody.Application.Status = body.ServiceDescription.Application.Status
	mdsBody.Application.Features = body.ServiceDescription.Application.Features
	mdsBody.Application.Title = body.ServiceDescription.Application.Title
	if err = n.apRepo.Update(ctx, mds.Application.ID, &mdsBody, tx); err != nil {
		return
	}

	// update for existing publication
	body.DefaultInformation.MdsID = mds.ID
	if err = n.mdpRepo.Update(ctx, body, pubID); err != nil {
		return
	}

	// transaction commit
	if err = tx.Commit(); err != nil {
		return
	}

	return
}
