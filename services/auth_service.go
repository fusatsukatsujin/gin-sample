package services

import (
	"database/sql"
	"gin-sample/pkg/setting"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	db *sql.DB
}

func NewAuthService(db *sql.DB) *AuthService {
	return &AuthService{db: db}
}

func (s *AuthService) Authenticate(username, password string) (string, error) {
	// TODO: ここでデータベースを使用したユーザー認証を実装する
	if username == "admin" && password == "password" {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": 1,
			"exp":     time.Now().Add(time.Hour * 24).Unix(),
		})

		tokenString, err := token.SignedString([]byte(setting.AppSetting.JwtSecret))
		if err != nil {
			return "", err
		}

		return tokenString, nil
	}

	return "", nil
}