package enrollments

import (
	"context"
	"errors"
	coursesC "github.com/Rasikrr/learning_platform_gateway/internal/clients/courses"
	usersC "github.com/Rasikrr/learning_platform_gateway/internal/clients/users"
	"github.com/Rasikrr/learning_platform_gateway/internal/domain/entity"
	"github.com/samber/lo"
)

type Service interface {
	GetUserEnrollments(ctx context.Context, userID string) ([]*entity.Enrollment, error)
	Enroll(ctx context.Context, userID string, courseID string) error
	CheckEnrollment(ctx context.Context, userID string, courseID string) (bool, error)
}

type service struct {
	coursesClient coursesC.Client
	usersClient   usersC.Client
}

func NewService(
	coursesClient coursesC.Client,
	usersClient usersC.Client,
) Service {
	return &service{
		usersClient:   usersClient,
		coursesClient: coursesClient,
	}
}

func (s *service) Enroll(ctx context.Context, userID string, courseID string) error {
	course, err := s.coursesClient.GetCourseByID(ctx, courseID)
	if err != nil {
		return err
	}
	enrolled, err := s.CheckEnrollment(ctx, userID, course.ID)
	if err != nil {
		return err
	}
	if enrolled {
		return errors.New("user already enrolled")
	}
	return s.usersClient.Enroll(ctx, userID, course.ID)
}

func (s *service) CheckEnrollment(ctx context.Context, userID string, courseID string) (bool, error) {
	course, err := s.coursesClient.GetCourseByID(ctx, courseID)
	if err != nil {
		return false, err
	}
	exists, err := s.usersClient.CheckEnrollment(ctx, userID, course.ID)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (s *service) GetUserEnrollments(ctx context.Context, userID string) ([]*entity.Enrollment, error) {
	enrollments, err := s.usersClient.GetUserEnrollments(ctx, userID)
	if err != nil {
		return nil, err
	}
	courses, err := s.coursesClient.GetCoursesByIDs(ctx, lo.Map(enrollments, func(enrollment *entity.Enrollment, _ int) string {
		return enrollment.Course.ID
	}))
	if err != nil {
		return nil, err
	}
	return s.mergeEnrollments(ctx, enrollments, courses)
}

func (s *service) mergeEnrollments(_ context.Context, enrollments []*entity.Enrollment, courses []*entity.Course) ([]*entity.Enrollment, error) {
	coursesMap := lo.SliceToMap(courses, func(course *entity.Course) (string, *entity.Course) {
		return course.ID, course
	})
	for _, enrollment := range enrollments {
		enrollment.Course = coursesMap[enrollment.Course.ID]
	}
	return enrollments, nil
}
