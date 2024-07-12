package repository

import (
	"go_learning/auth"
	"go_learning/database"
	"go_learning/internal/models"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	CreateUser(user models.User) (models.User, error)
	DeleteUser(user models.User) (error)
	FindOneByKey(id uint) (models.User,error)
}

type userRepository struct{
	authService auth.AuthService
}

func NewUserRepository() UserRepository {
    return &userRepository{
		authService: auth.NewAuthRepository(),
	}
}

func (r *userRepository) FindAll() ([]models.User, error){
	var user []models.User
	if err:= database.DB.Find(&user).Error; err != nil {
		return nil,err
	}
	return user, nil
}

func (r *userRepository) FindOneByKey(id uint) (models.User,error){
	var user models.User
	if err:= database.DB.First(&user,id).Error; err != nil {
		return user,err
	}
	return user,nil
}


func (r *userRepository) CreateUser(user models.User) (models.User,error) {
	encoded, err := r.authService.EncodingPassword(user.Password);
	if err != nil {
		return user,err
	}
	user.Password = encoded;
	if err := database.DB.Save(&user).Error; err != nil {
		return user, err
	}
	return user,nil
}


func (r *userRepository) DeleteUser(user models.User) (error) {
	if err := database.DB.Delete(&user,user).Error; err != nil {
		return err
	}
	return nil
}