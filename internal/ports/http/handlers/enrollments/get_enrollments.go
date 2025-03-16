package enrollments

import (
	"github.com/Rasikrr/learning_platform_core/api"
	"github.com/Rasikrr/learning_platform_core/http/session"
	"net/http"
)

// @Summary get user enrollments
// @Description get user enrollments
// @Tags enrollments
// @Produce json
// @Security     BearerAuth
// @param Authorization header string true "Authorization token"
// @Success 200 {object} getUserEnrollmentsResponse "Success"
// @Router /api/v1/courses/enrollments [get]
func (c *Controller) getUserEnrollments(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ses, err := session.GetFromCtx(ctx)
	if err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}
	enrollments, err := c.enrollmentsService.GetUserEnrollments(ctx, ses.UserID())
	if err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}
	api.SendData(w, convertToGetUserEnrollmentsResponse(enrollments), http.StatusOK)
}
