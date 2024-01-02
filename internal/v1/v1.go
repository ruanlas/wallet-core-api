package v1

import (
	"github.com/ruanlas/wallet-core-api/internal/v1/gain"
	"github.com/ruanlas/wallet-core-api/internal/v1/gainprojection"
)

type Api interface {
	GetGainProjectionHandler() gainprojection.Handler
	GetGainHandler() gain.Handler
}

func NewApi(gainProjectionHandler gainprojection.Handler, gainHandler gain.Handler) Api {
	return &api{
		gainProjectionHandler: gainProjectionHandler,

		gainHandler: gainHandler}
}

type api struct {
	gainProjectionHandler gainprojection.Handler
	gainHandler           gain.Handler
}

func (a *api) GetGainProjectionHandler() gainprojection.Handler {
	return a.gainProjectionHandler
}

func (a *api) GetGainHandler() gain.Handler {
	return a.gainHandler
}
