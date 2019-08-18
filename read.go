package logic

import (
	"errors"
	"github.com/kulichak/models"
)

func (base *BaseLogicHandler) DoPaginate(request *models.Request) (*models.PaginateResult, error) {
	base.handleRequestFunction(base.LogicHandler.Models, request)
	base.handleRequestFunction(base.LogicHandler.BeforeQuery, request)
	result, err := base.LogicHandler.Paginate(request)
	base.handleRequestFunction(base.LogicHandler.AfterQuery, request)
	return result, err
}

func (base *BaseLogicHandler) DoGet(request *models.Request) (interface{}, error) {
	base.handleRequestFunction(base.LogicHandler.Model, request)
	base.handleRequestFunction(base.LogicHandler.BeforeQuery, request)
	result, err := base.LogicHandler.Get(request)
	base.handleRequestFunction(base.LogicHandler.AfterQuery, request)
	return result, err
}

func (base *BaseLogicHandler) BeforeQuery(request *models.Request) {
}

func (base *BaseLogicHandler) Paginate(request *models.Request) (*models.PaginateResult, error) {
	if base.DataHandler != nil {
		base.handleRequestFunction(base.DataHandler.BeforeQuery, request)
		result, err := base.DataHandler.Paginate(request)
		base.handleRequestFunction(base.DataHandler.AfterQuery, request)
		return result, err
	}
	return nil, errors.New("data handler is not initialized")
}

func (base *BaseLogicHandler) Get(request *models.Request) (interface{}, error) {
	if base.DataHandler != nil {
		base.handleRequestFunction(base.DataHandler.BeforeQuery, request)
		result, err := base.DataHandler.Get(request)
		base.handleRequestFunction(base.DataHandler.AfterQuery, request)
		return result, err
	}
	return nil, errors.New("data handler is not initialized")
}

func (base *BaseLogicHandler) AfterQuery(request *models.Request) {
}
