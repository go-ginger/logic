package logic

import (
	"errors"
	"github.com/kulichak/dl"
	"github.com/kulichak/models"
)

type IBaseLogicHandler interface {
	Paginate(request *models.IRequest) (*models.PaginateResult, error)
}

type BaseLogicHandler struct {
	IBaseLogicHandler

	DbHandler dl.IBaseDbHandler
}

func (base *BaseLogicHandler) Paginate(request *models.IRequest) (*models.PaginateResult, error) {
	if base.DbHandler != nil {
		return base.DbHandler.Paginate(request)
	}
	return nil, errors.New("data handler is not initialized")
}
