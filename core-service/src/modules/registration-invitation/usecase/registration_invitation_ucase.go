package usecase

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
)

type regInvitationUcase struct {
	regInvitationRepo domain.RegistrationInvitationRepository
	userRepo          domain.UserRepository
	mailRepo          domain.MailRepository
	mailTemplateRepo  domain.TemplateRepository
	contextTimeout    time.Duration
}

func NewRegInvitationUsecase(regInvitationRepo domain.RegistrationInvitationRepository,
	userRepo domain.UserRepository, mailRepo domain.MailRepository, mtemplateRepo domain.TemplateRepository,
	contextTimeout time.Duration) domain.RegistrationInvitationUsecase {
	return &regInvitationUcase{
		regInvitationRepo: regInvitationRepo,
		userRepo:          userRepo,
		mailRepo:          mailRepo,
		mailTemplateRepo:  mtemplateRepo,
		contextTimeout:    contextTimeout,
	}
}

func (r *regInvitationUcase) generateInvitationToken() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", b), nil
}

func (r *regInvitationUcase) Invite(ctx context.Context,
	req domain.RegistrationInvitation) (regInvitation domain.RegistrationInvitation, err error) {

	// validate if email is already registered in users table
	if u, _ := r.userRepo.GetByEmail(ctx, req.Email); u.Email != "" {
		return regInvitation, errors.New("email already registered")
	}

	// find if email is already registered in registration_invitation table
	regInvitation, _ = r.regInvitationRepo.GetByEmail(ctx, req.Email)
	if regInvitation.Email != "" {
		return regInvitation, errors.New("email already invited")
	}

	// prepare registration invitation data
	regInvitation.Email = req.Email
	regInvitation.UnitID = req.UnitID
	regInvitation.InvitedBy = req.InvitedBy
	regInvitation.Token, _ = r.generateInvitationToken()

	// update invitation if email already exist
	// if regInvitation.ID != nil {
	// 	err = r.regInvitationRepo.Update(ctx, *regInvitation.ID, &regInvitation)
	// } else {
	// 	err = r.regInvitationRepo.Store(ctx, &regInvitation)
	// }

	t, _ := r.mailTemplateRepo.GetByTemplate(ctx, "registration_invitation")
	registrationLink := fmt.Sprintf("%s/daftar?token=%s", config.LoadAppConfig().CmsUrl, regInvitation.Token)

	r.mailRepo.Enqueue(ctx, domain.Mail{
		To:      regInvitation.Email,
		Subject: t.Subject,
		Body:    helpers.ReplaceBodyParams(t.Body, []string{registrationLink}),
	})

	return
}

func (r *regInvitationUcase) Authorize(ctx context.Context,
	token string) (claim domain.RegistrationInvitationClaim, err error) {

	regInvitation, err := r.regInvitationRepo.GetByToken(ctx, token)
	if err != nil {
		return claim, err
	}

	if err := helpers.IsInvitationTokenValid(regInvitation, token); err != nil {
		return claim, err
	}

	claim = domain.RegistrationInvitationClaim{
		Email: regInvitation.Email,
		Token: regInvitation.Token,
	}

	return
}
