package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() {
	var err error
	// Cấu hình kết nối MySQL
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

// Struct để chứa thông tin yêu cầu đăng nhập
type LoginRequest struct {
	UserName string `json:"username"`
	Pass     string `json:"password"`
}

// Hàm xử lý đăng nhập
func loginHandler(w http.ResponseWriter, r *http.Request) {
	var loginReq LoginRequest

	// Đọc dữ liệu từ yêu cầu đăng nhập
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		http.Error(w, "Yêu cầu không hợp lệ", http.StatusBadRequest)
		return
	}

	var userType, dbPass string

	// Kiểm tra thông tin đăng nhập trong bảng Users
	err = db.QueryRow("SELECT 'user' AS userType, Pass FROM Users WHERE UserName = ? UNION SELECT 'admin', Pass FROM admins WHERE UserName = ?", loginReq.UserName, loginReq.UserName).Scan(&userType, &dbPass)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Tên đăng nhập hoặc mật khẩu không đúng", http.StatusUnauthorized)
		} else {
			http.Error(w, "Lỗi khi truy vấn cơ sở dữ liệu", http.StatusInternalServerError)
		}
		return
	}

	// Kiểm tra mật khẩu
	if loginReq.Pass != dbPass {
		http.Error(w, "Tên đăng nhập hoặc mật khẩu không đúng", http.StatusUnauthorized)
		return
	}

	// Trả về phản hồi đăng nhập thành công
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Đăng nhập thành công", "userType": userType})
}

func main() {
	initDB()
	defer db.Close()

	http.HandleFunc("/login", loginHandler)
	fmt.Println("Server đang chạy tại http://localhost:8080...")
	http.ListenAndServe(":8080", nil)
}
