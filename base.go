package logic

import (
	"github.com/kulichak/dl"
	"github.com/kulichak/models"
)

type IBaseLogicHandler interface {
	Init(dataHandler dl.IBaseDbHandler)
	Model(request *models.Request)
	Models(request *models.Request)

	DoCreate(request *models.Request) (interface{}, error)
	DoPaginate(request *models.Request) (*models.PaginateResult, error)
	DoGet(request *models.Request) (interface{}, error)
	DoUpdate(request *models.Request) error
	DoDelete(request *models.Request) error

	BeforeCreate(request *models.Request)
	Create(request *models.Request) (interface{}, error)
	AfterCreate(request *models.Request)

	BeforeQuery(request *models.Request)
	Paginate(request *models.Request) (*models.PaginateResult, error)
	Get(request *models.Request) (interface{}, error)
	AfterQuery(request *models.Request)

	BeforeUpdate(request *models.Request)
	Update(request *models.Request) error
	AfterUpdate(request *models.Request)

	BeforeDelete(request *models.Request)
	Delete(request *models.Request) error
	AfterDelete(request *models.Request)
}

type BaseLogicHandler struct {
	IBaseLogicHandler

	LogicHandler IBaseLogicHandler
	DataHandler  dl.IBaseDbHandler
}

func (base *BaseLogicHandler) Init(dataHandler dl.IBaseDbHandler) {
	base.DataHandler = dataHandler
}

func (base *BaseLogicHandler) handleRequestFunction(
	function func(request *models.Request), request *models.Request) {
	if function != nil {
		function(request)
	}
}

func (base *BaseLogicHandler) Model(request *models.Request) {
}

func (base *BaseLogicHandler) Models(request *models.Request) {
}
