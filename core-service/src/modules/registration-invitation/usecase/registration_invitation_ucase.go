package usecase

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type regInvitationUcase struct {
	regInvitationRepo domain.RegistrationInvitationRepository
	userRepo          domain.UserRepository
	contextTimeout    time.Duration
}

func NewRegInvitationUsecase(regInvitationRepo domain.RegistrationInvitationRepository,
	userRepo domain.UserRepository, contextTimeout time.Duration) domain.RegistrationInvitationUsecase {
	return &regInvitationUcase{
		regInvitationRepo: regInvitationRepo,
		userRepo:          userRepo,
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
	email string) (regInvitation domain.RegistrationInvitation, err error) {

	// validate if email is already registered in users table
	if u, _ := r.userRepo.GetByEmail(ctx, email); u.Email != "" {
		return regInvitation, errors.New("email already registered")
	}

	// find if email is already registered in registration_invitation table
	regInvitation, err = r.regInvitationRepo.GetByEmail(ctx, email)
	if err != nil {
		return regInvitation, err
	}

	// prepare registration invitation data
	regInvitation.Email = email
	regInvitation.Token, _ = r.generateInvitationToken()
	regInvitation.ExpiredAt = time.Now().Add(time.Hour * 24 * 5) // expired in 5 days

	// update invitation if email already exist
	if regInvitation.ID > 0 {
		err = r.regInvitationRepo.Update(ctx, regInvitation.ID, &regInvitation)
	} else {
		err = r.regInvitationRepo.Store(ctx, &regInvitation)
	}

	// TODO: dispatch job to send email invitation

	return
}
