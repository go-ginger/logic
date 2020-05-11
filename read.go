package logic

import (
	"errors"
	"github.com/go-ginger/models"
)

func (base *BaseLogicHandler) DoPaginate(request models.IRequest) (result *models.PaginateResult, err error) {
	err = base.handleRequestFunction(base.IBaseLogicHandler.BeforeQuery, request)
	if err != nil {
		return
	}
	result, err = base.IBaseLogicHandler.Paginate(request)
	if err != nil {
		return
	}
	err = base.handleRequestParamFunction(base.IBaseLogicHandler.AfterQuery, request, result)
	if err != nil {
		return
	}
	return
}

func (base *BaseLogicHandler) DoGet(request models.IRequest) (result interface{}, err error) {
	err = base.handleRequestFunction(base.IBaseLogicHandler.BeforeQuery, request)
	if err != nil {
		return
	}
	result, err = base.IBaseLogicHandler.Get(request)
	if err != nil {
		return
	}
	err = base.handleRequestParamFunction(base.IBaseLogicHandler.AfterQuery, request, result)
	if err != nil {
		return
	}
	return
}

func (base *BaseLogicHandler) BeforeQuery(request models.IRequest) (err error) {
	return
}

func (base *BaseLogicHandler) Paginate(request models.IRequest) (result *models.PaginateResult, err error) {
	if base.DataHandler != nil {
		result, err = base.DataHandler.DoPaginate(request)
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
		result, err = base.DataHandler.DoGet(request)
		return
	}
	err = errors.New("data handler is not initialized")
	return
}

func (base *BaseLogicHandler) First(request models.IRequest) (result interface{}, err error) {
	if base.DataHandler != nil {
		result, err = base.DataHandler.DoGetFirst(request)
		return
	}
	err = errors.New("data handler is not initialized")
	return
}

func (base *BaseLogicHandler) AfterQuery(request models.IRequest, result interface{}) (err error) {
	return
}
