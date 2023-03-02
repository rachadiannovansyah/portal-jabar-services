package usecase

import (
	"context"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type masterDataServiceUsecase struct {
	mdsRepo        domain.MasterDataServiceRepository
	msRepo         domain.MainServiceRepository
	apRepo         domain.ApplicationRepository
	aiRepo         domain.AdditionalInformationRepository
	cfg            *config.Config
	contextTimeout time.Duration
}

// NewMasterDataServiceUsecase creates a new master-data-service usecase
func NewMasterDataServiceUsecase(mds domain.MasterDataServiceRepository, ms domain.MainServiceRepository, ap domain.ApplicationRepository, ai domain.AdditionalInformationRepository, cfg *config.Config, timeout time.Duration) domain.MasterDataServiceUsecase {
	return &masterDataServiceUsecase{
		mdsRepo:        mds,
		msRepo:         ms,
		apRepo:         ap,
		aiRepo:         ai,
		cfg:            cfg,
		contextTimeout: timeout,
	}
}

func (n *masterDataServiceUsecase) Store(ctx context.Context, au *domain.JwtCustomClaims, mds *domain.StoreMasterDataService) (err error) {
	tx, err := n.mdsRepo.GetTx(ctx)
	if err != nil {
		return
	}

	// store main_services repository
	msID, err := n.msRepo.Store(ctx, mds)
	if err != nil {
		return
	}
	mds.Services.ID = msID

	// store applications repository
	apID, err := n.apRepo.Store(ctx, mds)
	if err != nil {
		return
	}
	mds.Application.ID = apID

	// store additional_informations repository
	aID, err := n.aiRepo.Store(ctx, mds)
	if err != nil {
		return
	}
	mds.AdditionalInformation.ID = aID

	// store it on service public
	err = n.mdsRepo.Store(ctx, mds, tx)
	if err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		return
	}

	return
}
