package logic

import (
	"errors"
	"github.com/kulichak/dl"
	"github.com/kulichak/models"
)

type IBaseLogicHandler interface {
	Create(request *models.Request) (*models.IBaseModel, error)
	Paginate(request *models.Request) (*models.PaginateResult, error)
	Get(request *models.Request) (*models.IBaseModel, error)
	Update(request *models.Request) error
	Delete(request *models.Request) error
	Init(dataHandler dl.IBaseDbHandler)
}

type BaseLogicHandler struct {
	IBaseLogicHandler

	DataHandler dl.IBaseDbHandler
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

func (base *BaseLogicHandler) Create(request *models.Request) (*models.IBaseModel, error) {
	if base.DataHandler != nil {
		base.DataHandler.BeforeInsert(request)
		result, err := base.DataHandler.Insert(request)
		base.DataHandler.AfterInsert(request)
		return result, err
	}
	return nil, errors.New("data handler is not initialized")
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

func (base *BaseLogicHandler) Get(request *models.Request) (*models.IBaseModel, error) {
	if base.DataHandler != nil {
		base.handleRequestFunction(base.DataHandler.BeforeQuery, request)
		result, err := base.DataHandler.Get(request)
		base.handleRequestFunction(base.DataHandler.AfterQuery, request)
		return result, err
	}
	return nil, errors.New("data handler is not initialized")
}

func (base *BaseLogicHandler) Update(request *models.Request) error {
	if base.DataHandler != nil {
		base.handleRequestFunction(base.DataHandler.BeforeUpdate, request)
		err := base.DataHandler.Update(request)
		base.handleRequestFunction(base.DataHandler.AfterUpdate, request)
		return err
	}
	return errors.New("data handler is not initialized")
}

func (base *BaseLogicHandler) Delete(request *models.Request) error {
	if base.DataHandler != nil {
		base.handleRequestFunction(base.DataHandler.BeforeDelete, request)
		err := base.DataHandler.Delete(request)
		base.handleRequestFunction(base.DataHandler.AfterDelete, request)
		return err
	}
	return errors.New("data handler is not initialized")
}
