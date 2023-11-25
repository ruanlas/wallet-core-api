package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/ruanlas/wallet-core-api/internal/routes"
	v1 "github.com/ruanlas/wallet-core-api/internal/v1"
	"github.com/ruanlas/wallet-core-api/internal/v1/gainprojection"
	gainprojectionrepository "github.com/ruanlas/wallet-core-api/internal/v1/gainprojection/repository"
	gainprojectionservice "github.com/ruanlas/wallet-core-api/internal/v1/gainprojection/service"
	uuid "github.com/satori/go.uuid"
	"go.elastic.co/apm/module/apmsql"
	_ "go.elastic.co/apm/module/apmsql/mysql"
)

var (
	db *sql.DB
)

func init() {
	env := os.Getenv("ENV")
	if env == "" {
		godotenv.Load()
	}
	log.Println("ELASTIC_APM_SERVICE_NAME:", os.Getenv("ELASTIC_APM_SERVICE_NAME"))

	var err error
	// db, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/wallet_core?charset=utf8&parseTime=True&loc=Local")
	db, err = apmsql.Open("mysql", "root:123456@tcp(localhost:3306)/wallet_core?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		panic(err)
	}
	go startPrometheus()
}

// @title Wallet Core
// @version 0.1.0
// @description API que dispões de recursos para gerenciar as finanças pessoais
// @host localhost:8080
// @BasePath /api/
func main() {
	fmt.Println("Project Started!")

	gainProjectionRepository := gainprojectionrepository.New(db)
	gainProjectionStorageProcess := gainprojectionservice.NewStorageProcess(gainProjectionRepository, uuid.NewV4)
	gainProjectionHandler := gainprojection.NewHandler(gainProjectionStorageProcess)

	apiV1 := v1.NewApi(gainProjectionHandler)
	router := routes.NewRouter(apiV1)
	router.SetupRoutes()

	// r := gin.Default()

	// // docs.SwaggerInfo.BasePath = "/api"
	// r.GET("/", func(c *gin.Context) {

	// 	tx := apm.TransactionFromContext(c.Request.Context())
	// 	span := tx.StartSpan("Default span", "handler_main", nil)
	// 	span.End()
	// 	c.JSON(http.StatusOK, gin.H{"message": "Welcome to API!"})
	// })

	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// r.Run(":8080")
}

func startPrometheus() {
	log.Println("Prometheus metrics on /metrics port 2112")
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		res := "{ \"status\": \"ok\" }"
		w.Write([]byte(res))
	})
	http.ListenAndServe(":2112", nil)
}
