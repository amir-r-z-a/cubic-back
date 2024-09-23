package repos

import (
	
	"github.com/amir-r-z-a/cubic-back/models"
)

type UserRepo struct{
	Repo *AppRepo
}

func NewUserRepo (appRep *AppRepo) *UserRepo {
	return &UserRepo{
		Repo: appRep,
	}
}

type UserRepoInterface interface{
	createUser(username string, password string) error
}

func (ur UserRepo) CreateUser(username string, password string) (int, error) {

	passHash, hashErr := models.HashPassword(password)

	if  hashErr != nil {
		return 0, hashErr
	}
	
	user := models.User{
		Username: username,
		PasswordHash: passHash,
	}

	result := ur.Repo.DB.Create(&user)

	if result.Error != nil {
		return 0, result.Error
	}
	return user.ID, nil
}

func (ur UserRepo) GetUser(username string) (models.User, error) {
	user := models.User{}
	result := ur.Repo.DB.Where("username = ?", username).First(&user)

	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}
