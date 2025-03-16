package courses

import (
	"context"
	"github.com/Rasikrr/learning_platform_gateway/internal/domain/entity"
	pb "github.com/Rasikrr/learning_platform_gateway/pkg/deps/api/proto/courses"
)

func (c *client) CreateCourse(ctx context.Context, params *entity.CreateCourseParams) error {
	_, err := c.client.CreateCourse(ctx, &pb.CreateCourseRequest{
		Title:       params.Title,
		ImageUrl:    params.ImageURL,
		CategoryId:  params.CategoryID,
		Description: params.Description,
		CreatedBy:   params.CreatedBy,
	})
	return err
}
