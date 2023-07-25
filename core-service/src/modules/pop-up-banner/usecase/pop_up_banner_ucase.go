package usecase

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/sirupsen/logrus"
)

type popUpBannerUsecase struct {
	popUpBannerRepo domain.PopUpBannerRepository
	cfg             *config.Config
	contextTimeout  time.Duration
}

// NewPopUpBannerUsecase creates a new service-public usecase
func NewPopUpBannerUsecase(pb domain.PopUpBannerRepository, cfg *config.Config, timeout time.Duration) domain.PopUpBannerUsecase {
	return &popUpBannerUsecase{
		popUpBannerRepo: pb,
		cfg:             cfg,
		contextTimeout:  timeout,
	}
}

func (u *popUpBannerUsecase) Fetch(c context.Context, _ *domain.JwtCustomClaims, params *domain.Request) (res []domain.PopUpBanner, total int64, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, total, err = u.popUpBannerRepo.Fetch(ctx, params)
	if err != nil {
		return nil, 0, err
	}

	return
}

func (u *popUpBannerUsecase) GetByID(c context.Context, id int64) (res domain.PopUpBanner, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err = u.popUpBannerRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	return
}

func (u *popUpBannerUsecase) Store(c context.Context, _ *domain.JwtCustomClaims, body *domain.StorePopUpBannerRequest) (err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	// set flag if use scheduler
	body.Scheduler.Status = "NON-ACTIVE"
	if body.Scheduler.IsScheduled == 1 { // 1 is mean true
		body.Scheduler.Status = "ACTIVE"
	}

	if err = u.popUpBannerRepo.Store(ctx, body); err != nil {
		return
	}

	return
}

func (u *popUpBannerUsecase) Delete(c context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	if err = u.popUpBannerRepo.Delete(ctx, id); err != nil {
		return
	}

	return
}

func (n *popUpBannerUsecase) UpdateStatus(ctx context.Context, ID int64, body *domain.UpdateStatusPopUpBannerRequest) (err error) {
	// deactive existing active pop up banner
	if err = n.popUpBannerRepo.DeactiveStatus(ctx); err != nil {
		return
	}

	pop, err := n.popUpBannerRepo.GetByID(ctx, ID)
	if err != nil {
		return
	}

	// update within selected banner to live publish
	body.Duration = pop.Duration
	body.IsLive = int64(0)
	if body.Status == "ACTIVE" {
		body.IsLive = int64(1)
	}
	if err = n.popUpBannerRepo.UpdateStatus(ctx, ID, body); err != nil {
		return
	}

	return
}

func (n *popUpBannerUsecase) Update(ctx context.Context, _ *domain.JwtCustomClaims, ID int64, body *domain.StorePopUpBannerRequest) (err error) {
	// set flag if use scheduler
	body.Scheduler.Status = "NON-ACTIVE"
	if body.Scheduler.IsScheduled == 1 { // 1 is mean true
		body.Scheduler.Status = "ACTIVE"
	}

	if err = n.popUpBannerRepo.Update(ctx, ID, body); err != nil {
		return
	}

	return
}

func (u *popUpBannerUsecase) GetMetaDataImage(_ context.Context, link string) (meta domain.DetailMetaDataImage, err error) {
	subStringsSlice := strings.Split(link, "/")
	fileName := subStringsSlice[len(subStringsSlice)-1]

	resp, err := http.Head(link)
	if err != nil {
		logrus.Error(err)
		return domain.DetailMetaDataImage{}, err
	}

	if resp.StatusCode != http.StatusOK {
		logrus.Error(err)
		return domain.DetailMetaDataImage{}, err
	}

	size, _ := strconv.Atoi(resp.Header.Get("Content-Length"))
	fileSize := int64(size)

	meta.FileName = fileName
	meta.FileDownloadUri = link
	meta.Size = fileSize

	return meta, err
}

func (u *popUpBannerUsecase) LiveBanner(c context.Context) (res domain.PopUpBanner, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err = u.popUpBannerRepo.LiveBanner(ctx)
	if err != nil {
		return
	}

	return
}
