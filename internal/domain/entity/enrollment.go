package entity

import (
	"github.com/Rasikrr/learning_platform_gateway/internal/domain/enum"
	"time"
)

//go:generate easyjson -all course.go

type Enrollment struct {
	ID        string              `json:"id"`
	UserID    string              `json:"user_id"`
	Course    *Course             `json:"course"`
	Status    enum.CourseProgress `json:"status"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
}
