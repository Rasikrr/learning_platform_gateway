package users

import (
	"github.com/Rasikrr/learning_platform_gateway/internal/domain/entity"
)

//go:generate easyjson -all models.go

type updateUserRequest struct {
	Name     *string `json:"name"`
	LastName *string `json:"last_name"`
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type userResponse struct {
	User User `json:"user"`
}

func convertUserResponse(user *entity.User) userResponse {
	return userResponse{
		User: convertUser(user),
	}
}

func convertUser(user *entity.User) User {
	return User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		//ID:        user.ID.String(),
		//Name:      user.Name,
		//LastName:  user.LastName,
		//Email:     user.Email,
		//Role:      user.AccountRole,
		//CreatedAt: user.CreatedAt,
		//UpdatedAt: user.UpdatedAt,
		//DeletedAt: user.DeletedAt,
	}
}
