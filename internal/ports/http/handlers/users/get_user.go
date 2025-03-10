package users

import (
	"github.com/Rasikrr/learning_platform_core/api"
	"net/http"
)

// @Summary Get user
// @Description Get user by id
// @Tags users
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} userResponse "Success"
// @Router /api/v1/users/{id} [get]
func (c *Controller) getUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	email := r.PathValue("email")

	user, err := c.usersService.GetByEmail(ctx, email)
	if err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}
	api.SendData(w, convertUserResponse(user), http.StatusOK)
}
