package services

import (
	"log/slog"
	"regexp"
	"time"

	"github.com/amir-r-z-a/cubic-back/config"
	"github.com/amir-r-z-a/cubic-back/models"
	"github.com/amir-r-z-a/cubic-back/repos"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserService struct {
	Repo      *repos.UserRepo
	Logger    *slog.Logger
	SecretKey []byte
}

func (us *UserService) createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 500).Unix(),
		})

	tokenString, err := token.SignedString(us.SecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func NewUserService(userRepo *repos.UserRepo, appConf *config.AppConfig, secret []byte) *UserService {
	return &UserService{Repo: userRepo, Logger: appConf.Logger, SecretKey: secret}
}

func (us UserService) SignUp(c *gin.Context) {

	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

	signUpStruct := struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required,min=6"`
	}{}

	if err := c.ShouldBindJSON(&signUpStruct); err != nil {
		us.Logger.Error("Failed to bind request body", "error", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if !emailRegex.MatchString(signUpStruct.Username) {
		us.Logger.Error("Invalid email format", "username", signUpStruct.Username)
		c.JSON(400, gin.H{"error": "Invalid email format"})
		return
	}

	res, err := us.Repo.CreateUser(signUpStruct.Username, signUpStruct.Password)

	if err != nil {
		us.Logger.Error("Failed to create user", "error", err)
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}

	token, err := us.createToken(signUpStruct.Username)
	if err != nil {
		us.Logger.Error("Failed to create token", "error", err)
		c.JSON(500, gin.H{"error": "Failed to create token"})
		return
	}

	c.JSON(201, gin.H{"token": token})
	us.Logger.Info("User created successfully", "id", res)
}

func (us UserService) GetUser(c *gin.Context) {
	claims := c.MustGet("claims").(jwt.MapClaims)

	username := claims["username"].(string)

	user, getUserErr := us.Repo.GetUser(username)
	if getUserErr != nil {
		c.JSON(500, gin.H{"error": "Failed to get user"})
		us.Logger.Error("Failed to get user", "error", getUserErr)
		return
	}

	c.JSON(200, gin.H{"user": map[string]interface{}{"username": user.Username, "id": user.ID}})
	us.Logger.Info("User retrieved successfully", "id", user.ID)
}

func (us UserService) SignIn(c *gin.Context) {
	signInStruct := struct {
		Username string `json:"username" binding:"required,min=4"`
		Password string `json:"password" binding:"required,min=6"`
	}{}

	if err := c.ShouldBindJSON(&signInStruct); err != nil {
		us.Logger.Error("Failed to bind request body", "error", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, getUserErr := us.Repo.GetUser(signInStruct.Username)

	if getUserErr != nil {
		c.JSON(401, gin.H{"error": "Invalid username or password"})
		us.Logger.Error("Failed to get user", "error", getUserErr)
		return
	}

	if !models.VerifyPassword(signInStruct.Password, user.PasswordHash) {
		c.JSON(401, gin.H{"error": "Invalid username or password"})
		us.Logger.Error("Invalid username or password")
		return
	}

	token, err := us.createToken(signInStruct.Username)
	if err != nil {
		us.Logger.Error("Failed to create token", "error", err)
		c.JSON(500, gin.H{"error": "Failed to create token"})
		return
	}

	c.JSON(200, gin.H{"token": token})
	us.Logger.Info("User signed in successfully", "id", user.ID)
}

func (us UserService) UpdateUserProfile(c *gin.Context) {
	claims := c.MustGet("claims").(jwt.MapClaims)
	userID := int(claims["user_id"].(int))

	var input models.UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		us.Logger.Error("Failed to bind request body", "error", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := us.Repo.UpdateUser(userID, input)
	if err != nil {
		us.Logger.Error("Failed to update user", "error", err)
		c.JSON(500, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(200, gin.H{"message": "Profile updated successfully"})
	us.Logger.Info("User profile updated successfully", "id", userID)
}

func (us UserService) GetUserProfile(c *gin.Context) {
	claims := c.MustGet("claims").(jwt.MapClaims)
	userID := int(claims["user_id"].(int))

	user, err := us.Repo.GetUserProfile(userID)
	if err != nil {
		us.Logger.Error("Failed to get user profile", "error", err)
		c.JSON(500, gin.H{"error": "Failed to get user profile"})
		return
	}

	c.JSON(200, gin.H{
		"profile": map[string]interface{}{
			"id":              user.ID,
			"username":        user.Username,
			"name":            user.Name,
			"last_name":       user.LastName,
			"age":             user.Age,
			"gender":          user.Gender,
			"height":          user.Height,
			"weight":          user.Weight,
			"disease_history": user.DiseaseHistory,
		},
	})
	us.Logger.Info("User profile retrieved successfully", "id", user.ID)
}
