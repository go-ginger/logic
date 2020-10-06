package logic

import (
	"errors"
	"fmt"
	"github.com/go-ginger/dl"
	"github.com/go-ginger/helpers"
	"github.com/go-ginger/models"
	"log"
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
	secondaryDbs := base.DataHandler.GetSecondaryDBs()
	if secondaryDbs != nil {
		for _, secondaryDB := range secondaryDbs {
			if secondaryDB.UpdateInBackgroundEnabled() {
				go func(db dl.IBaseDbHandler) {
					err = base.handleSecondaryUpdate(request, db)
					if err != nil {
						log.Println(fmt.Sprintf("error on handleSecondaryUpdate, err: %v", err))
						return
					}
				}(secondaryDB)
			} else {
				err = base.handleSecondaryUpdate(request, secondaryDB)
			}
		}
	}
	return
}

func (base *BaseLogicHandler) handleSecondaryUpdate(request models.IRequest, secondaryDB dl.IBaseDbHandler) (err error) {
	secondaryRequest := helpers.Clone(request).(models.IRequest)
	if secondaryDB.IsFullObjOnUpdateRequired() {
		objID := request.GetID()
		req := secondaryRequest.GetBaseRequest()
		req.Filters = &models.Filters{
			"id": objID,
		}
		item, e := base.IBaseLogicHandler.DoGet(secondaryRequest)
		if e != nil {
			err = e
			return
		}
		secondaryRequest.SetBody(item)
	}
	err = secondaryDB.DoUpdate(secondaryRequest)
	return
}
