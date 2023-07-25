package usecase

import (
	"context"
	"database/sql"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
)

type masterDataServiceUsecase struct {
	mdsRepo        domain.MasterDataServiceRepository
	msRepo         domain.MainServiceRepository
	apRepo         domain.ApplicationRepository
	aiRepo         domain.AdditionalInformationRepository
	userRepo       domain.UserRepository
	cfg            *config.Config
	contextTimeout time.Duration
}

// NewMasterDataServiceUsecase creates a new master-data-service usecase
func NewMasterDataServiceUsecase(mdsArgs domain.MasterDataServiceUsecaseArgs) domain.MasterDataServiceUsecase {
	return &masterDataServiceUsecase{
		mdsRepo:        mdsArgs.MdsRepo,
		msRepo:         mdsArgs.MsRepo,
		apRepo:         mdsArgs.ApRepo,
		aiRepo:         mdsArgs.AiRepo,
		userRepo:       mdsArgs.UserRepo,
		cfg:            mdsArgs.Cfg,
		contextTimeout: mdsArgs.ContextTimeout,
	}
}

func (n *masterDataServiceUsecase) Store(ctx context.Context, _ *domain.JwtCustomClaims, mds *domain.StoreMasterDataService) (err error) {
	tx, err := n.mdsRepo.GetTx(ctx)
	if err != nil {
		return
	}

	// storing mds support entity
	if err = n.storeMdsSupport(ctx, mds, tx); err != nil {
		return
	}

	// store it on mds domain
	if err = n.mdsRepo.Store(ctx, mds, tx); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		return
	}

	return
}

// private func to support of mds main_service, application, additional_information entities
func (n *masterDataServiceUsecase) storeMdsSupport(ctx context.Context, mds *domain.StoreMasterDataService, tx *sql.Tx) (err error) {
	// store main_services repository
	msID, err := n.msRepo.Store(ctx, mds, tx)
	if err != nil {
		return
	}
	mds.Services.ID = msID

	// store applications repository
	apID, err := n.apRepo.Store(ctx, mds, tx)
	if err != nil {
		return
	}
	mds.Application.ID = apID

	// store additional_informations repository
	aID, err := n.aiRepo.Store(ctx, mds, tx)
	if err != nil {
		return
	}
	mds.AdditionalInformation.ID = aID

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

func (n *masterDataServiceUsecase) Fetch(c context.Context, au *domain.JwtCustomClaims, params *domain.Request) (
	res []domain.MasterDataService, total int64, err error) {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	params = filterByRoleAcces(au, params)

	res, total, err = n.mdsRepo.Fetch(ctx, params)
	if err != nil {
		return nil, 0, err
	}

	return
}

func (u *masterDataServiceUsecase) Delete(c context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	// check deleteable mds
	delMds, err := u.mdsRepo.CheckHasPublication(ctx, id)
	if err != nil {
		return
	}

	if err = u.mdsRepo.Delete(ctx, delMds.ID); err != nil {
		return
	}

	return
}

func (u *masterDataServiceUsecase) GetByID(c context.Context, id int64) (res domain.MasterDataService, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err = u.mdsRepo.GetByID(ctx, id)
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

func (n *masterDataServiceUsecase) Update(ctx context.Context, body *domain.StoreMasterDataService, mdsID int64) (err error) {
	tx, err := n.mdsRepo.GetTx(ctx)
	if err != nil {
		return
	}

	mds, err := n.GetByID(ctx, mdsID)
	if err != nil {
		return
	}

	// update mds support entity
	if err = n.updateMdsSupport(ctx, mds, body, tx); err != nil {
		return
	}

	// updated it on mds domain
	mdsEntityID := domain.MasterDataServiceEntityID{ // placed here on struct to reduce args code complexity
		ID:                      mdsID,
		MainServiceID:           mds.MainService.ID,
		ApplicationID:           mds.Application.ID,
		AdditionalInformationID: mds.AdditionalInformation.ID,
	}
	if err = n.mdsRepo.Update(ctx, body, &mdsEntityID, tx); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		return
	}

	return
}

// private func to support of mds main_service, application, additional_information entities
func (n *masterDataServiceUsecase) updateMdsSupport(ctx context.Context, mds domain.MasterDataService, body *domain.StoreMasterDataService, tx *sql.Tx) (err error) {
	// update main_services repository
	if err = n.msRepo.Update(ctx, mds.MainService.ID, body, tx); err != nil {
		return
	}

	// update applications repository
	if err = n.apRepo.Update(ctx, mds.Application.ID, body, tx); err != nil {
		return
	}

	// update additional_informations repository
	if err = n.aiRepo.Update(ctx, mds.AdditionalInformation.ID, body, tx); err != nil {
		return
	}

	return
}

func (n *masterDataServiceUsecase) TabStatus(ctx context.Context, au *domain.JwtCustomClaims, params *domain.Request) (res []domain.TabStatusResponse, err error) {
	params = filterByRoleAcces(au, params)
	res, err = n.mdsRepo.TabStatus(ctx, params)
	if err != nil {
		return
	}
	return
}

func (n *masterDataServiceUsecase) Archive(c context.Context, au *domain.JwtCustomClaims, params *domain.Request) (
	res []domain.MasterDataService, err error) {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	params = filterByRoleAcces(au, params)
	params.Filters["status"] = domain.ArchiveStatus

	res, err = n.mdsRepo.Archive(ctx, params)
	if err != nil {
		return nil, err
	}

	return
}
