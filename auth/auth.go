package auth

import (
	"go_learning/config"
	"go_learning/database"
	"go_learning/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)


type AuthRepository interface {
	GenerateToken(username string) (string,error)
	EncodingPassword(string) (string,error)
	CompareHashAndPassword(hash, password string) (bool)
	LoginHandler(username string,password string) (UserData,error)
}

type authRepository struct {}

func NewAuthRepository() AuthRepository {
	return &authRepository{}
}

func (r *authRepository) GenerateToken(username string) (string,error){
	claim := &jwt.MapClaims{
		"Username" : string(username),
		"ExpiresAt": jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		"IssuedAt": jwt.NewNumericDate(time.Now()),
		"NotBefore": jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim);
	secret := config.GetSecret();
	ss,err := token.SignedString([]byte(secret));
	if err != nil {
		return "",err
	}
	return ss,nil
}

func (r *authRepository) EncodingPassword(password string) (string,error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost);
	return string(bytes),err	
}

func (r *authRepository) CompareHashAndPassword(hash, password string) (bool){
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}


func (r *authRepository) LoginHandler(username string,	password string) (UserData,error) {
	var user User
	var userData UserData

	query := `SELECT * FROM users WHERE "Username" = ?  AND "IsActive" = true LIMIT 1`
	if err := database.DB.Raw(query, username).Scan(&user).Error; err != nil {
		return userData, utils.ErrInvalidCredentials
	}

	if !r.CompareHashAndPassword(user.Password,password){
		return userData, utils.ErrInvalidCredentials
	}

	token, err := r.GenerateToken(user.Username)
	if err != nil {
		return userData, utils.ErrTokenGeneration
	}

	userData.Username = user.Username
	userData.AccessToken = token
	
	return userData,nil
}