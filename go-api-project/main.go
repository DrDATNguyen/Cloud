package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
)

var db *sql.DB
var jwtKey = []byte("44448888") // Khóa bí mật để ký JWT

// Cấu trúc chứa thông tin người dùng trong token
type Claims struct {
	ID          int    `json:"id"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	UserName    string `json:"username"`
	jwt.RegisteredClaims
}

// Hàm khởi tạo kết nối cơ sở dữ liệu MySQL
func initDB() {
	var err error
	dsn := "root:Dat210903@@tcp(127.0.0.1:3306)/cloud"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Không thể kết nối với MySQL:", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Không thể kết nối với MySQL:", err)
	}
	fmt.Println("Kết nối MySQL thành công!")
}

// Hàm tạo JWT chứa thông tin người dùng
func generateJWT(id int, email, phoneNumber, username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token có hiệu lực trong 24 giờ

	claims := &Claims{
		ID:          id,
		Email:       email,
		PhoneNumber: phoneNumber,
		UserName:    username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Hàm lấy thông tin từ token
func getUserFromToken(tokenString string) (Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return *claims, err
	}
	if !token.Valid {
		return *claims, fmt.Errorf("Token không hợp lệ")
	}
	return *claims, nil
}

// Xử lý yêu cầu API để lấy thông tin người dùng
func loginHandler(w http.ResponseWriter, r *http.Request) {
	tokenString := r.URL.Query().Get("token")
	if tokenString == "" {
		http.Error(w, "Token không được cung cấp", http.StatusBadRequest)
		return
	}

	claims, err := getUserFromToken(tokenString)
	if err != nil {
		http.Error(w, "Token không hợp lệ hoặc đã hết hạn", http.StatusUnauthorized)
		return
	}

	var id int
	var email, phoneNumber, userName, pass, address, vipUser string
	var wallet, credit float64

	// Truy vấn thông tin người dùng từ bảng Users
	err = db.QueryRow("SELECT ID, Email, PhoneNumber, UserName, Pass, Wallet, Credit, Address, VIPuser FROM Users WHERE ID = ?", claims.ID).Scan(&id, &email, &phoneNumber, &userName, &pass, &wallet, &credit, &address, &vipUser)
	if err != nil {
		http.Error(w, "Không tìm thấy người dùng", http.StatusNotFound)
		return
	}

	// Trả về thông tin người dùng
	userInfo := map[string]interface{}{
		"ID":          id,
		"Email":       email,
		"PhoneNumber": phoneNumber,
		"UserName":    userName,
		"Pass":        pass,
		"Wallet":      wallet,
		"Credit":      credit,
		"Address":     address,
		"VIPuser":     vipUser,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userInfo)
}

// Tạo token cho tất cả người dùng trong cơ sở dữ liệu
func generateTokensForAllUsers() {
	rows, err := db.Query("SELECT ID, Email, PhoneNumber, UserName FROM Users")
	if err != nil {
		log.Fatal("Lỗi khi truy vấn cơ sở dữ liệu:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var email, phoneNumber, username string
		if err := rows.Scan(&id, &email, &phoneNumber, &username); err != nil {
			log.Println("Lỗi khi đọc dữ liệu:", err)
			continue
		}

		// Tạo token cho từng người dùng
		token, err := generateJWT(id, email, phoneNumber, username)
		if err != nil {
			log.Println("Lỗi khi tạo token cho user:", username, err)
			continue
		}

		fmt.Printf("Token cho user %s: %s\n", username, token)
	}
}

func main() {
	initDB()
	defer db.Close()

	// Tạo token cho tất cả người dùng và in ra màn hình
	generateTokensForAllUsers()

	http.HandleFunc("/login", loginHandler)
	fmt.Println("Server đang chạy tại http://localhost:8080...")
	http.ListenAndServe(":8080", nil)
}
