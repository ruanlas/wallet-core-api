package v1

import "github.com/ruanlas/wallet-core-api/internal/infra/handlers"

type API struct {
	Gain              *GainHandler
	GainProjection    *GainProjectionHandler
	Invoice           *InvoiceHandler
	InvoiceProjection *InvoiceProjectionHandler
}

type GainHandler struct {
	Create  handlers.Handler
	Update  handlers.Handler
	Delete  handlers.Handler
	GetById handlers.Handler
	GetAll  handlers.Handler
}

type GainProjectionHandler struct {
	Create     handlers.Handler
	CreateGain handlers.Handler
	Update     handlers.Handler
	Delete     handlers.Handler
	GetById    handlers.Handler
	GetAll     handlers.Handler
}

type InvoiceHandler struct {
	Create  handlers.Handler
	Update  handlers.Handler
	Delete  handlers.Handler
	GetById handlers.Handler
	GetAll  handlers.Handler
}

type InvoiceProjectionHandler struct {
	Create        handlers.Handler
	CreateInvoice handlers.Handler
	Update        handlers.Handler
	Delete        handlers.Handler
	GetById       handlers.Handler
	GetAll        handlers.Handler
}
