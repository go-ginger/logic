package logic

import (
	"errors"
	"github.com/go-ginger/models"
)

func (base *BaseLogicHandler) DoUpdate(request models.IRequest) (err error) {
	err = base.handleRequestFunction(base.IBaseLogicHandler.BeforeUpdate, request)
	if err != nil {
		return
	}
	err = base.IBaseLogicHandler.Update(request)
	if err != nil {
		return
	}
	err = base.handleRequestFunction(base.IBaseLogicHandler.AfterUpdate, request)
	if err != nil {
		return
	}
	return
}

func (base *BaseLogicHandler) BeforeUpdate(request models.IRequest) (err error) {
	return
}

func (base *BaseLogicHandler) Update(request models.IRequest) (err error) {
	if base.DataHandler != nil {
		err = base.DataHandler.DoUpdate(request)
		return
	}
	return errors.New("data handler is not initialized")
}

func (base *BaseLogicHandler) AfterUpdate(request models.IRequest) (err error) {
	return
}
