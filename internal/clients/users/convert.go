package users

import (
	"github.com/Rasikrr/learning_platform_gateway/internal/domain/entity"
	pb "github.com/Rasikrr/learning_platform_gateway/pkg/deps/api/proto/users"
)

func convert(reply *pb.GetByEmailReply) *entity.User {
	return &entity.User{
		FirstName: reply.FirstName,
		LastName:  reply.LastName,
	}
}
