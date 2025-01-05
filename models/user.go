package models

import (
    "golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
    ID                int
    Username         string
    PasswordHash     string
    Name             string
    LastName         string
    Age              *int
    Gender           string
    Height           *float64
    Weight           *float64
    DiseaseHistory   string
    CreatedAt        time.Time
    UpdatedAt        time.Time
}

// Add a new struct for update input
type UpdateUserInput struct {
    Name            string   `json:"name"`
    LastName        string   `json:"last_name"`
    Age             *int     `json:"age"`
    Gender          string   `json:"gender"`
    Height          *float64 `json:"height"`
    Weight          *float64 `json:"weight"`
    DiseaseHistory  string   `json:"disease_history"`
}

type SignUpInput struct {
	Username string `json:"username" binding:"required,min=4"`
	Password string `json:"password" binding:"required,min=6"`
}

type SignInInput struct {
	Username string `json:"username" binding:"required,min=4"`
	Password string `json:"password" binding:"required,min=6"`
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}


func VerifyPassword(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
