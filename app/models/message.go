package models

type Message struct {
	Email		string	`json:"email" validate:"required"`
	Username	string	`json:"username" validate:"required"`
	Messages	string	`json:"messages" validate:"required"`
} 
