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
	"github.com/ruanlas/wallet-core-api/internal/v1/gain"
	gainservice "github.com/ruanlas/wallet-core-api/internal/v1/gain/gservice"
	gainrepository "github.com/ruanlas/wallet-core-api/internal/v1/gain/repository"
	"github.com/ruanlas/wallet-core-api/internal/v1/gainprojection"
	gainprojectionservice "github.com/ruanlas/wallet-core-api/internal/v1/gainprojection/gpservice"
	gainprojectionrepository "github.com/ruanlas/wallet-core-api/internal/v1/gainprojection/repository"
	"github.com/ruanlas/wallet-core-api/internal/v1/invoiceprojection"
	invoiceprojectionservice "github.com/ruanlas/wallet-core-api/internal/v1/invoiceprojection/ipservice"
	invoiceprojectionrepository "github.com/ruanlas/wallet-core-api/internal/v1/invoiceprojection/repository"
	uuid "github.com/satori/go.uuid"
	"go.elastic.co/apm/module/apmsql"
	_ "go.elastic.co/apm/module/apmsql/mysql"
)

var (
	db *sql.DB
)

var requiredEnvs = []string{
	"SERVICE_HOST", "SERVICE_PORT", "PROMETHEUS_PORT",
	"DATABASE_HOST", "DATABASE_PORT", "DATABASE_NAME", "DATABASE_USERNAME", "DATABASE_PASSWORD",
	"IDP_HOST", "IDP_PORT", "IDP_MAIN_REALM", "IDP_USER_ADMIN", "IDP_PASSWORD_ADMIN", "IDP_REALM", "IDP_CLIENT_IDENTIFIER", "IDP_CLIENT_SECRET",
}

func checkRequiredEnvs() {
	for _, envName := range requiredEnvs {
		if os.Getenv(envName) == "" {
			panic(fmt.Sprintf("You must to define %s env", envName))
		}
	}
}

func init() {
	env := os.Getenv("ENV")
	if env == "" {
		godotenv.Load()
	}
	checkRequiredEnvs()

	dbUsername := os.Getenv("DATABASE_USERNAME")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("DATABASE_NAME")
	var err error
	// db, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/wallet_core?charset=utf8&parseTime=True&loc=Local")
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUsername, dbPassword, dbHost, dbPort, dbName)
	db, err = apmsql.Open("mysql", connectionString)
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
	gainProjectionReadingProcess := gainprojectionservice.NewReadingProcess(gainProjectionRepository)
	gainProjectionHandler := gainprojection.NewHandler(gainProjectionStorageProcess, gainProjectionReadingProcess)

	gainRepository := gainrepository.New(db)
	gainStorageProcess := gainservice.NewStorageProcess(gainRepository, uuid.NewV4)
	gainReadingProcess := gainservice.NewReadingProcess(gainRepository)
	gainHandler := gain.NewHandler(gainStorageProcess, gainReadingProcess)

	invoiceProjectionRepository := invoiceprojectionrepository.New(db)
	invoiceProjectionStorageProcess := invoiceprojectionservice.NewStorageProcess(invoiceProjectionRepository, uuid.NewV4)
	invoiceProjectionReadingProcess := invoiceprojectionservice.NewReadingProcess(invoiceProjectionRepository)
	invoiceProjectionHandler := invoiceprojection.NewHandler(invoiceProjectionStorageProcess, invoiceProjectionReadingProcess)

	apiV1 := v1.NewApi(gainProjectionHandler, gainHandler, invoiceProjectionHandler)
	router := routes.NewRouter(apiV1)
	router.SetupRoutes()
}

func startPrometheus() {
	prometheusPort := os.Getenv("PROMETHEUS_PORT")
	log.Println("Prometheus metrics on /metrics port", prometheusPort)
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		res := "{ \"status\": \"ok\" }"
		w.Write([]byte(res))
	})
	prometheusAddr := fmt.Sprintf(":%s", prometheusPort)
	http.ListenAndServe(prometheusAddr, nil)
}
