package routes

import "github.com/ruanlas/wallet-core-api/internal/infra/handlers"

var createGainProjection, updateGainProjection, deleteGainProjection,
	getAllGainProjection, getGainProjectionById, createGainForProjection handlers.Handler

func RegisterGainProjectionHandlers(create, update, delete,
	getAll, getById, createGain handlers.Handler) {
	createGainProjection = createGain
	updateGainProjection = update
	deleteGainProjection = delete
	getAllGainProjection = getAll
	getGainProjectionById = getById
	createGainForProjection = createGain
}
