package route

import (
	"time"

	"logo-lab-demo/bootstrap"
	"logo-lab-demo/middleware"
	"logo-lab-demo/mongo"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.Auth0Domain, env.Auth0Audience))
	// All Private APIs
	NewMessageRouter(timeout, db, protectedRouter)
}
