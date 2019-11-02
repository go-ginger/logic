package logic

import (
	"errors"
	"github.com/go-ginger/models"
)

func (base *BaseLogicHandler) DoUpdate(request models.IRequest) (err error) {
	base.handleRequestFunction(base.BeforeUpdate, request)
	err = base.LogicHandler.Update(request)
	if err != nil {
		return
	}
	base.handleRequestFunction(base.LogicHandler.AfterUpdate, request)
	return
}

func (base *BaseLogicHandler) BeforeUpdate(request models.IRequest) {
}

func (base *BaseLogicHandler) Update(request models.IRequest) (err error) {
	if base.DataHandler != nil {
		base.handleRequestFunction(base.DataHandler.BeforeUpdate, request)
		err = base.DataHandler.Update(request)
		if err != nil {
			return
		}
		base.handleRequestFunction(base.DataHandler.AfterUpdate, request)
		return
	}
	return errors.New("data handler is not initialized")
}

func (base *BaseLogicHandler) AfterUpdate(request models.IRequest) {
}
