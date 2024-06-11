package route

import (
	"time"

	"logo-lab-demo/controller"
	"logo-lab-demo/domain"
	"logo-lab-demo/mongo"
	"logo-lab-demo/repository"
	"logo-lab-demo/service"

	"github.com/gin-gonic/gin"
)

func NewMessageRouter(timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewMessageRepository(db, domain.CollectionMessage)
	pc := &controller.MessageController{
		MessageService: service.NewMessageService(ur, timeout),
	}
	group.GET("/messages", pc.Fetch)
	group.POST("/messages", pc.Create)
}
