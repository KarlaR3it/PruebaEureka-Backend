package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/KarlaR3it/PruebaEureka-Backend/internal/handler"
)

func SetupRoutes(areaHandler *handler.AreaHandler, personHandler *handler.PersonHandler) *gin.Engine {
	router := gin.Default()

	// Configurar CORS (permitir todos los orígenes - solo para desarrollo)
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
	}))

	// Grupo de rutas para áreas
	areaGroup := router.Group("/areas")
	{
		areaGroup.POST("", areaHandler.Create)
		areaGroup.GET("", areaHandler.GetAll)
		areaGroup.GET("/count", areaHandler.GetAreaCounts)
	}

	// Grupo de rutas para personas
	personGroup := router.Group("/persons")
	{
		personGroup.POST("", personHandler.Create)
		personGroup.GET("", personHandler.GetAll)
	}

	return router
}
