package users

import (
	"github.com/Rasikrr/learning_platform_core/api"
	"github.com/Rasikrr/learning_platform_core/http/session"
	"net/http"
)

// @Summary Get my profile
// @Description Get my profile
// @Tags users
// @Produce json
// @Security     BearerAuth
// @param Authorization header string true "Authorization token"
// @Success 200 {object} userResponse "Success"
// @Router /api/v1/users/me [get]
func (c *Controller) getMyProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ses, err := session.GetFromCtx(ctx)
	if err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}
	user, err := c.usersClient.GetByID(ctx, ses.UserID())
	if err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}
	api.SendData(w, convertUserResponse(user), http.StatusOK)
}
