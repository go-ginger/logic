package logic

import (
	"errors"
	"github.com/go-ginger/models"
)

func (base *BaseLogicHandler) DoCreate(request models.IRequest) (result interface{}, err error) {
	base.handleRequestFunction(base.LogicHandler.Model, request)
	base.handleRequestFunction(base.LogicHandler.BeforeCreate, request)
	result, err = base.LogicHandler.Create(request)
	if err != nil {
		return
	}
	base.handleRequestFunction(base.LogicHandler.AfterCreate, request)
	return
}

func (base *BaseLogicHandler) BeforeCreate(request models.IRequest) {
}

func (base *BaseLogicHandler) Create(request models.IRequest) (result interface{}, err error) {
	if base.DataHandler != nil {
		base.DataHandler.BeforeInsert(request)
		result, err = base.DataHandler.Insert(request)
		if err != nil {
			return nil, err
		}
		base.DataHandler.AfterInsert(request)
		return result, err
	}
	err = errors.New("data handler is not initialized")
	return
}

func (base *BaseLogicHandler) AfterCreate(request models.IRequest) {
}
