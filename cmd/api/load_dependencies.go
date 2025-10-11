package main

import (
	"fmt"

	"github.com/ruanlas/wallet-core-api/internal/configs"
	gainhandler "github.com/ruanlas/wallet-core-api/internal/infra/handlers/v1/gain"
	gainprojectionhandler "github.com/ruanlas/wallet-core-api/internal/infra/handlers/v1/gainprojection"
	invoicehandler "github.com/ruanlas/wallet-core-api/internal/infra/handlers/v1/invoice"
	invoiceprojectionhandler "github.com/ruanlas/wallet-core-api/internal/infra/handlers/v1/invoiceprojection"
	gainrepo "github.com/ruanlas/wallet-core-api/internal/infra/repository/gain"
	gainprojectionrepo "github.com/ruanlas/wallet-core-api/internal/infra/repository/gainprojection"
	invoicerepo "github.com/ruanlas/wallet-core-api/internal/infra/repository/invoice"
	invoiceprojectionrepo "github.com/ruanlas/wallet-core-api/internal/infra/repository/invoiceprojection"
	"github.com/ruanlas/wallet-core-api/internal/usecases/v1/gain"
	"github.com/ruanlas/wallet-core-api/internal/usecases/v1/gainprojection"
	"github.com/ruanlas/wallet-core-api/internal/usecases/v1/invoice"
	"github.com/ruanlas/wallet-core-api/internal/usecases/v1/invoiceprojection"
	uuid "github.com/satori/go.uuid"
	"go.elastic.co/apm/module/apmsql"
	_ "go.elastic.co/apm/module/apmsql/mysql"
)

func initializeDatabaseConnection() {
	var err error
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		configs.App.Database.Username, configs.App.Database.Password,
		configs.App.Database.Host, configs.App.Database.Port, configs.App.Database.Name)
	db, err = apmsql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		panic(err)
	}
}

func initializeRepositories() {
	gainRepository = gainrepo.New(db)
	invoiceRepository = invoicerepo.New(db)
	gainProjectionRepository = gainprojectionrepo.New(db)
	invoiceProjectionRepository = invoiceprojectionrepo.New(db)
}

func initializeUseCases() {
	createGainUseCase = gain.NewCreateUseCase(gainRepository, uuid.NewV4)
	updateGainUseCase = gain.NewUpdateUseCase(gainRepository)
	deleteGainUseCase = gain.NewDeleteUseCase(gainRepository)
	getGainByIdUseCase = gain.NewGetByIdUseCase(gainRepository)
	getAllGainsUseCase = gain.NewGetAllUseCase(gainRepository)

	createGainProjectionUseCase = gainprojection.NewCreateUseCase(gainProjectionRepository, uuid.NewV4)
	updateGainProjectionUseCase = gainprojection.NewUpdateUseCase(gainProjectionRepository)
	deleteGainProjectionUseCase = gainprojection.NewDeleteUseCase(gainProjectionRepository)
	getGainProjectionByIdUseCase = gainprojection.NewGetByIdUseCase(gainProjectionRepository)
	getAllGainsProjectionUseCase = gainprojection.NewGetAllUseCase(gainProjectionRepository)
	createGainForProjectionUseCase = gainprojection.NewCreateGainUseCase(gainProjectionRepository, uuid.NewV4)

	createInvoiceUseCase = invoice.NewCreateUseCase(invoiceRepository, uuid.NewV4)
	updateInvoiceUseCase = invoice.NewUpdateUseCase(invoiceRepository)
	deleteInvoiceUseCase = invoice.NewDeleteUseCase(invoiceRepository)
	getInvoiceByIdUseCase = invoice.NewGetByIdUseCase(invoiceRepository)
	getAllInvoicesUseCase = invoice.NewGetAllUseCase(invoiceRepository)

	createInvoiceProjectionUseCase = invoiceprojection.NewCreateUseCase(invoiceProjectionRepository, uuid.NewV4)
	updateInvoiceProjectionUseCase = invoiceprojection.NewUpdateUseCase(invoiceProjectionRepository)
	deleteInvoiceProjectionUseCase = invoiceprojection.NewDeleteUseCase(invoiceProjectionRepository)
	getInvoiceProjectionByIdUseCase = invoiceprojection.NewGetByIdUseCase(invoiceProjectionRepository)
	getAllInvoicesProjectionUseCase = invoiceprojection.NewGetAllUseCase(invoiceProjectionRepository)
	createInvoiceForProjectionUseCase = invoiceprojection.NewCreateInvoiceUseCase(invoiceProjectionRepository, uuid.NewV4)
}

func initializeHandlers() {
	createGainHandler = gainhandler.NewCreateHandler(createGainUseCase)
	updateGainHandler = gainhandler.NewUpdateHandler(updateGainUseCase)
	deleteGainHandler = gainhandler.NewDeleteHandler(deleteGainUseCase)
	getAllGainsHandler = gainhandler.NewGetAllHandler(getAllGainsUseCase)
	getGainByIdHandler = gainhandler.NewGetByIdHandler(getGainByIdUseCase)

	createGainProjectionHandler = gainprojectionhandler.NewCreateHandler(createGainProjectionUseCase)
	updateGainProjectionHandler = gainprojectionhandler.NewUpdateHandler(updateGainProjectionUseCase)
	deleteGainProjectionHandler = gainprojectionhandler.NewDeleteHandler(deleteGainProjectionUseCase)
	getAllGainsProjectionHandler = gainprojectionhandler.NewGetAllHandler(getAllGainsProjectionUseCase)
	getGainProjectionByIdHandler = gainprojectionhandler.NewGetByIdHandler(getGainProjectionByIdUseCase)
	createGainForProjectionHandler = gainprojectionhandler.NewCreateGainHandler(createGainForProjectionUseCase)

	createInvoiceHandler = invoicehandler.NewCreateHandler(createInvoiceUseCase)
	updateInvoiceHandler = invoicehandler.NewUpdateHandler(updateInvoiceUseCase)
	deleteInvoiceHandler = invoicehandler.NewDeleteHandler(deleteInvoiceUseCase)
	getAllInvoicesHandler = invoicehandler.NewGetAllHandler(getAllInvoicesUseCase)
	getInvoiceByIdHandler = invoicehandler.NewGetByIdHandler(getInvoiceByIdUseCase)

	createInvoiceProjectionHandler = invoiceprojectionhandler.NewCreateHandler(createInvoiceProjectionUseCase)
	updateInvoiceProjectionHandler = invoiceprojectionhandler.NewUpdateHandler(updateInvoiceProjectionUseCase)
	deleteInvoiceProjectionHandler = invoiceprojectionhandler.NewDeleteHandler(deleteInvoiceProjectionUseCase)
	getAllInvoicesProjectionHandler = invoiceprojectionhandler.NewGetAllHandler(getAllInvoicesProjectionUseCase)
	getInvoiceProjectionByIdHandler = invoiceprojectionhandler.NewGetByIdHandler(getInvoiceProjectionByIdUseCase)
	createInvoiceForProjectionHandler = invoiceprojectionhandler.NewCreateInvoiceHandler(createInvoiceForProjectionUseCase)
}

func loadDependencies() {
	// Initialize database connection
	initializeDatabaseConnection()

	// Initialize repositories
	initializeRepositories()

	// Initialize use cases
	initializeUseCases()

	// Initialize handlers
	initializeHandlers()
}
