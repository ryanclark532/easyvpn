package utils

import (
	"easyvpn/src/dtos"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

var (
	secretKey = []byte("your-secret-key")
)

func CreateToken(user dtos.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["iat"] = time.Now().Unix()
	claims["user_id"] = user.ID
	claims["is_admin"] = strconv.FormatBool(user.IsAdmin)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*dtos.CheckTokenResponse, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		isAdminClaim, ok := claims["is_admin"].(string)
		if !ok {
			return nil, fmt.Errorf("is_admin claim not found or not a string")
		}

		isAdmin, err := strconv.ParseBool(isAdminClaim)
		if err != nil {
			return nil, err
		}

		return &dtos.CheckTokenResponse{
			IsAdmin:    isAdmin,
			TokenValid: true,
		}, nil
	}
	return &dtos.CheckTokenResponse{
		IsAdmin:    false,
		TokenValid: false,
	}, nil
}
