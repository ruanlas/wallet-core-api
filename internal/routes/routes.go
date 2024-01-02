package routes

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ruanlas/wallet-core-api/docs"
	"github.com/ruanlas/wallet-core-api/internal/idpauth"
	v1 "github.com/ruanlas/wallet-core-api/internal/v1"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.elastic.co/apm/module/apmgin"
)

type Router struct {
	apiV1 v1.Api
}

func NewRouter(apiV1 v1.Api) *Router {
	return &Router{apiV1: apiV1}
}

func (r *Router) SetupRoutes() {
	servicePort := os.Getenv("SERVICE_PORT")
	serviceHost := os.Getenv("SERVICE_HOST")
	router := gin.Default()
	router.Use(apmgin.Middleware(router), idpauth.AuthenticationMiddleware, idpauth.AuthorizationMiddleware)

	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", serviceHost, servicePort)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	v1router := router.Group("/v1")
	v1router.POST("/gain-projection", r.apiV1.GetGainProjectionHandler().Create)
	v1router.GET("/gain-projection", r.apiV1.GetGainProjectionHandler().GetAll)
	v1router.GET("/gain-projection/:id", r.apiV1.GetGainProjectionHandler().GetById)
	v1router.PUT("/gain-projection/:id", r.apiV1.GetGainProjectionHandler().Update)
	v1router.DELETE("/gain-projection/:id", r.apiV1.GetGainProjectionHandler().Delete)
	v1router.POST("/gain-projection/:id/create-gain", r.apiV1.GetGainProjectionHandler().CreateGain)

	v1router.POST("/gain", r.apiV1.GetGainHandler().Create)
	v1router.GET("/gain", r.apiV1.GetGainHandler().GetAll)
	v1router.GET("/gain/:id", r.apiV1.GetGainHandler().GetById)
	v1router.PUT("/gain/:id", r.apiV1.GetGainHandler().Update)
	v1router.DELETE("/gain/:id", r.apiV1.GetGainHandler().Delete)

	serviceAddr := fmt.Sprintf(":%s", servicePort)
	router.Run(serviceAddr)
}
