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

func NewUserRouter(timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pc := &controller.UserController{
		UserService: service.NewUserService(ur, timeout),
	}
	group.GET("/user", pc.Fetch)
}
