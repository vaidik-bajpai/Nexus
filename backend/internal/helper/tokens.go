package helper

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	mathrand "math/rand"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
)

func CreateRandomToken() int {
	r := mathrand.New(mathrand.NewSource(time.Now().UnixNano()))
	return r.Intn(1000000)
}

func GenerateAccessAndRefreshTokens(user *types.User) (string, string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": *user,
		"exp":  time.Now().Add(24 * time.Hour).Unix(),
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": *user,
		"exp":  time.Now().Add(7 * 24 * time.Hour).Unix(),
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

func VerifyToken(tokenString string, tokenSecret string) (*types.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Extract user from claims (it's stored as a map[string]interface{})
		userMap, ok := claims["user"].(map[string]interface{})
		if !ok {
			return nil, errors.New("invalid token: user claim not found")
		}

		// Marshal the map to JSON
		userJSON, err := json.Marshal(userMap)
		if err != nil {
			return nil, errors.New("invalid token: failed to marshal user")
		}

		// Unmarshal JSON back into User struct
		var user types.User
		if err := json.Unmarshal(userJSON, &user); err != nil {
			return nil, errors.New("invalid token: failed to unmarshal user")
		}

		return &user, nil
	}

	return nil, errors.New("invalid token")
}

func GenerateOAuthState() (string, error) {
	nonceBytes := make([]byte, 64)
	_, err := io.ReadFull(rand.Reader, nonceBytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(nonceBytes), nil
}
