package routes

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ruanlas/wallet-core-api/docs"
	"github.com/ruanlas/wallet-core-api/internal/idpauth"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.elastic.co/apm/module/apmgin"
)

func Start() {
	servicePort := os.Getenv("SERVICE_PORT")
	serviceHost := os.Getenv("SERVICE_HOST")
	router := gin.Default()
	router.Use(apmgin.Middleware(router), idpauth.AuthenticationMiddleware, idpauth.AuthorizationMiddleware)

	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", serviceHost, servicePort)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	v1router := router.Group("/v1")
	v1router.POST("/gain-projection", createGainProjection.Handle)
	v1router.GET("/gain-projection", getAllGainProjection.Handle)
	v1router.GET("/gain-projection/:id", getGainProjectionById.Handle)
	v1router.PUT("/gain-projection/:id", updateGainProjection.Handle)
	v1router.DELETE("/gain-projection/:id", deleteGainProjection.Handle)
	v1router.POST("/gain-projection/:id/create-gain", createGainForProjection.Handle)

	v1router.POST("/gain", createGain.Handle)
	v1router.GET("/gain", getAllGain.Handle)
	v1router.GET("/gain/:id", getGainById.Handle)
	v1router.PUT("/gain/:id", updateGain.Handle)
	v1router.DELETE("/gain/:id", deleteGain.Handle)

	v1router.POST("/invoice-projection", createInvoiceProjection.Handle)
	v1router.GET("/invoice-projection", getAllInvoiceProjection.Handle)
	v1router.GET("/invoice-projection/:id", getInvoiceProjectionById.Handle)
	v1router.PUT("/invoice-projection/:id", updateInvoiceProjection.Handle)
	v1router.DELETE("/invoice-projection/:id", deleteInvoiceProjection.Handle)
	v1router.POST("/invoice-projection/:id/create-invoice", createInvoiceForProjection.Handle)

	v1router.POST("/invoice", createInvoice.Handle)
	v1router.GET("/invoice", getAllInvoice.Handle)
	v1router.GET("/invoice/:id", getInvoiceById.Handle)
	v1router.PUT("/invoice/:id", updateInvoice.Handle)
	v1router.DELETE("/invoice/:id", deleteInvoice.Handle)

	serviceAddr := fmt.Sprintf(":%s", servicePort)
	router.Run(serviceAddr)
}
