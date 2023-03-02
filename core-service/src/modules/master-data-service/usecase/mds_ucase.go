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
func NewMasterDataServiceUsecase(mdsArgs domain.MasterDataServiceUsecaseArgs) domain.MasterDataServiceUsecase {
	return &masterDataServiceUsecase{
		mdsRepo:        mdsArgs.MdsRepo,
		msRepo:         mdsArgs.MsRepo,
		apRepo:         mdsArgs.ApRepo,
		aiRepo:         mdsArgs.AiRepo,
		cfg:            mdsArgs.Cfg,
		contextTimeout: mdsArgs.ContextTimeout,
	}
}

func (n *masterDataServiceUsecase) Store(ctx context.Context, au *domain.JwtCustomClaims, mds *domain.StoreMasterDataService) (err error) {
	tx, err := n.mdsRepo.GetTx(ctx)
	if err != nil {
		return
	}

	// storing mds support entuty
	n.storeMdsSupport(ctx, mds) // use priv func to reduce exceeding return statements

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
func (n *masterDataServiceUsecase) storeMdsSupport(ctx context.Context, mds *domain.StoreMasterDataService) {
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
}
