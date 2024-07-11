package models

type User struct {
	ID       uint  `json:"id" gorm:"primary_key;auto_increment;column:ID;type:int4"`
	Username string `json:"username" gorm:"column:Username" binding:"required"`
	Email    string `json:"email" gorm:"column:Email" binding:"required"`
}

type CreateUser struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
}