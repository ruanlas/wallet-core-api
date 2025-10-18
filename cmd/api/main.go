package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/ruanlas/wallet-core-api/internal/configs"
	"github.com/ruanlas/wallet-core-api/internal/routes"
)

func init() {
	godotenv.Load()
	configs.Init()

	loadDependencies()

	routes.RegisterGainHandlers(
		createGainHandler,
		updateGainHandler,
		deleteGainHandler,
		getAllGainsHandler,
		getGainByIdHandler,
	)

	routes.RegisterInvoiceHandlers(
		createInvoiceHandler,
		updateInvoiceHandler,
		deleteInvoiceHandler,
		getAllInvoicesHandler,
		getInvoiceByIdHandler,
	)

	routes.RegisterGainProjectionHandlers(
		createGainProjectionHandler,
		updateGainProjectionHandler,
		deleteGainProjectionHandler,
		getAllGainsProjectionHandler,
		getGainProjectionByIdHandler,
		createGainForProjectionHandler,
	)

	routes.RegisterInvoiceProjectionHandlers(
		createInvoiceProjectionHandler,
		updateInvoiceProjectionHandler,
		deleteInvoiceProjectionHandler,
		getAllInvoicesProjectionHandler,
		getInvoiceProjectionByIdHandler,
		createInvoiceForProjectionHandler,
	)
}

// TODO: isolar em outro pacote de monitoramento para expor as métricas e o healthcheck
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

func main() {
	go startPrometheus()
	routes.Start()
}
