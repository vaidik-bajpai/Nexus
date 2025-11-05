package helper

import (
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateRandomToken() int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Intn(1000000)
}

func GenerateAccessAndRefreshTokens(userID string) (string, string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(15 * time.Minute),
	})
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(7 * 24 * time.Hour),
	})

	accessTokenString, err := accessToken.SignedString([]byte(GetStrEnvOrPanic("ACCESS_TOKEN_SECRET")))
	if err != nil {
		return "", "", err
	}
	refreshTokenString, err := refreshToken.SignedString([]byte(GetStrEnvOrPanic("REFRESH_TOKEN_SECRET")))
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}
