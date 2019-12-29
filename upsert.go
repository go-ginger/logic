package logic

import (
	"errors"
	"github.com/go-ginger/models"
)

func (base *BaseLogicHandler) DoUpsert(request models.IRequest) (err error) {
	err = base.handleRequestFunction(base.IBaseLogicHandler.BeforeUpsert, request)
	if err != nil {
		return
	}
	err = base.IBaseLogicHandler.Upsert(request)
	if err != nil {
		return
	}
	err = base.handleRequestFunction(base.IBaseLogicHandler.AfterUpsert, request)
	if err != nil {
		return
	}
	return
}

func (base *BaseLogicHandler) BeforeUpsert(request models.IRequest) (err error) {
	return
}

func (base *BaseLogicHandler) Upsert(request models.IRequest) (err error) {
	if base.DataHandler != nil {
		err = base.DataHandler.DoUpsert(request)
		return
	}
	return errors.New("data handler is not initialized")
}

func (base *BaseLogicHandler) AfterUpsert(request models.IRequest) (err error) {
	return
}
