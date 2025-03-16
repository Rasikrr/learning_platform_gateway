package entity

import (
	"github.com/Rasikrr/learning_platform_gateway/internal/domain/enum"
	"time"
)

//go:generate easyjson -all course.go

type Course struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	ImageURL    *string   `json:"image_url"`
	Category    *Category `json:"category"`
	Description string    `json:"description"`
	Topics      []*Topic  `json:"topics,omitempty"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Category struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Topic struct {
	ID             string           `json:"id"`
	CourseID       string           `json:"course_id"`
	Title          string           `json:"title"`
	Description    string           `json:"description"`
	Content        *TopicContent    `json:"content,omitempty"`
	Quizzes        []*Quiz          `json:"submissions,omitempty"`
	PracticalTasks []*PracticalTask `json:"practical_tasks,omitempty"`
	OrderNumber    int              `json:"order_number"`
	CreatedAt      time.Time        `json:"created_at"`
	UpdatedAt      time.Time        `json:"updated_at"`
}

type TopicContent struct {
	ID                  string    `json:"id"`
	TopicID             string    `json:"topic_id"`
	Content             string    `json:"content"`
	AdditionalResources []string  `json:"additional_resources"`
	VideoURLs           []string  `json:"video_urls"`
	ImageURLs           []string  `json:"image_urls"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

type Quiz struct {
	ID             string    `json:"id"`
	TopicID        string    `json:"topic_id"`
	Question       string    `json:"question"`
	Options        []string  `json:"options"`
	CorrectAnswers []bool    `json:"correct_answers"`
	MultipleChoice bool      `json:"multiple_choice"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type PracticalTask struct {
	ID              string                   `json:"id"`
	TopicID         string                   `json:"topic_id"`
	Description     string                   `json:"description"`
	DifficultyLevel enum.Difficulty          `json:"difficulty_level"`
	StarterCode     string                   `json:"starter_code"`
	ExpectedOutput  *string                  `json:"expected_output"`
	OrderNumber     int                      `json:"order_number"`
	CreatedAt       time.Time                `json:"created_at"`
	UpdatedAt       time.Time                `json:"updated_at"`
	TestCases       bool                     `json:"test_cases"`
	Language        enum.ProgrammingLanguage `json:"language"`
}

type AnswerQuiz struct {
	QuestionID string `json:"question_id"`
	Answer     []bool `json:"answer"`
}

type CreateCourseParams struct {
	Title       string  `json:"title"`
	ImageURL    *string `json:"image"`
	CategoryID  string  `json:"category_id"`
	Description string  `json:"description"`
	CreatedBy   string  `json:"created_by"`
}
