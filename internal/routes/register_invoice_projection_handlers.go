package routes

import "github.com/ruanlas/wallet-core-api/internal/infra/handlers"

var createInvoiceProjection, updateInvoiceProjection, deleteInvoiceProjection,
	getAllInvoiceProjection, getInvoiceProjectionById, createInvoiceForProjection handlers.Handler

func RegisterInvoiceProjectionHandlers(create, update, delete,
	getAll, getById, createInvoice handlers.Handler) {
	createInvoiceProjection = create
	updateInvoiceProjection = update
	deleteInvoiceProjection = delete
	getAllInvoiceProjection = getAll
	getInvoiceProjectionById = getById
	createInvoiceForProjection = createInvoice
}
