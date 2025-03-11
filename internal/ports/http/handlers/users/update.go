package users

import (
	"github.com/Rasikrr/learning_platform/api"
	"github.com/Rasikrr/learning_platform_core/http/session"
	"log"
	"net/http"
)

// @Summary Update user
// @Description Update user
// @Tags users
// @Produce json
// @Security     BearerAuth
// @param Authorization header string true "Authorization token"
// @Param request body updateUserRequest true "request"
// @Success 200 {object} api.EmptySuccessResponse "Success"
// @Router /api/v1/users/update [put]
func (c *Controller) updateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ses, err := session.GetFromCtx(ctx)
	if err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}
	log.Println("ID ", ses.UserID())
	var req updateUserRequest
	if err := api.GetData(r, &req); err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	if err := c.usersClient.Update(ctx, req.ToEntity(ses)); err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}
	api.SendData(w, api.NewEmptySuccessResponse(), http.StatusOK)
}
