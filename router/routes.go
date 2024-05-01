package router

import (
	controllers "github.com/IsaelVVI/goremovebg.git/controllers/image"
	"github.com/IsaelVVI/goremovebg.git/middleware"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	basePath := "/api/v1"

	v1 := router.Group(basePath, middleware.CORSMiddleware())
	{
		v1.POST("/removebg", controllers.HandleRemovebg)
	}
}
