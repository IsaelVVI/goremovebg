package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()

	initializeRoutes(router)

	router.Use()

	router.Run(":3000") // listen and serve on 0.0.0.0:3000
}
