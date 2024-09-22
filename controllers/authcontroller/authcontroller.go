package authcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/jeypc/go-jwt-mux/models"
	"golang.org/x/crypto/bcrypt"
)

// Handler for Login
func Login(w http.ResponseWriter, r *http.Request) {
	// Logic for login
}

// Handler for Register
func Register(w http.ResponseWriter, r *http.Request) {
	var userInput models.User

	// Decode request body
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		http.Error(w, "Gagal memproses JSON input", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Validasi input
	if userInput.Username == "" || userInput.NamaLengkap == "" || userInput.Password == "" {
		http.Error(w, "Data tidak lengkap", http.StatusBadRequest)
		return
	}

	// Hash password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Gagal mengenkripsi password", http.StatusInternalServerError)
		return
	}
	userInput.Password = string(hashPassword)

	// Insert ke database
	if err := models.DB.Create(&userInput).Error; err != nil {
		http.Error(w, "Gagal menyimpan data user", http.StatusInternalServerError)
		return
	}

	// Kirim response sukses
	response, err := json.Marshal(map[string]string{"message": "User registered successfully"})
	if err != nil {
		http.Error(w, "Gagal membuat response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// Handler for Logout
func Logout(w http.ResponseWriter, r *http.Request) {
	// Logic for logout
	w.Write([]byte("Logout successful"))
}
