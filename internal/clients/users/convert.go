package users

import (
	coreEnum "github.com/Rasikrr/learning_platform_core/enum"
	coreErrors "github.com/Rasikrr/learning_platform_core/errors"
	"github.com/Rasikrr/learning_platform_core/grpc/converters"
	"github.com/Rasikrr/learning_platform_gateway/internal/domain/entity"
	"github.com/Rasikrr/learning_platform_gateway/internal/domain/enum"
	pb "github.com/Rasikrr/learning_platform_gateway/pkg/deps/api/proto/users"
)

func convertUser(user *pb.User) (*entity.User, error) {
	if user == nil {
		return nil, coreErrors.ErrNotFound
	}
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

func convertEnrollments(in ...*pb.Enrollment) ([]*entity.Enrollment, error) {
	enrollments := make([]*entity.Enrollment, 0, len(in))
	for _, enrollment := range in {
		e, err := convertEnrollment(enrollment)
		if err != nil {
			return nil, err
		}
		if e != nil {
			enrollments = append(enrollments, e)
		}
	}
	return enrollments, nil
}

func convertEnrollment(in *pb.Enrollment) (*entity.Enrollment, error) {
	status, err := enum.CourseProgressString(in.GetStatus())
	if err != nil {
		return nil, err
	}
	return &entity.Enrollment{
		ID:     in.Id,
		UserID: in.UserId,
		Course: &entity.Course{
			ID: in.CourseId,
		},
		Status:    status,
		CreatedAt: converters.ConvertToTime(in.CreatedAt),
		UpdatedAt: converters.ConvertToTime(in.UpdatedAt),
	}, nil
}
