package brand

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Service struct {
	log           *slog.Logger
	secretKey     []byte
	tokenDuration time.Duration
}

func NewService(logger *slog.Logger, secretKey []byte, tokenDuration time.Duration) *Service {
	return &Service{
		log:           logger,
		secretKey:     secretKey,
		tokenDuration: tokenDuration,
	}
}

func (s *Service) GenerateToken(userID uint16) (string, error) {
	log := s.log.With(slog.String("op", "services.jwt.generate_token"))

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(s.tokenDuration).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	log.Info("token successfully generated")

	return token.SignedString(s.secretKey)
}

func (s *Service) ParseToken(token string) (jwt.MapClaims, error) {
	log := s.log.With(slog.String("op", "services.jwt.parse_token"))

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return s.secretKey, nil
	})

	if err != nil {
		log.Error("parse error", slog.String("error", err.Error()))
		return nil, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}
