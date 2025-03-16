package enrollments

import (
	"github.com/Rasikrr/learning_platform_core/api"
	"github.com/Rasikrr/learning_platform_core/http/session"
	"net/http"
)

// @Summary enroll to course
// @Description enroll to course by course id
// @Tags enrollments
// @Accept json
// @Produce json
// @Security     BearerAuth
// @param Authorization header string true "Authorization token"
// @Param  request body enrollToCourseRequest true "request"
// @Success 200 {object} api.EmptySuccessResponse "Success"
// @Router /api/v1/courses/enroll [post]
func (c *Controller) enrollToCourse(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ses, err := session.GetFromCtx(ctx)
	if err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}
	var req enrollToCourseRequest
	if err := api.GetData(r, &req); err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	if err := c.enrollmentsService.Enroll(ctx, ses.UserID(), req.CourseID); err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	api.SendData(w, api.NewEmptySuccessResponse(), http.StatusOK)
}
