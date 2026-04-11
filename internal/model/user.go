package model

import (
	"learningSpace/internal/entity"

	"github.com/google/uuid"
)

type UserRegister struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserResponse struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func ToUserResponse(User entity.User) UserResponse {
	return UserResponse{
		Id:       User.Id,
		Name:     User.Name,
		Email:    User.Email,
		Password: User.Password,
	}
}

var (
	UserRole  = "user"
	AdminRole = "admin"
)
