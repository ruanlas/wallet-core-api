package v1

import "github.com/ruanlas/wallet-core-api/internal/v1/gainprojection"

type Api interface {
	GetGainProjectionHandler() gainprojection.Handler
}

func NewApi(gainProjectionHandler gainprojection.Handler) Api {
	return &api{gainProjectionHandler: gainProjectionHandler}
}

type api struct {
	gainProjectionHandler gainprojection.Handler
}

func (a *api) GetGainProjectionHandler() gainprojection.Handler {
	return a.gainProjectionHandler
}
