package users

import (
	"context"
	"github.com/Rasikrr/learning_platform_core/grpc"
	"github.com/Rasikrr/learning_platform_gateway/internal/domain/entity"
	pb "github.com/Rasikrr/learning_platform_gateway/pkg/deps/api/proto/users"
	grpc2 "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type Client interface {
	GetByID(ctx context.Context, id string) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, params *entity.UpdateUserParams) error
	GetUserEnrollments(ctx context.Context, userID string) ([]*entity.Enrollment, error)
	Enroll(ctx context.Context, userID string, courseID string) error
	CheckEnrollment(ctx context.Context, userID string, courseID string) (bool, error)
}

type client struct {
	client pb.UsersClient
}

func NewClient(ctx context.Context, addr string) (Client, error) {
	conn, err := grpc.NewClient(
		ctx,
		addr,
		grpc2.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return nil, err
	}
	return &client{
		client: pb.NewUsersClient(conn),
	}, nil
}

func (c *client) GetByID(ctx context.Context, id string) (*entity.User, error) {
	reply, err := c.client.GetByID(ctx, &pb.GetByIDRequest{UserId: id})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return convertUser(reply.User)
}

func (c *client) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	reply, err := c.client.GetByEmail(ctx, &pb.GetByEmailRequest{Email: email})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return convertUser(reply.User)
}

func (c *client) Update(ctx context.Context, params *entity.UpdateUserParams) error {
	_, err := c.client.UpdateUser(ctx, &pb.UpdateUserRequest{
		UserId:   params.ID,
		Name:     params.Name,
		LastName: params.LastName,
	})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (c *client) Delete(ctx context.Context, id string) error {
	_, err := c.client.Delete(ctx, &pb.DeleteUserRequest{UserId: id})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (c *client) Enroll(ctx context.Context, userID string, courseID string) error {
	_, err := c.client.Enroll(ctx, &pb.EnrollRequest{
		UserId:   userID,
		CourseId: courseID,
	})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (c *client) GetUserEnrollments(ctx context.Context, userID string) ([]*entity.Enrollment, error) {
	reply, err := c.client.GetUserEnrollments(ctx, &pb.GetUserEnrollmentsRequest{UserId: userID})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return convertEnrollments(reply.Enrollments...)
}

func (c *client) CheckEnrollment(ctx context.Context, userID string, courseID string) (bool, error) {
	reply, err := c.client.CheckEnrollment(ctx, &pb.CheckEnrollmentRequest{
		UserId:   userID,
		CourseId: courseID,
	})
	if err != nil {
		log.Println(err)
		return false, err
	}
	return reply.Enrolled, nil
}
