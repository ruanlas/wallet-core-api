package routes

import "github.com/ruanlas/wallet-core-api/internal/infra/handlers"

var createInvoice, updateInvoice, deleteInvoice,
	getAllInvoice, getInvoiceById handlers.Handler

func RegisterInvoiceHandlers(create, update, delete,
	getAll, getById handlers.Handler) {
	createInvoice = create
	updateInvoice = update
	deleteInvoice = delete
	getAllInvoice = getAll
	getInvoiceById = getById
}
