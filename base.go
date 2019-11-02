package logic

import (
	"github.com/go-ginger/dl"
	"github.com/go-ginger/models"
)

type IBaseLogicHandler interface {
	Init(logicHandler IBaseLogicHandler, dataHandler dl.IBaseDbHandler)

	DoCreate(request models.IRequest) (result interface{}, err error)
	DoPaginate(request models.IRequest) (result *models.PaginateResult, err error)
	DoGet(request models.IRequest) (result interface{}, err error)
	DoUpdate(request models.IRequest) (err error)
	DoDelete(request models.IRequest) (err error)

	BeforeCreate(request models.IRequest)
	Create(request models.IRequest) (result interface{}, err error)
	AfterCreate(request models.IRequest)

	BeforeQuery(request models.IRequest)
	Paginate(request models.IRequest) (result *models.PaginateResult, err error)
	Get(request models.IRequest) (result interface{}, err error)
	AfterQuery(request models.IRequest)

	BeforeUpdate(request models.IRequest)
	Update(request models.IRequest) (err error)
	AfterUpdate(request models.IRequest)

	BeforeDelete(request models.IRequest)
	Delete(request models.IRequest) (err error)
	AfterDelete(request models.IRequest)
}

type BaseLogicHandler struct {
	IBaseLogicHandler

	LogicHandler IBaseLogicHandler
	DataHandler  dl.IBaseDbHandler
}

func (base *BaseLogicHandler) Init(logicHandler IBaseLogicHandler, dataHandler dl.IBaseDbHandler) {
	base.DataHandler = dataHandler
	base.LogicHandler = logicHandler
}

func (base *BaseLogicHandler) handleRequestFunction(
	function func(request models.IRequest), request models.IRequest) {
	if function != nil {
		function(request)
	}
}

func (base *BaseLogicHandler) handleRequestParamFunction(
	function func(request models.IRequest, param interface{}), request models.IRequest, param interface{}) {
	if function != nil {
		function(request, param)
	}
}
