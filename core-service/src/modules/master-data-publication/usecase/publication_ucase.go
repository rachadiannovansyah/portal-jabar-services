package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type masterDataPublicationUsecase struct {
	mdpRepo        domain.MasterDataPublicationRepository
	mdsRepo        domain.MasterDataServiceRepository
	msRepo         domain.MainServiceRepository
	cfg            *config.Config
	contextTimeout time.Duration
}

// NewMasterDataPublicationUsecase creates a new master-data-publication usecase
func NewMasterDataPublicationUsecase(mdpRepo domain.MasterDataPublicationRepository, mdsRepo domain.MasterDataServiceRepository, msRepo domain.MainServiceRepository, cfg *config.Config, contextTimeout time.Duration) domain.MasterDataPublicationUsecase {
	return &masterDataPublicationUsecase{
		mdpRepo:        mdpRepo,
		mdsRepo:        mdsRepo,
		msRepo:         msRepo,
		cfg:            cfg,
		contextTimeout: contextTimeout,
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

	if err = n.msRepo.UpdateFromPublication(ctx, mds.MainService.ID, &mdsBody, tx); err != nil {
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
