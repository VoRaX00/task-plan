package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"task-plan/internal/application/requestModels"
	"task-plan/pkg/tokenManager"
)

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept json
// @Produce json
// @Param input body domain.User true "account info"
// @Success 200 {integer} integer 1
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input requestModels.UserToAdd
	if err := c.ShouldBindJSON(&input); err != nil {
		logrus.Error(err)
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.IAuthService.Create(input)
	if err != nil {
		logrus.Error(err)
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	manager := tokenManager.NewManager(h.signingKey)
	emailToken, err := h.services.GenerateEmailConfirmationToken(id.String(), manager)
	if err != nil {
		logrus.Error(err)
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.IMessageEmailService.SendConfirmEmail(input.Email, emailToken)
	if err != nil {
		logrus.Error(err)
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}

// @Summary SignIn
// @Tags auth
// @Description login account
// @ID login-account
// @Accept json
// @Produce json
// @Param input body domain.User true "account info"
// @Success 200 {integer} integer 1
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {

}
