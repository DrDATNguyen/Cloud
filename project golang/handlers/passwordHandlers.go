package handlers

import (
	"myproject/initializers"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	jwt.StandardClaims
}

// Hàm kiểm tra mật khẩu mạnh
func IsStrongPassword(password string) bool {
	if len(password) < 12 {
		return false
	}

	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#\$%\^&\*\(\)_\+\-=\[\]\{\};:'",.<>/?\\|]`).MatchString(password)

	return hasUpper && hasLower && hasDigit && hasSpecial
}

// Hàm băm mật khẩu
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// Tạo token JWt
func GenerateJWT(id uint, email, phone string) (string, error) {
	// Thời gian hết hạn Token
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		ID:    id,
		Email: email,
		Phone: phone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(initializers.JwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
