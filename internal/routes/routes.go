package routes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ruanlas/wallet-core-api/internal/v1"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	apiV1 v1.Api
}

func NewRouter(apiV1 v1.Api) *Router {
	return &Router{apiV1: apiV1}
}

func (r *Router) SetupRoutes() {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	v1router := router.Group("/v1")
	v1router.POST("/gain-projection", r.apiV1.GetGainProjectionHandler().Create)

	router.Run(":8080")
}
