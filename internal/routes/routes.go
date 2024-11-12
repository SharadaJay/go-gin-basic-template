package routes

import (
	"basic-gin-app/internal/controllers"
	"basic-gin-app/internal/middleware"
	"basic-gin-app/internal/utils"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Routes struct {
}

func (c Routes) SetupRouter() {
	router := gin.Default()
	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("my_session", store))
	router.Use(middleware.CORSMiddleware())

	controller := controllers.NewController()

	healthCheckAPI := router.Group("/base-app")
	{
		healthCheckAPI.GET("/health-check", controller.HealthStatus)
	}

	resourceAPI := router.Group("/resources")
	{
		resourceAPI.POST("/v1", controller.CreateResource)
		resourceAPI.GET("/v1/:id", controller.GetResourceByID)
		resourceAPI.GET("/v1", controller.GetResources)
		resourceAPI.PUT("/v1/:id", controller.UpdateResource)
		resourceAPI.DELETE("/v1/:id", controller.DeleteResource)
	}

	err := router.Run(utils.Cfg["application"].Port)
	if err != nil {
		log.Error("Error while running router:", err)
	}
}
