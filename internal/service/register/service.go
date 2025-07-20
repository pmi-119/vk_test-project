package register

import (
	"context"
	"errors"
	"net/mail"
	"regexp"

	"VK_test_proect/internal/model"

	"github.com/google/uuid"
)

var (
	ErrInvalidEmail = errors.New("invalid email format")
	ErrWeakPassword = errors.New("password must be at least 8 characters, contain a letter and a number")
)

type Service struct {
	user userRepo
}

func New(user userRepo) *Service {
	return &Service{
		user: user,
	}
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func isStrongPassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	hasLetter := regexp.MustCompile(`[A-Za-z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)
	return hasLetter && hasDigit
}

func (s *Service) RegisterUser(ctx context.Context, in In) (model.User, error) {
	if !isValidEmail(in.Email) {
		return model.User{}, ErrInvalidEmail
	}
	if !isStrongPassword(in.Password) {
		return model.User{}, ErrWeakPassword
	}

	id := uuid.New()

	user := model.User{
		Id:       id,
		Login:    in.Login,
		Email:    in.Email,
		Password: in.Password,
	}

	err := s.user.Save(ctx, user)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
