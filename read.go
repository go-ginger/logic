package logic

import (
	"errors"
	"github.com/kulichak/models"
)

func (base *BaseLogicHandler) DoPaginate(request models.IRequest) (*models.PaginateResult, error) {
	base.handleRequestFunction(base.LogicHandler.Models, request)
	base.handleRequestFunction(base.LogicHandler.BeforeQuery, request)
	result, err := base.LogicHandler.Paginate(request)
	if err != nil {
		return nil, err
	}
	base.handleRequestFunction(base.LogicHandler.AfterQuery, request)
	return result, err
}

func (base *BaseLogicHandler) DoGet(request models.IRequest) (interface{}, error) {
	base.handleRequestFunction(base.LogicHandler.Model, request)
	base.handleRequestFunction(base.LogicHandler.BeforeQuery, request)
	result, err := base.LogicHandler.Get(request)
	if err != nil {
		return nil, err
	}
	base.handleRequestFunction(base.LogicHandler.AfterQuery, request)
	return result, err
}

func (base *BaseLogicHandler) BeforeQuery(request models.IRequest) {
}

func (base *BaseLogicHandler) Paginate(request models.IRequest) (*models.PaginateResult, error) {
	if base.DataHandler != nil {
		base.LogicHandler.Models(request)
		base.handleRequestFunction(base.DataHandler.BeforeQuery, request)
		result, err := base.DataHandler.Paginate(request)
		if err != nil {
			return nil, err
		}
		base.handleRequestFunction(base.DataHandler.AfterQuery, request)
		return result, err
	}
	return nil, errors.New("data handler is not initialized")
}

func (base *BaseLogicHandler) Get(request models.IRequest) (interface{}, error) {
	if base.DataHandler != nil {
		base.LogicHandler.Model(request)
		base.handleRequestFunction(base.DataHandler.BeforeQuery, request)
		result, err := base.DataHandler.Get(request)
		if err != nil {
			return nil, err
		}
		base.handleRequestFunction(base.DataHandler.AfterQuery, request)
		return result, err
	}
	return nil, errors.New("data handler is not initialized")
}

func (base *BaseLogicHandler) AfterQuery(request models.IRequest) {
}
