package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"logo-lab-demo/domain"

	"github.com/gin-gonic/gin"
)

type MessageController struct {
	MessageService domain.MessageService
}

func (pc *MessageController) Fetch(c *gin.Context) {
	queryPage, exists := c.GetQuery("page")
	page := 0
	if exists {
		v, err := strconv.Atoi(queryPage)
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
			return
		}
		page = v
	}

	res, err := pc.MessageService.GetMessages(c, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (pc *MessageController) Create(c *gin.Context) {
	message := &domain.CreateMessage{}
	json.NewDecoder(c.Request.Body).Decode(message)

	err := pc.MessageService.InsertMessage(c, *message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}
