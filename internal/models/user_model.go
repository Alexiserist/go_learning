package models

type User struct {
	ID       uint  `json:"id" gorm:"column:ID;type:int4;"`
	Username string `json:"username" gorm:"column:Username"`
	Email    string `json:"email" gorm:"column:Email"`
}
