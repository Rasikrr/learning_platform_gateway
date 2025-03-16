package courses

import (
	"context"
	"github.com/Rasikrr/learning_platform_core/grpc"
	pb "github.com/Rasikrr/learning_platform_gateway/pkg/deps/api/proto/courses"
	grpc2 "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/Rasikrr/learning_platform_gateway/internal/domain/entity"
)

type Client interface {
	CreateCourse(ctx context.Context, params *entity.CreateCourseParams) error

	GetCoursesByParams(ctx context.Context, params *entity.GetCoursesParams) ([]*entity.Course, error)
	GetCourseByID(ctx context.Context, id string) (*entity.Course, error)
	GetCoursesByIDs(ctx context.Context, ids []string) ([]*entity.Course, error)

	GetAllCategories(ctx context.Context) ([]*entity.Category, error)

	GetContentByTopicID(ctx context.Context, courseID, topicID string) (*entity.TopicContent, error)

	GetQuizzesByTopicID(ctx context.Context, userID, topicID string) ([]*entity.Quiz, bool, error)

	GetTasksByTopicIDAndOrderNum(
		ctx context.Context,
		id string,
		order int,
		userID string) (*entity.PracticalTask, *entity.TaskSubmission, error)
}

type client struct {
	client pb.CoursesClient
}

func NewClient(ctx context.Context, addr string) (Client, error) {
	c, err := grpc.NewClient(
		ctx,
		addr,
		grpc2.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	return &client{client: pb.NewCoursesClient(c)}, nil
}
