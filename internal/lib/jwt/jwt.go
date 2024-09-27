package customjwt

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"kanban/internal/domain/models"
	"time"
)

func GenerateJWTToken(secret string, uuid uuid.UUID) (*models.AuthTokens, error) {
	claims := models.JWTData{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(time.Hour)).Unix(),
		},
		CustomClaims: map[string]any{
			"id": uuid,
		},
	}

	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenString.SignedString([]byte(secret))
	if err != nil {
		return &models.AuthTokens{}, err
	}

	refreshClaims := models.JWTData{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(time.Hour)).Unix(),
		},
		// TODO: Change to user id
		CustomClaims: map[string]any{
			"id": uuid,
		},
	}

	refreshTokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	refreshToken, err := refreshTokenString.SignedString([]byte(secret))
	if err != nil {
		return &models.AuthTokens{}, err
	}

	return &models.AuthTokens{
		AccessToken:  token,
		RefreshToken: refreshToken,
	}, nil
}
