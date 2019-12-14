package logic

import (
	"errors"
	"github.com/go-ginger/models"
)

func (base *BaseLogicHandler) DoCreate(request models.IRequest) (result interface{}, err error) {
	err = base.handleRequestFunction(base.IBaseLogicHandler.BeforeCreate, request)
	if err != nil {
		return
	}
	result, err = base.IBaseLogicHandler.Create(request)
	if err != nil {
		return
	}
	err = base.handleRequestFunction(base.IBaseLogicHandler.AfterCreate, request)
	if err != nil {
		return
	}
	return
}

func (base *BaseLogicHandler) BeforeCreate(request models.IRequest) (err error) {
	return
}

func (base *BaseLogicHandler) Create(request models.IRequest) (result interface{}, err error) {
	if base.DataHandler != nil {
		err = base.DataHandler.BeforeInsert(request)
		if err != nil {
			return
		}
		result, err = base.DataHandler.Insert(request)
		if err != nil {
			return nil, err
		}
		err = base.DataHandler.AfterInsert(request)
		if err != nil {
			return
		}
		return result, err
	}
	err = errors.New("data handler is not initialized")
	return
}

func (base *BaseLogicHandler) AfterCreate(request models.IRequest) (err error) {
	return
}
