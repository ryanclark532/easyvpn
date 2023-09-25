package utils

import (
	"easyvpn/src/auth/auth_dtos"
	user_dtos "easyvpn/src/user/user-dtos"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

var (
	secretKey = []byte("your-secret-key")
)

func CreateToken(user *user_dtos.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()
	claims["iat"] = time.Now().Unix()
	claims["user_id"] = user.ID
	claims["is_admin"] = strconv.FormatBool(user.IsAdmin)
	claims["password_expiry"] = user.PasswordExpiry.Format(time.DateTime)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func CheckUserToken(tokenString string) (*auth_dtos.CheckTokenResponse, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return secretKey, nil
	})
	if err != nil {
		var ve *jwt.ValidationError
		if errors.As(err, &ve) {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return &auth_dtos.CheckTokenResponse{
					IsAdmin:    false,
					TokenValid: false,
				}, nil
			}
		}
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		isAdminClaim, ok := claims["is_admin"].(string)
		if !ok {
			return &auth_dtos.CheckTokenResponse{
				IsAdmin:    false,
				TokenValid: false,
			}, nil
		}

		isAdmin, err := strconv.ParseBool(isAdminClaim)
		if err != nil {

			return &auth_dtos.CheckTokenResponse{
				IsAdmin:    false,
				TokenValid: false,
			}, nil
		}

		passwordExpirationClaim, ok := claims["password_expiry"].(string)
		if !ok {
			return &auth_dtos.CheckTokenResponse{
				IsAdmin:    false,
				TokenValid: false,
			}, nil
		}
		date, err := time.Parse(time.DateTime, passwordExpirationClaim)
		if err != nil {
			return &auth_dtos.CheckTokenResponse{
				IsAdmin:    false,
				TokenValid: false,
			}, nil
		}

		return &auth_dtos.CheckTokenResponse{
			IsAdmin:         isAdmin,
			TokenValid:      true,
			PasswordExpired: date.Before(time.Now()),
		}, nil
	}
	return &auth_dtos.CheckTokenResponse{
		IsAdmin:    false,
		TokenValid: false,
	}, nil
}
