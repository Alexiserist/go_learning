package models

type Testing struct {
    ID    int    `json:"id" `
	Username string `json:"username"  binding:"required"`
	Email string `json:"email"  binding:"required"`
}