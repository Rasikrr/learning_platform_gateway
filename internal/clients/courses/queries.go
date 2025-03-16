package courses

import (
	"context"
	"github.com/Rasikrr/learning_platform_gateway/internal/domain/entity"
	pb "github.com/Rasikrr/learning_platform_gateway/pkg/deps/api/proto/courses"
)

func (c *client) GetCoursesByParams(ctx context.Context, params *entity.GetCoursesParams) ([]*entity.Course, error) {
	out, err := c.client.GetCoursesByParams(ctx, &pb.GetCoursesByParamsRequest{
		Limit:       int64(params.Limit),
		Offset:      int64(params.Offset),
		CategoryIds: params.CategoriesIDs,
	})
	if err != nil {
		return nil, err
	}
	return convertCoursesToEntity(out.Courses)
}

func (c *client) GetCourseByID(ctx context.Context, id string) (*entity.Course, error) {
	out, err := c.client.GetCourseByID(ctx, &pb.GetCourseByIDRequest{
		Id: id,
	})
	if err != nil {
		return nil, err
	}
	return convertCourseToEntity(out.Course)
}

func (c *client) GetCoursesByIDs(ctx context.Context, ids []string) ([]*entity.Course, error) {
	out, err := c.client.GetCoursesByIDs(ctx, &pb.GetCoursesByIDsRequest{
		Ids: ids,
	})
	if err != nil {
		return nil, err
	}
	return convertCoursesToEntity(out.Courses)
}

func (c *client) GetAllCategories(ctx context.Context) ([]*entity.Category, error) {
	out, err := c.client.GetAllCategories(ctx, &pb.GetCategoriesRequest{})
	if err != nil {
		return nil, err
	}
	return convertCategoriesToEntity(out.Categories)
}

func (c *client) GetContentByTopicID(ctx context.Context, courseID, topicID string) (*entity.TopicContent, error) {
	out, err := c.client.GetContentByTopicID(ctx, &pb.GetContentByTopicIDRequest{
		TopicId:  topicID,
		CourseId: courseID,
	})
	if err != nil {
		return nil, err
	}
	return convertContentToEntity(out.Content), nil
}

func (c *client) GetQuizzesByTopicID(ctx context.Context, userID, topicID string) ([]*entity.Quiz, bool, error) {
	out, err := c.client.GetQuizzesByTopicID(ctx, &pb.GetQuizzesByTopicIDRequest{
		UserId:  userID,
		TopicId: topicID,
	})
	if err != nil {
		return nil, false, err
	}
	quizzes := convertQuizzesToEntity(out.Quizzes...)
	return quizzes, out.Passed, nil
}

func (c *client) GetTasksByTopicIDAndOrderNum(
	ctx context.Context,
	id string,
	order int,
	userID string) (*entity.PracticalTask, *entity.TaskSubmission, error) {
	out, err := c.client.GetTasksByTopicIDAndOrderNum(ctx, &pb.GetTasksByTopicIDAndOrderNumRequest{
		Id:     id,
		Order:  int32(order),
		UserId: userID,
	})
	if err != nil {
		return nil, nil, err
	}
	task, err := convertPracticalTaskToEntity(out.Task)
	if err != nil {
		return nil, nil, err
	}
	submission, err := convertTaskSubmissionToEntity(out.Submission)
	if err != nil {
		return nil, nil, err
	}
	return task, submission, nil
}
