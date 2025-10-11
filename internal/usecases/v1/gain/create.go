package gain

import (
	"time"

	"github.com/ruanlas/wallet-core-api/internal/idpauth"
	gainrepo "github.com/ruanlas/wallet-core-api/internal/infra/repository/gain"
	uuid "github.com/satori/go.uuid"
)

type CreateUseCase interface {
	Execute(createCtx CreateContext) (*GainResponse, error)
}

type createUseCase struct {
	repository   gainrepo.Repository
	generateUUID func() uuid.UUID
}

func NewCreateUseCase(repository gainrepo.Repository, generateUUID func() uuid.UUID) CreateUseCase {
	return &createUseCase{
		repository:   repository,
		generateUUID: generateUUID,
	}
}

func (uc *createUseCase) Execute(createCtx CreateContext) (*GainResponse, error) {
	request := createCtx.Request
	user := idpauth.GetUser(createCtx.UserToken)
	createdAt := time.Now()
	gainBuilder := gainrepo.NewGainBuilder().
		AddId(uc.generateUUID().String()).
		AddCreatedAt(createdAt).
		AddPayIn(request.PayIn).
		AddIsPassive(request.IsPassive).
		AddCategory(gainrepo.GainCategory{Id: request.CategoryId}).
		AddDescription(request.Description).
		AddValue(request.Value).
		AddUserId(user.Id)
	if request.PayIn.IsZero() {
		gainBuilder.AddPayIn(createdAt)
	}

	gain := gainBuilder.Build()
	gainSaved, err := uc.repository.Save(createCtx.Ctx, *gain)
	if err != nil {
		return nil, err
	}

	gainSaved, err = uc.repository.GetById(createCtx.Ctx, gainSaved.Id, user.Id)
	if err != nil {
		return nil, err
	}

	return NewGainResponseBuilder().
		AddId(gain.Id).
		AddPayIn(gain.PayIn).
		AddDescription(gain.Description).
		AddValue(gain.Value).
		AddIsPassive(gain.IsPassive).
		AddCategory(CategoryResponse{Id: gainSaved.Category.Id, Category: gainSaved.Category.Category}).
		Build(), nil
}
