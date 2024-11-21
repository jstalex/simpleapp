package handler

import "simpleapp/internal/domain"

type ResponseErr struct {
	Message string `json:"message,omitempty"`
}

type User struct {
	domain.User
}
