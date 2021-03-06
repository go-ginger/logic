package logic

import (
	"github.com/go-ginger/dl"
	"github.com/go-ginger/models"
)

type IBaseLogicHandler interface {
	Init(logicHandler IBaseLogicHandler, dataHandler dl.IBaseDbHandler)
	GetDataHandler() dl.IBaseDbHandler

	DoCreate(request models.IRequest) (result models.IBaseModel, err error)
	DoPaginate(request models.IRequest) (result *models.PaginateResult, err error)
	DoGet(request models.IRequest) (result models.IBaseModel, err error)
	DoUpdate(request models.IRequest) (err error)
	DoUpsert(request models.IRequest) (err error)
	DoDelete(request models.IRequest) (err error)

	BeforeCreate(request models.IRequest) (err error)
	Create(request models.IRequest) (result models.IBaseModel, err error)
	AfterCreate(request models.IRequest) (err error)

	BeforeQuery(request models.IRequest) (err error)
	Paginate(request models.IRequest) (result *models.PaginateResult, err error)
	Get(request models.IRequest) (result models.IBaseModel, err error)
	First(request models.IRequest) (result models.IBaseModel, err error)
	Exists(request models.IRequest) (exists bool, err error)
	Count(request models.IRequest) (count uint64, err error)
	AfterQuery(request models.IRequest, result models.IBaseModel) (err error)

	BeforeUpdate(request models.IRequest) (err error)
	Update(request models.IRequest) (err error)
	AfterUpdate(request models.IRequest) (err error)

	BeforeUpsert(request models.IRequest) (err error)
	Upsert(request models.IRequest) (err error)
	AfterUpsert(request models.IRequest) (err error)

	BeforeDelete(request models.IRequest) (err error)
	Delete(request models.IRequest) (err error)
	AfterDelete(request models.IRequest) (err error)
}

type BaseLogicHandler struct {
	IBaseLogicHandler

	DataHandler dl.IBaseDbHandler
}

func (base *BaseLogicHandler) Init(logicHandler IBaseLogicHandler, dataHandler dl.IBaseDbHandler) {
	base.DataHandler = dataHandler
	base.IBaseLogicHandler = logicHandler
}

func (base *BaseLogicHandler) handleRequestFunction(
	function func(request models.IRequest) (err error), request models.IRequest) (err error) {
	if function != nil {
		err = function(request)
	}
	return
}

func (base *BaseLogicHandler) GetDataHandler() dl.IBaseDbHandler {
	return base.DataHandler
}

func (base *BaseLogicHandler) handleRequestModelFunction(
	function func(request models.IRequest, model models.IBaseModel) (err error),
	request models.IRequest, model models.IBaseModel) (err error) {
	if function != nil {
		err = function(request, model)
	}
	return
}
