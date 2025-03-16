package queries

import (
	"github.com/Rasikrr/learning_platform_core/api"
	"github.com/Rasikrr/learning_platform_core/http/session"
	"net/http"
)

// @Summary Get course topic tasks
// @Description Get course topic tasks by topic id and. Also you must pass task order
// @Tags tasks
// @Produce json
// @Security     BearerAuth
// @param Authorization header string true "Authorization token"
// @Param course_id path string true "course id"
// @Param topic_id path string true "topic id"
// @Param order path string true "number of task"
// @Success 200 {object} getTaskResponse "Success"
// @Router /api/v1/courses/{course_id}/topic/{topic_id}/tasks/{order} [get]
func (c *Controller) getCourseTopicTasks(w http.ResponseWriter, r *http.Request) {
	var req getTopicTasksRequest
	if err := api.GetData(r, &req); err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	ctx := r.Context()
	ses, err := session.GetFromCtx(ctx)
	if err != nil {
		api.SendError(w, http.StatusUnauthorized, err)
		return
	}
	tasks, solution, err := c.coursesClient.GetTasksByTopicIDAndOrderNum(ctx, req.CourseID, req.TopicID, req.Order, ses.UserID())
	if err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	api.SendData(w, convertToGetTaskResponse(tasks, solution), http.StatusOK)
}
