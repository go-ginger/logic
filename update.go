package logic

import (
	"errors"
	"github.com/kulichak/models"
)

func (base *BaseLogicHandler) DoUpdate(request *models.Request) error {
	base.handleRequestFunction(base.BeforeUpdate, request)
	err := base.LogicHandler.Update(request)
	if err != nil {
		return err
	}
	base.handleRequestFunction(base.LogicHandler.AfterUpdate, request)
	return err
}

func (base *BaseLogicHandler) BeforeUpdate(request *models.Request) {
}

func (base *BaseLogicHandler) Update(request *models.Request) error {
	if base.DataHandler != nil {
		base.handleRequestFunction(base.DataHandler.BeforeUpdate, request)
		err := base.DataHandler.Update(request)
		if err != nil {
			return err
		}
		base.handleRequestFunction(base.DataHandler.AfterUpdate, request)
		return err
	}
	return errors.New("data handler is not initialized")
}

func (base *BaseLogicHandler) AfterUpdate(request *models.Request) {
}
