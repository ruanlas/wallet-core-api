package v1

import (
	"github.com/ruanlas/wallet-core-api/internal/v1/gain"
	"github.com/ruanlas/wallet-core-api/internal/v1/gainprojection"
	"github.com/ruanlas/wallet-core-api/internal/v1/invoice"
	"github.com/ruanlas/wallet-core-api/internal/v1/invoiceprojection"
)

type Api interface {
	GetGainProjectionHandler() gainprojection.Handler
	GetGainHandler() gain.Handler
	GetInvoiceProjectionHandler() invoiceprojection.Handler
	GetInvoiceHandler() invoice.Handler
}

func NewApi(gainProjectionHandler gainprojection.Handler, gainHandler gain.Handler, invoiceProjectionHandler invoiceprojection.Handler, invoiceHandler invoice.Handler) Api {
	return &api{
		gainProjectionHandler:    gainProjectionHandler,
		gainHandler:              gainHandler,
		invoiceProjectionHandler: invoiceProjectionHandler,
		invoiceHandler:           invoiceHandler}
}

type api struct {
	gainProjectionHandler    gainprojection.Handler
	gainHandler              gain.Handler
	invoiceProjectionHandler invoiceprojection.Handler
	invoiceHandler           invoice.Handler
}

func (a *api) GetGainProjectionHandler() gainprojection.Handler {
	return a.gainProjectionHandler
}

func (a *api) GetGainHandler() gain.Handler {
	return a.gainHandler
}

func (a *api) GetInvoiceProjectionHandler() invoiceprojection.Handler {
	return a.invoiceProjectionHandler
}

func (a *api) GetInvoiceHandler() invoice.Handler {
	return a.invoiceHandler
}
