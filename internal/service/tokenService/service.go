package token_service

import (
	"errors"
	"fmt"
	"time"

	"VK_test_proect/internal/model"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const (
	userIDKey     = "user_id"
	expirationKey = "exp"
)

type Service struct {
	secret string
}

func New(secret string) *Service {
	return &Service{secret: secret}
}

func (s *Service) Generate(user model.User) (string, error) {
	claims := jwt.MapClaims{
		userIDKey:     user.Id.String(),
		expirationKey: time.Now().UTC().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(s.secret))
}

func (s *Service) Validate(tokenString string) (uuid.UUID, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.secret), nil
	})
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("cannot parse token string: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.UUID{}, errors.New("invalid token")
	}

	expirationTime, err := getExpirationTimeFromClaims(claims)
	if err != nil {
		return uuid.UUID{}, errors.New("invalid token")
	}

	if isTokenExpired(expirationTime) {
		return uuid.UUID{}, errors.New("invalid token")
	}

	return getUserIDFromClaims(claims)
}

func getExpirationTimeFromClaims(claims jwt.MapClaims) (time.Time, error) {
	untypedExpiration, ok := claims[expirationKey]
	if !ok {
		return time.Time{}, errors.New("invalid token")
	}

	unixExpiration, ok := untypedExpiration.(float64)
	if !ok {
		return time.Time{}, errors.New("invalid token")
	}

	return time.Unix(int64(unixExpiration), 0), nil
}

func isTokenExpired(expirationTime time.Time) bool {
	return time.Now().UTC().After(expirationTime)
}

func getUserIDFromClaims(claims jwt.MapClaims) (uuid.UUID, error) {
	untypedUserID, ok := claims[userIDKey]
	if !ok {
		return uuid.UUID{}, errors.New("invalid token")
	}

	stringUserID, ok := untypedUserID.(string)
	if !ok {
		return uuid.UUID{}, errors.New("invalid token")
	}

	userID, err := uuid.Parse(stringUserID)
	if err != nil {
		return uuid.UUID{}, errors.New("invalid token")
	}

	return userID, nil
}
