package auth

import (
	"go_learning/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


type AuthRepository interface {
	GenerateToken() (string,error)
}

type authRepository struct {}

func NewAuthRepository() AuthRepository {
	return &authRepository{}
}

func (r *authRepository) GenerateToken() (string,error){
	claim := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt: jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim);
	ss,err := token.SignedString(config.GetSecret());
	if err != nil {
		return "",err
	}
	return ss,nil
}