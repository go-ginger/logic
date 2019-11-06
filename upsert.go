package logic

import (
	"errors"
	"github.com/go-ginger/models"
)

func (base *BaseLogicHandler) DoUpsert(request models.IRequest) (err error) {
	base.handleRequestFunction(base.BeforeUpsert, request)
	err = base.LogicHandler.Upsert(request)
	if err != nil {
		return
	}
	base.handleRequestFunction(base.LogicHandler.AfterUpsert, request)
	return
}

func (base *BaseLogicHandler) BeforeUpsert(request models.IRequest) {
}

func (base *BaseLogicHandler) Upsert(request models.IRequest) (err error) {
	if base.DataHandler != nil {
		base.handleRequestFunction(base.DataHandler.BeforeUpsert, request)
		err = base.DataHandler.Upsert(request)
		if err != nil {
			return
		}
		base.handleRequestFunction(base.DataHandler.AfterUpsert, request)
		return
	}
	return errors.New("data handler is not initialized")
}

func (base *BaseLogicHandler) AfterUpsert(request models.IRequest) {
}
