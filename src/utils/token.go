package utils

import (
	"easyvpn/src/dtos"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	secretKey = []byte("your-secret-key")
)

func CreateToken(userID uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["iat"] = time.Now().Unix()
	claims["user_id"] = userID

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*dtos.CheckTokenResponse, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, HandleError("unexpected signing method", "VerifyToken")
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, HandleError(err.Error(), "VerifyToken")

	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["user_id"])
		//TODO add db check if userid is admin
		response := &dtos.CheckTokenResponse{
			IsAdmin:    true,
			TokenValid: true,
		}
		return response, nil
	}

	response := &dtos.CheckTokenResponse{
		IsAdmin:    false,
		TokenValid: false,
	}
	return response, nil
}
