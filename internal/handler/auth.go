package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"task-plan/internal/application"
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
	var input application.UserRegister
	if err := c.ShouldBindJSON(&input); err != nil {
		logrus.Error(err)
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
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
