package users

import (
	coreEnum "github.com/Rasikrr/learning_platform_core/enum"
	"github.com/Rasikrr/learning_platform_core/grpc/converters"
	"github.com/Rasikrr/learning_platform_gateway/internal/domain/entity"
	pb "github.com/Rasikrr/learning_platform_gateway/pkg/deps/api/proto/users"
)

func convert(res *pb.GetByEmailResponse) (*entity.User, error) {
	user := res.GetUser()
	role, err := coreEnum.AccountRoleString(user.GetAccountRole())
	if err != nil {
		return nil, err
	}
	return &entity.User{
		ID:          user.Id,
		Name:        user.Name,
		LastName:    user.LastName,
		Email:       user.Email,
		Password:    user.Password,
		AccountRole: role,
		CreatedAt:   converters.ConvertToTime(user.CreatedAt),
		UpdatedAt:   converters.ConvertToTime(user.UpdatedAt),
		DeletedAt:   converters.ConvertToTimePtr(user.DeletedAt),
	}, nil
}
