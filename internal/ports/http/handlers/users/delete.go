package users

import (
	"github.com/Rasikrr/learning_platform_core/api"
	"github.com/Rasikrr/learning_platform_core/http/session"
	"net/http"
)

// @Summary Delete my profile
// @Description Delete my profile
// @Tags users
// @Produce json
// @Security     BearerAuth
// @param Authorization header string true "Authorization token"
// @Success 200 {object} api.EmptySuccessResponse "Success"
// @Router /api/v1/users/me/delete [delete]
func (c *Controller) deleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ses, err := session.GetFromCtx(ctx)
	if err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}
	if err := c.usersClient.Delete(ctx, ses.UserID()); err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}
	api.SendData(w, api.NewEmptySuccessResponse(), http.StatusOK)
}
