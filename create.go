package logic

import (
	"errors"
	"github.com/kulichak/models"
)

func (base *BaseLogicHandler) DoCreate(request models.IRequest) (interface{}, error) {
	base.handleRequestFunction(base.LogicHandler.Model, request)
	base.handleRequestFunction(base.LogicHandler.BeforeCreate, request)
	result, err := base.LogicHandler.Create(request)
	if err != nil {
		return nil, err
	}
	base.handleRequestFunction(base.LogicHandler.AfterCreate, request)
	return result, err
}

func (base *BaseLogicHandler) BeforeCreate(request models.IRequest) {
}

func (base *BaseLogicHandler) Create(request models.IRequest) (interface{}, error) {
	if base.DataHandler != nil {
		base.DataHandler.BeforeInsert(request)
		result, err := base.DataHandler.Insert(request)
		if err != nil {
			return nil, err
		}
		base.DataHandler.AfterInsert(request)
		return result, err
	}
	return nil, errors.New("data handler is not initialized")
}

func (base *BaseLogicHandler) AfterCreate(request models.IRequest) {
}
