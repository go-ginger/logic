package logic

import (
	"errors"
	"github.com/kulichak/models"
)

func (base *BaseLogicHandler) DoDelete(request *models.Request) error {
	base.handleRequestFunction(base.BeforeDelete, request)
	err := base.Delete(request)
	base.handleRequestFunction(base.AfterDelete, request)
	return err
}

func (base *BaseLogicHandler) BeforeDelete(request *models.Request) {
}

func (base *BaseLogicHandler) Delete(request *models.Request) error {
	if base.DataHandler != nil {
		base.handleRequestFunction(base.DataHandler.BeforeDelete, request)
		err := base.DataHandler.Delete(request)
		base.handleRequestFunction(base.DataHandler.AfterDelete, request)
		return err
	}
	return errors.New("data handler is not initialized")
}

func (base *BaseLogicHandler) AfterDelete(request *models.Request) {
}
