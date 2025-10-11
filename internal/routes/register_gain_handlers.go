package routes

import "github.com/ruanlas/wallet-core-api/internal/infra/handlers"

var createGain, updateGain, deleteGain,
	getAllGain, getGainById handlers.Handler

func RegisterGainHandlers(create, update, delete,
	getAll, getById handlers.Handler) {
	createGain = create
	updateGain = update
	deleteGain = delete
	getAllGain = getAll
	getGainById = getById
}
