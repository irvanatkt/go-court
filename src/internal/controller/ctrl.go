package controller

import "court.com/src/internal/service"

type Controller struct {
	placeSvc    service.PlaceService
	locationSvc service.Location
}

func Init(placeSvc service.PlaceService, locSvc service.Location) *Controller {
	return &Controller{placeSvc: placeSvc, locationSvc: locSvc}
}
