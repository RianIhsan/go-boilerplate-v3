package utils

import (
	"ams-sentuh/config"
	"ams-sentuh/internal/entities"
	"ams-sentuh/pkg/httpErrors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
)

type Claim struct {
	ID    uint64
	Email string
	Role  int
	Name  string
	jwt.RegisteredClaims
}

func GenerateJwtToken(user *entities.User, cfg *config.Config, expire time.Duration) (string, error) {
	claims := Claim{
		ID:    user.ID,
		Email: user.Email,
		Role:  int(user.RoleID),
		Name:  user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expire)),
			Issuer:    "jwt",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(cfg.Server.JWTSecretKey))
	if err != nil {
		return "", errors.Wrap(err, "GenerateJWTTokenPair.SignedString")
	}
	return tokenString, nil
}

func GenerateTokenPair(user *entities.User, cfg *config.Config) (accToken, refToken string, err error) {
	accToken, err = GenerateJwtToken(user, cfg, 30*time.Minute) // 30 minute
	if err != nil {
		return
	}
	refToken, err = GenerateJwtToken(user, cfg, 1*24*time.Hour) // 1 day
	return
}

func ValidateJwtToken(tokenString string, cfg *config.Config) (*Claim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		secretKey := []byte(cfg.Server.JWTSecretKey)
		return secretKey, nil
	})

	if err != nil {
		return nil, httpErrors.NewInvalidJwtTokenError(errors.Wrap(err, "ValidateJwtToken.ParseWithClaims"))
	}

	claims, ok := token.Claims.(*Claim)
	if !ok || !token.Valid {
		return nil, httpErrors.NewInternalServerError("unknown claims type, cannot proceed")
	}

	return claims, nil
}
