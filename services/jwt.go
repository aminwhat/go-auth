package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService interface {
	GenerateToken(userId string) (token string, err error)
	ValidateToken(token string) (userId string, err error)
}

type jwtService struct {
	secretKey string
}

func NewJwtService(secretKey string) JwtService {
	return &jwtService{
		secretKey: secretKey,
	}
}

func (j *jwtService) ValidateToken(tokenString string) (userId string, err error) {
	// Parse the token with the claims
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return "", err
	}

	// Validate claims and extract user_id
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if uid, ok := claims["user_id"].(string); ok {
			return uid, nil
		}
		return "", jwt.ErrTokenInvalidClaims
	}

	return "", jwt.ErrTokenInvalidClaims
}

func (j *jwtService) GenerateToken(userId string) (token string, err error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		"iat":     time.Now().Unix(),                     // Issued at time
	}

	// Create a new token with the specified signing method and claims
	theToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := theToken.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
