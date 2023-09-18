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

func VerifyToken(tokenString string) dtos.CheckTokenResponse {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return secretKey, nil
	})
	if err != nil {
		return dtos.CheckTokenResponse{
			IsAdmin:    false,
			TokenValid: false,
		}
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		isAdminClaim, ok := claims["is_admin"].(string)
		if !ok {
			fmt.Println("is_admin claim")
			return dtos.CheckTokenResponse{
				IsAdmin:    false,
				TokenValid: false,
			}
		}

		isAdmin, err := strconv.ParseBool(isAdminClaim)
		if err != nil {
			fmt.Println("is_admin conversion")

			return dtos.CheckTokenResponse{
				IsAdmin:    false,
				TokenValid: false,
			}
		}

		passwordExpirationClaim, ok := claims["password_expiry"].(string)
		if !ok {
			fmt.Println(passwordExpirationClaim)
			fmt.Println(ok)
			fmt.Println("passwordExpirationClaim claim")
			return dtos.CheckTokenResponse{
				IsAdmin:    false,
				TokenValid: false,
			}
		}
		date, err := time.Parse(time.DateTime, passwordExpirationClaim)
		if err != nil {
			return dtos.CheckTokenResponse{
				IsAdmin:    false,
				TokenValid: false,
			}
		}

		return dtos.CheckTokenResponse{
			IsAdmin:         isAdmin,
			TokenValid:      true,
			PasswordExpired: date.Before(time.Now()),
		}
	}
	fmt.Println("down here")
	return dtos.CheckTokenResponse{
		IsAdmin:    false,
		TokenValid: false,
	}
}
