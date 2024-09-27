package customjwt

import (
	"fmt"
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

func ValidateToken(tokenString string, secret []byte) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secret, nil
	})
}
