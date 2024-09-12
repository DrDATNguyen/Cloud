package main

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var jwtKey = []byte("your_secret_key")

// Structure to store user information in the token
type Claims struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	SDT   string `json:"SDT"`
	jwt.RegisteredClaims
}

// Initialize the MySQL database connection
func initDB() {
	var err error
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPass, dbHost, dbName)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Could not connect to MySQL: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Could not connect to MySQL: %v", err)
	}
	fmt.Println("Connected to MySQL successfully!")
}

// Generate a JWT token with user information
func generateJWT(id int, email string, SDT string) (string, error) {
	claims := &Claims{
		ID:    id,
		Email: email,
		SDT:   SDT,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
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

func TestJWTFunctions(t *testing.T) {
	// Giả lập thông tin người dùng
	userID := 4
	userEmail := "example@example.com"
	userSDT := "1234567890"

	// Tạo token
	token, err := generateJWT(userID, userEmail, userSDT)
	if err != nil {
		t.Fatalf("Error generating JWT: %v", err)
	}
	fmt.Println("Generated Token:", token)

	// Lấy thông tin người dùng từ token
	claims, err := getUserFromToken(token)
	if err != nil {
		t.Fatalf("Error getting user from token: %v", err)
	}

	// So sánh thông tin người dùng
	if claims.ID != userID {
		t.Errorf("Expected user ID %d, but got %d", userID, claims.ID)
	}
	if claims.Email != userEmail {
		t.Errorf("Expected email %s, but got %s", userEmail, claims.Email)
	}
	if claims.SDT != userSDT {
		t.Errorf("Expected SDT %s, but got %s", userSDT, claims.SDT)
	}
}

// Hàm xử lý yêu cầu đăng nhập
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Xử lý yêu cầu GET với token
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
		var email, pass string

		// Truy vấn thông tin người dùng từ bảng Users
		err = db.QueryRow("SELECT ID, Email, Pass FROM Users WHERE ID = ?", claims.ID).Scan(&id, &email, &pass)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Không tìm thấy người dùng", http.StatusNotFound)
			} else {
				http.Error(w, "Lỗi truy vấn cơ sở dữ liệu", http.StatusInternalServerError)
			}
			return
		}

		// Trả về thông tin người dùng dưới dạng JSON
		userInfo := map[string]interface{}{
			"Email": email,
			"Pass":  pass,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(userInfo)
		return
	} else if r.Method == http.MethodPost {
		// Xử lý yêu cầu POST với email và password
		var requestData struct {
			Email string `json:"email"`
			Pass  string `json:"pass"`
		}

		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			http.Error(w, "Invalid JSON input", http.StatusBadRequest)
			return
		}

		email := requestData.Email
		password := requestData.Pass

		if email == "" || password == "" {
			http.Error(w, "Email và Password là bắt buộc", http.StatusBadRequest)
			return
		}

		var id int
		var storedPassword string
		var token, userName *string

		// Mã hóa mật khẩu nhập vào bằng MD5 ba lần
		hashedPassword := hashPassword(password)

		// Truy vấn cơ sở dữ liệu để kiểm tra thông tin người dùng
		err = db.QueryRow("SELECT ID, UserName, Pass, Token FROM Users WHERE Email = ?", email).Scan(&id, &userName, &storedPassword, &token)
		fmt.Println("err", err)
		if err != nil {
			http.Error(w, "Email hoặc mật khẩu không hợp lệ", http.StatusUnauthorized)
			return
		}
		fmt.Println("storedPassword", storedPassword)
		fmt.Println("hashedPassword", hashedPassword)

		if storedPassword != hashedPassword {
			http.Error(w, "Email hoặc mật khẩu không hợp lệ", http.StatusUnauthorized)
			return
		}

		if token == nil || *token == "" {
			newToken, err := generateJWT(id, email, "")
			if err != nil {
				http.Error(w, "Lỗi khi tạo token", http.StatusInternalServerError)
				return
			}

			_, err = db.Exec("UPDATE Users SET Token = ? WHERE ID = ?", newToken, id)
			if err != nil {
				http.Error(w, "Lỗi khi cập nhật token trong cơ sở dữ liệu", http.StatusInternalServerError)
				return
			}

			token = &newToken
		}

		// Trả về thông tin người dùng dưới dạng JSON
		userInfo := map[string]interface{}{
			"ID":       id,
			"UserName": userName,
			"Email":    email,
			"Token":    *token,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(userInfo)
		return
	}

	// Trả về lỗi nếu phương thức không hợp lệ
	http.Error(w, "Phương thức không được hỗ trợ", http.StatusMethodNotAllowed)
}

// Hàm mã hóa mật khẩu bằng MD5 ba lần
func hashPassword(password string) string {
	hasher := md5.New()
	hasher.Write([]byte(password))
	hash := hasher.Sum(nil)
	for i := 0; i < 2; i++ {
		hasher.Reset()
		hasher.Write(hash)
		hash = hasher.Sum(nil)
	}
	return hex.EncodeToString(hash)
}

// Hàm kiểm tra mật khẩu mạnh
// Hàm kiểm tra mật khẩu mạnh
func isStrongPassword(password string) bool {
	if len(password) < 12 {
		return false
	}

	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#\$%\^&\*\(\)_\+\-=\[\]\{\};:'",.<>/?\\|]`).MatchString(password)

	return hasUpper && hasLower && hasDigit && hasSpecial
}

// Hàm xử lý yêu cầu đăng ký
func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Chỉ cho phép phương thức POST", http.StatusMethodNotAllowed)
		return
	}

	var requestData struct {
		Email string `json:"email"`
		Phone string `json:"phone"`
		Pass  string `json:"pass"`
	}

	// Phân tích dữ liệu JSON từ yêu cầu
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Dữ liệu JSON không hợp lệ", http.StatusBadRequest)
		return
	}

	email := requestData.Email
	phone := requestData.Phone
	password := requestData.Pass

	if email == "" || phone == "" || password == "" {
		http.Error(w, "Email, SĐT và Password là bắt buộc", http.StatusBadRequest)
		return
	}

	// Kiểm tra tính mạnh mẽ của mật khẩu
	if !isStrongPassword(password) {
		http.Error(w, "Mật khẩu phải dài ít nhất 12 ký tự và bao gồm chữ hoa, chữ thường, số và ký hiệu đặc biệt", http.StatusBadRequest)
		return
	}

	// Mã hóa mật khẩu
	hashedPassword := hashPassword(password)

	// Kiểm tra tính duy nhất của email và số điện thoại
	var existingID int
	err = db.QueryRow("SELECT ID FROM Users WHERE Email = ? OR PhoneNumber = ?", email, phone).Scan(&existingID)
	if err == nil {
		http.Error(w, "Email hoặc số điện thoại đã được sử dụng", http.StatusConflict)
		return
	}

	// Tạo token cho người dùng mới
	newID := 0
	token, err := generateJWT(newID, email, phone)
	if err != nil {
		http.Error(w, "Lỗi khi tạo token", http.StatusInternalServerError)
		return
	}

	// Thực hiện việc thêm người dùng vào cơ sở dữ liệu
	_, err = db.Exec("INSERT INTO Users (Email, PhoneNumber, Pass, Token) VALUES (?, ?, ?, ?)", email, phone, hashedPassword, token)
	if err != nil {
		http.Error(w, "Lỗi khi đăng ký người dùng", http.StatusInternalServerError)
		return
	}

	// Trả về thông tin người dùng và thông báo thành công
	userInfo := map[string]interface{}{
		"ID":      newID,
		"Email":   email,
		"Phone":   phone,
		"Token":   token,
		"Message": "Đăng ký thành công",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userInfo)
}

// Hàm xử lý yêu cầu đổi mật khẩu
func changePasswordHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Chỉ cho phép phương thức POST", http.StatusMethodNotAllowed)
		return
	}

	var requestData struct {
		Email       string `json:"email"`
		OldPassword string `json:"old_pass"`
		NewPassword string `json:"new_pass"`
	}

	// Phân tích dữ liệu JSON từ yêu cầu
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Dữ liệu JSON không hợp lệ", http.StatusBadRequest)
		return
	}

	email := requestData.Email
	oldPassword := requestData.OldPassword
	newPassword := requestData.NewPassword

	if email == "" || oldPassword == "" || newPassword == "" {
		http.Error(w, "Email, mật khẩu cũ và mật khẩu mới là bắt buộc", http.StatusBadRequest)
		return
	}

	// Kiểm tra tính mạnh mẽ của mật khẩu mới
	if !isStrongPassword(newPassword) {
		http.Error(w, "Mật khẩu mới phải dài ít nhất 12 ký tự và bao gồm chữ hoa, chữ thường, số và ký hiệu đặc biệt", http.StatusBadRequest)
		return
	}

	// Mã hóa mật khẩu cũ và mới bằng MD5 ba lần
	hashedOldPassword := hashPassword(oldPassword)
	hashedNewPassword := hashPassword(newPassword)

	var id int
	var storedPassword, token, SDT string

	// Truy vấn cơ sở dữ liệu để kiểm tra thông tin người dùng
	err = db.QueryRow("SELECT ID, Pass, Token,PhoneNumber FROM Users WHERE Email = ?", email).Scan(&id, &storedPassword, &token, &SDT)
	fmt.Println("err", err)

	if err != nil {
		http.Error(w, "Email không hợp lệ", http.StatusUnauthorized)
		return
	}

	// Kiểm tra mật khẩu cũ
	if storedPassword != hashedOldPassword {
		http.Error(w, "Mật khẩu cũ không đúng", http.StatusUnauthorized)
		return
	}

	// Cập nhật mật khẩu mới và tạo token mới
	newToken, err := generateJWT(id, email, SDT)
	if err != nil {
		http.Error(w, "Lỗi khi tạo token mới", http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("UPDATE Users SET Pass = ?, Token = ? WHERE ID = ?", hashedNewPassword, newToken, id)
	if err != nil {
		http.Error(w, "Lỗi khi cập nhật mật khẩu hoặc token trong cơ sở dữ liệu", http.StatusInternalServerError)
		return
	}

	// Trả về thông báo thành công
	response := map[string]interface{}{
		"Message": "Đổi mật khẩu thành công",
		"Token":   newToken,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Handle profile request
func profileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Chỉ cho phép phương thức GET", http.StatusMethodNotAllowed)
		return
	}

	// Lấy token từ header Authorization
	tokenString := r.Header.Get("Authorization")
	fmt.Println("tokenString", tokenString)

	if tokenString == "" {
		http.Error(w, "Token không được cung cấp", http.StatusBadRequest)
		return
	}

	// Xác thực token và lấy thông tin người dùng
	claims, err := getUserFromToken(tokenString)
	fmt.Println("err", err)
	if err != nil {
		http.Error(w, "Token không hợp lệ hoặc đã hết hạn", http.StatusUnauthorized)
		return
	}

	// Truy vấn thông tin người dùng từ cơ sở dữ liệu
	var userInfo struct {
		UserName    string  `json:"UserName"`
		Email       string  `json:"Email"`
		PhoneNumber *int    `json:"PhoneNumber"`
		Wallet      float64 `json:"Wallet"`
		Credit      float64 `json:"Credit"`
		Address     string  `json:"Address"`
		VIPuser     string  `json:"VIPuser"`
	}

	err = db.QueryRow("SELECT UserName, Email, PhoneNumber, Wallet, Credit, Address, VIPuser FROM Users WHERE ID = ?", claims.ID).
		Scan(&userInfo.UserName, &userInfo.Email, &userInfo.PhoneNumber, &userInfo.Wallet, &userInfo.Credit, &userInfo.Address, &userInfo.VIPuser)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Không tìm thấy người dùng", http.StatusNotFound)
		} else {
			http.Error(w, "Lỗi truy vấn cơ sở dữ liệu", http.StatusInternalServerError)
		}
		return
	}

	// Trả về thông tin người dùng dưới dạng JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userInfo)
}

func main() {
	initDB()
	defer db.Close()
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/change-password", changePasswordHandler)
	http.HandleFunc("/profile", profileHandler)
	TestJWTFunctions(nil)
	fmt.Println("Server is running on http://localhost:8080...")
	http.ListenAndServe(":8080", nil)
}
