package logic

import (
	"errors"
	"github.com/kulichak/dl"
	"github.com/kulichak/models"
)

type IBaseLogicHandler interface {
	Create(request models.IRequest) (*models.IBaseModel, error)
	Paginate(request models.IRequest) (*models.PaginateResult, error)
	Get(request models.IRequest) (*models.IBaseModel, error)
	Update(request models.IRequest) error
	Delete(request models.IRequest) error
}

type BaseLogicHandler struct {
	IBaseLogicHandler

	DataHandler dl.IBaseDbHandler
}

func (base *BaseLogicHandler) handleRequestFunction(
	function func(request models.IRequest),	request models.IRequest) {
	if function != nil {
		function(request)
	}
}

func (base *BaseLogicHandler) Create(request models.IRequest) (*models.IBaseModel, error) {
	if base.DataHandler != nil {
		base.DataHandler.BeforeInsert(request)
		result, err := base.DataHandler.Insert(request)
		base.DataHandler.AfterInsert(request)
		return result, err
	}
	return nil, errors.New("data handler is not initialized")
}

func (base *BaseLogicHandler) Paginate(request models.IRequest) (*models.PaginateResult, error) {
	if base.DataHandler != nil {
		base.handleRequestFunction(base.DataHandler.BeforeQuery, request)
		result, err := base.DataHandler.Paginate(request)
		base.handleRequestFunction(base.DataHandler.AfterQuery, request)
		return result, err
	}
	return nil, errors.New("data handler is not initialized")
}

func (base *BaseLogicHandler) Get(request models.IRequest) (*models.IBaseModel, error) {
	if base.DataHandler != nil {
		base.handleRequestFunction(base.DataHandler.BeforeQuery, request)
		result, err := base.DataHandler.Get(request)
		base.handleRequestFunction(base.DataHandler.AfterQuery, request)
		return result, err
	}
	return nil, errors.New("data handler is not initialized")
}

func (base *BaseLogicHandler) Update(request models.IRequest) error {
	if base.DataHandler != nil {
		base.handleRequestFunction(base.DataHandler.BeforeUpdate, request)
		err := base.DataHandler.Update(request)
		base.handleRequestFunction(base.DataHandler.AfterUpdate, request)
		return err
	}
	return errors.New("data handler is not initialized")
}

func (base *BaseLogicHandler) Delete(request models.IRequest) error {
	if base.DataHandler != nil {
		base.handleRequestFunction(base.DataHandler.BeforeDelete, request)
		err := base.DataHandler.Delete(request)
		base.handleRequestFunction(base.DataHandler.AfterDelete, request)
		return err
	}
	return errors.New("data handler is not initialized")
}
