package logic

import (
	"errors"
	"github.com/go-ginger/models"
)

func (base *BaseLogicHandler) DoPaginate(request models.IRequest) (result *models.PaginateResult, err error) {
	base.handleRequestFunction(base.LogicHandler.BeforeQuery, request)
	result, err = base.LogicHandler.Paginate(request)
	if err != nil {
		return
	}
	base.handleRequestFunction(base.LogicHandler.AfterQuery, request)
	return
}

func (base *BaseLogicHandler) DoGet(request models.IRequest) (result interface{}, err error) {
	base.handleRequestFunction(base.LogicHandler.BeforeQuery, request)
	result, err = base.LogicHandler.Get(request)
	if err != nil {
		return
	}
	base.handleRequestFunction(base.LogicHandler.AfterQuery, request)
	return
}

func (base *BaseLogicHandler) BeforeQuery(request models.IRequest) {
}

func (base *BaseLogicHandler) Paginate(request models.IRequest) (result *models.PaginateResult, err error) {
	if base.DataHandler != nil {
		base.handleRequestFunction(base.DataHandler.BeforeQuery, request)
		result, err = base.DataHandler.Paginate(request)
		if err != nil {
			return
		}
		base.handleRequestParamFunction(base.DataHandler.AfterQuery, request, result)
		return
	}
	err = errors.New("data handler is not initialized")
	return
}

func (base *BaseLogicHandler) Get(request models.IRequest) (result interface{}, err error) {
	req := request.GetBaseRequest()
	if req.ID != nil {
		// handle id
		if req.Filters == nil {
			req.Filters = &models.Filters{}
		}
		(*req.Filters)["id"] = req.ID
	}

	if base.DataHandler != nil {
		base.handleRequestFunction(base.DataHandler.BeforeQuery, request)
		result, err = base.DataHandler.Get(request)
		if err != nil {
			return
		}
		base.handleRequestParamFunction(base.DataHandler.AfterQuery, request, result)
		return
	}
	err = errors.New("data handler is not initialized")
	return
}

func (base *BaseLogicHandler) AfterQuery(request models.IRequest) {
}
