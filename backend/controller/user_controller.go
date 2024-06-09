package controller

import (
	"net/http"

	"logo-lab-demo/domain"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService domain.UserService
}

func (pc *UserController) Fetch(c *gin.Context) {
	email := c.GetString("x-user-id")

	user, err := pc.UserService.GetProfileByEmail(c, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
