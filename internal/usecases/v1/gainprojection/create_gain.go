package gainprojection

import (
	"time"

	"github.com/ruanlas/wallet-core-api/internal/idpauth"
	"github.com/ruanlas/wallet-core-api/internal/infra/repository/gainprojection"
	uuid "github.com/satori/go.uuid"
)

type CreateGainUseCase interface {
	Execute(createGainCtx CreateGainContext) (*GainStat, error)
}

func NewCreateGainUseCase(repository gainprojection.Repository, generateUUID func() uuid.UUID) CreateGainUseCase {
	return &createGainUseCase{
		repository:   repository,
		generateUUID: generateUUID,
	}
}

type createGainUseCase struct {
	repository   gainprojection.Repository
	generateUUID func() uuid.UUID
}

func (uc *createGainUseCase) Execute(createGainCtx CreateGainContext) (*GainStat, error) {
	request := createGainCtx.Request
	user := idpauth.GetUser(createGainCtx.UserToken)
	gainProjection, err := uc.repository.GetById(createGainCtx.Ctx, createGainCtx.Id, user.Id)
	if err != nil {
		return nil, err
	}
	if gainProjection == nil {
		return &GainStat{ProjectionIsFound: false, ProjectionIsAlreadyDone: false}, nil
	}
	if gainProjection.IsAlreadyDone == true {
		return &GainStat{ProjectionIsFound: true, ProjectionIsAlreadyDone: true}, nil
	}
	gainBuilder := gainprojection.NewGainBuilder().
		AddId(uc.generateUUID().String()).
		AddCategory(gainProjection.Category).
		AddCreatedAt(time.Now()).
		AddDescription(gainProjection.Description).
		AddGainProjectionId(gainProjection.Id).
		AddIsPassive(gainProjection.IsPassive).
		AddUserId(gainProjection.UserId).
		AddValue(gainProjection.Value).
		AddPayIn(gainProjection.PayIn)
	if request.Value != 0 {
		gainBuilder.AddValue(request.Value)
	}
	if !request.PayIn.IsZero() {
		gainBuilder.AddPayIn(request.PayIn)
	}
	gain, err := uc.repository.SaveGain(createGainCtx.Ctx, *gainBuilder.Build())
	if err != nil {
		return nil, err
	}
	gainProjection.IsAlreadyDone = true
	_, err = uc.repository.Edit(createGainCtx.Ctx, *gainProjection)
	if err != nil {
		return nil, err
	}

	gainResponse := NewGainResponseBuilder().
		AddId(gain.Id).
		AddGainProjectionId(gain.GainProjectionId).
		AddPayIn(gain.PayIn).
		AddDescription(gain.Description).
		AddValue(gain.Value).
		AddIsPassive(gain.IsPassive).
		AddCategory(CategoryResponse{Id: gain.Category.Id, Category: gain.Category.Category}).
		Build()
	return &GainStat{ProjectionIsFound: true, ProjectionIsAlreadyDone: false, Gain: gainResponse}, nil
}
