package logic

import (
	"errors"
	"github.com/go-ginger/models"
)

func (base *BaseLogicHandler) DoDelete(request models.IRequest) (err error) {
	err = base.handleRequestFunction(base.IBaseLogicHandler.BeforeDelete, request)
	if err != nil {
		return
	}
	err = base.IBaseLogicHandler.Delete(request)
	if err != nil {
		return
	}
	err = base.handleRequestFunction(base.IBaseLogicHandler.AfterDelete, request)
	if err != nil {
		return
	}
	return
}

func (base *BaseLogicHandler) BeforeDelete(request models.IRequest) (err error) {
	return
}

func (base *BaseLogicHandler) Delete(request models.IRequest) (err error) {
	if base.DataHandler != nil {
		err = base.handleRequestFunction(base.DataHandler.BeforeDelete, request)
		if err != nil {
			return
		}
		err = base.DataHandler.Delete(request)
		if err != nil {
			return
		}
		err = base.handleRequestFunction(base.DataHandler.AfterDelete, request)
		if err != nil {
			return
		}
		return
	}
	return errors.New("data handler is not initialized")
}

func (base *BaseLogicHandler) AfterDelete(request models.IRequest) (err error) {
	return
}
