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

	DataHandler dl.IBaseData
}

func (base *BaseLogicHandler) Paginate(request *models.IRequest) (*models.PaginateResult, error)  {
	if base.DataHandler != nil {
		return base.DataHandler.Paginate(request), nil
	}
	return nil, errors.New("data handler is not initialized")
}
