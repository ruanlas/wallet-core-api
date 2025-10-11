package main

import (
	"database/sql"

	"github.com/ruanlas/wallet-core-api/internal/usecases/v1/gain"
	"github.com/ruanlas/wallet-core-api/internal/usecases/v1/gainprojection"
	"github.com/ruanlas/wallet-core-api/internal/usecases/v1/invoice"
	"github.com/ruanlas/wallet-core-api/internal/usecases/v1/invoiceprojection"

	"github.com/ruanlas/wallet-core-api/internal/infra/handlers"
	gainrepo "github.com/ruanlas/wallet-core-api/internal/infra/repository/gain"
	gainprojectionrepo "github.com/ruanlas/wallet-core-api/internal/infra/repository/gainprojection"
	invoicerepo "github.com/ruanlas/wallet-core-api/internal/infra/repository/invoice"
	invoiceprojectionrepo "github.com/ruanlas/wallet-core-api/internal/infra/repository/invoiceprojection"
)

var (
	db *sql.DB

	gainRepository              gainrepo.Repository
	invoiceRepository           invoicerepo.Repository
	gainProjectionRepository    gainprojectionrepo.Repository
	invoiceProjectionRepository invoiceprojectionrepo.Repository

	createGainUseCase  gain.CreateUseCase
	updateGainUseCase  gain.UpdateUseCase
	deleteGainUseCase  gain.DeleteUseCase
	getGainByIdUseCase gain.GetByIdUseCase
	getAllGainsUseCase gain.GetAllUseCase

	createGainProjectionUseCase    gainprojection.CreateUseCase
	updateGainProjectionUseCase    gainprojection.UpdateUseCase
	deleteGainProjectionUseCase    gainprojection.DeleteUseCase
	getGainProjectionByIdUseCase   gainprojection.GetByIdUseCase
	getAllGainsProjectionUseCase   gainprojection.GetAllUseCase
	createGainForProjectionUseCase gainprojection.CreateGainUseCase

	createInvoiceUseCase  invoice.CreateUseCase
	updateInvoiceUseCase  invoice.UpdateUseCase
	deleteInvoiceUseCase  invoice.DeleteUseCase
	getInvoiceByIdUseCase invoice.GetByIdUseCase
	getAllInvoicesUseCase invoice.GetAllUseCase

	createInvoiceProjectionUseCase    invoiceprojection.CreateUseCase
	updateInvoiceProjectionUseCase    invoiceprojection.UpdateUseCase
	deleteInvoiceProjectionUseCase    invoiceprojection.DeleteUseCase
	getInvoiceProjectionByIdUseCase   invoiceprojection.GetByIdUseCase
	getAllInvoicesProjectionUseCase   invoiceprojection.GetAllUseCase
	createInvoiceForProjectionUseCase invoiceprojection.CreateInvoiceUseCase

	createGainHandler  handlers.Handler
	deleteGainHandler  handlers.Handler
	updateGainHandler  handlers.Handler
	getAllGainsHandler handlers.Handler
	getGainByIdHandler handlers.Handler

	createGainProjectionHandler    handlers.Handler
	deleteGainProjectionHandler    handlers.Handler
	updateGainProjectionHandler    handlers.Handler
	getAllGainsProjectionHandler   handlers.Handler
	getGainProjectionByIdHandler   handlers.Handler
	createGainForProjectionHandler handlers.Handler

	createInvoiceHandler  handlers.Handler
	deleteInvoiceHandler  handlers.Handler
	updateInvoiceHandler  handlers.Handler
	getAllInvoicesHandler handlers.Handler
	getInvoiceByIdHandler handlers.Handler

	createInvoiceProjectionHandler    handlers.Handler
	deleteInvoiceProjectionHandler    handlers.Handler
	updateInvoiceProjectionHandler    handlers.Handler
	getAllInvoicesProjectionHandler   handlers.Handler
	getInvoiceProjectionByIdHandler   handlers.Handler
	createInvoiceForProjectionHandler handlers.Handler
)
