// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Swagger docs.
	"github.com/ferralucho/digital-course-api/config"
	_ "github.com/ferralucho/digital-course-api/docs"
	"github.com/ferralucho/digital-course-api/internal/usecase"
	"github.com/ferralucho/digital-course-api/pkg/logger"
)

// NewRouter -.
// Swagger spec:
// @title       Digital Course API
// @description Api for digital courses
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(handler *gin.Engine, l logger.Interface, t usecase.CoursePlanning) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// K8s probe
	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Prometheus metrics
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	firebaseAuth := config.SetupFirebase()

	handler.Use(func(c *gin.Context) {
		c.Set("firebaseAuth", firebaseAuth)
	})

	// using the auth middle ware to validate api requests
	//handler.Use(middleware.AuthMiddleware)

	// Routers
	h := handler.Group("/v1")
	{
		newCoursePlanningRoutes(h, t, l)
	}
}
