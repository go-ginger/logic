package logic

import (
	"errors"
	"github.com/kulichak/models"
)

func (base *BaseLogicHandler) DoDelete(request models.IRequest) error {
	base.handleRequestFunction(base.LogicHandler.BeforeDelete, request)
	err := base.LogicHandler.Delete(request)
	if err != nil {
		return err
	}
	base.handleRequestFunction(base.LogicHandler.AfterDelete, request)
	return err
}

func (base *BaseLogicHandler) BeforeDelete(request models.IRequest) {
}

func (base *BaseLogicHandler) Delete(request models.IRequest) error {
	if base.DataHandler != nil {
		base.handleRequestFunction(base.DataHandler.BeforeDelete, request)
		err := base.DataHandler.Delete(request)
		if err != nil {
			return err
		}
		base.handleRequestFunction(base.DataHandler.AfterDelete, request)
		return err
	}
	return errors.New("data handler is not initialized")
}

func (base *BaseLogicHandler) AfterDelete(request models.IRequest) {
}
