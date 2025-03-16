package courses

import (
	"github.com/Rasikrr/learning_platform_core/grpc/converters"
	"github.com/Rasikrr/learning_platform_gateway/internal/domain/entity"
	"github.com/Rasikrr/learning_platform_gateway/internal/domain/enum"
	pb "github.com/Rasikrr/learning_platform_gateway/pkg/deps/api/proto/courses"
	"github.com/samber/lo"
)

func convertCoursesToEntity(courses []*pb.Course) ([]*entity.Course, error) {
	out := make([]*entity.Course, 0, len(courses))
	for _, course := range courses {
		e, err := convertCourseToEntity(course)
		if err != nil {
			return nil, err
		}
		out = append(out, e)
	}
	return out, nil
}

func convertCourseToEntity(course *pb.Course) (*entity.Course, error) {
	if course == nil {
		return nil, nil
	}
	topics, err := convertTopicsToEntity(course.Topics...)
	if err != nil {
		return nil, err
	}
	return &entity.Course{
		ID:          course.Id,
		Title:       course.Title,
		ImageURL:    course.ImageUrl,
		Category:    convertCategoryToEntity(course.Category),
		Description: course.Description,
		Topics:      topics,
		CreatedBy:   course.CreatedBy,
		CreatedAt:   converters.ConvertToTime(course.CreatedAt),
		UpdatedAt:   converters.ConvertToTime(course.UpdatedAt),
	}, nil
}

func convertCategoriesToEntity(categories []*pb.Category) ([]*entity.Category, error) {
	out := make([]*entity.Category, len(categories))
	for _, category := range categories {
		out = append(out, convertCategoryToEntity(category))
	}
	return out, nil
}

func convertCategoryToEntity(category *pb.Category) *entity.Category {
	if category == nil {
		return nil
	}
	return &entity.Category{
		ID:        category.Id,
		Name:      category.Name,
		CreatedBy: category.CreatedBy,
		CreatedAt: converters.ConvertToTime(category.CreatedAt),
		UpdatedAt: converters.ConvertToTime(category.UpdatedAt),
	}
}

func convertTopicsToEntity(topics ...*pb.Topic) ([]*entity.Topic, error) {
	if len(topics) == 0 {
		return nil, nil
	}
	out := make([]*entity.Topic, 0, len(topics))
	for _, topic := range topics {
		e, err := convertTopicToEntity(topic)
		if err != nil {
			return nil, err
		}
		out = append(out, e)
	}
	return out, nil
}
func convertTopicToEntity(topic *pb.Topic) (*entity.Topic, error) {
	if topic == nil {
		return nil, nil
	}
	tasks, err := convertPracticalTasksToEntity(topic.PracticalTasks...)
	if err != nil {
		return nil, err
	}
	return &entity.Topic{
		ID:             topic.Id,
		CourseID:       topic.CourseId,
		Title:          topic.Title,
		Description:    topic.Description,
		Content:        convertContentToEntity(topic.Content),
		Quizzes:        convertQuizzesToEntity(topic.Quizzes...),
		PracticalTasks: tasks,
		OrderNumber:    int(topic.OrderNumber),
		CreatedAt:      converters.ConvertToTime(topic.CreatedAt),
		UpdatedAt:      converters.ConvertToTime(topic.UpdatedAt),
	}, nil
}

func convertContentToEntity(content *pb.TopicContent) *entity.TopicContent {
	if content == nil {
		return nil
	}
	return &entity.TopicContent{
		ID:                  content.Id,
		TopicID:             content.TopicId,
		Content:             content.Content,
		AdditionalResources: content.AdditionalResources,
		VideoURLs:           content.VideoUrls,
		ImageURLs:           content.ImageUrls,
		CreatedAt:           converters.ConvertToTime(content.CreatedAt),
		UpdatedAt:           converters.ConvertToTime(content.UpdatedAt),
	}
}

func convertQuizzesToEntity(quizzes ...*pb.Quiz) []*entity.Quiz {
	if len(quizzes) == 0 {
		return nil
	}
	return lo.Map(quizzes, func(quiz *pb.Quiz, _ int) *entity.Quiz {
		return &entity.Quiz{
			ID:             quiz.Id,
			TopicID:        quiz.TopicId,
			Question:       quiz.Question,
			Options:        quiz.Options,
			CorrectAnswers: quiz.CorrectAnswers,
			MultipleChoice: quiz.MultipleChoice,
			CreatedAt:      converters.ConvertToTime(quiz.CreatedAt),
			UpdatedAt:      converters.ConvertToTime(quiz.UpdatedAt),
		}
	})
}

func convertPracticalTasksToEntity(tasks ...*pb.PracticalTask) ([]*entity.PracticalTask, error) {
	tasks = lo.Filter(tasks, func(task *pb.PracticalTask, _ int) bool {
		return task != nil
	})
	if len(tasks) == 0 {
		return nil, nil
	}
	t := make([]*entity.PracticalTask, len(tasks))
	for _, task := range tasks {
		out, err := convertPracticalTaskToEntity(task)
		if err != nil {
			return nil, err
		}
		t = append(t, out)
	}
	return t, nil
}

func convertPracticalTaskToEntity(task *pb.PracticalTask) (*entity.PracticalTask, error) {
	if task == nil {
		return nil, nil
	}
	lvl, err := enum.DifficultyString(task.DifficultyLevel)
	if err != nil {
		return nil, err
	}
	lang, err := enum.ProgrammingLanguageString(task.Language)
	if err != nil {
		return nil, err
	}
	return &entity.PracticalTask{
		ID:              task.Id,
		TopicID:         task.TopicId,
		Description:     task.Description,
		DifficultyLevel: lvl,
		StarterCode:     task.StarterCode,
		ExpectedOutput:  task.ExpectedOutput,
		OrderNumber:     int(task.OrderNumber),
		CreatedAt:       converters.ConvertToTime(task.CreatedAt),
		UpdatedAt:       converters.ConvertToTime(task.UpdatedAt),
		TestCases:       task.TestCases,
		Language:        lang,
	}, nil
}

func convertTaskSubmissionToEntity(submission *pb.TaskSubmission) (*entity.TaskSubmission, error) {
	if submission == nil {
		return nil, nil
	}
	return &entity.TaskSubmission{
		ID:        submission.Id,
		TaskID:    submission.TaskId,
		UserID:    submission.UserId,
		Input:     submission.Input,
		Passed:    submission.Passed,
		Error:     submission.Error,
		CreatedAt: converters.ConvertToTime(submission.CreatedAt),
	}, nil
}
