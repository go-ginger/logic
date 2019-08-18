package logic

import (
	"errors"
	"github.com/kulichak/models"
)

func (base *BaseLogicHandler) DoCreate(request *models.Request) (interface{}, error) {
	base.handleRequestFunction(base.LogicHandler.Model, request)
	base.handleRequestFunction(base.BeforeCreate, request)
	result, err := base.Create(request)
	base.handleRequestFunction(base.AfterCreate, request)
	return result, err
}

func (base *BaseLogicHandler) BeforeCreate(request *models.Request) {
}

func (base *BaseLogicHandler) Create(request *models.Request) (interface{}, error) {
	if base.DataHandler != nil {
		base.DataHandler.BeforeInsert(request)
		result, err := base.DataHandler.Insert(request)
		base.DataHandler.AfterInsert(request)
		return result, err
	}
	return nil, errors.New("data handler is not initialized")
}

func (base *BaseLogicHandler) AfterCreate(request *models.Request) {
}
