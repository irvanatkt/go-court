package controller

import "court.com/src/internal/service"

type Controller struct {
	placeSvc service.PlaceService
}

func Init(placeSvc service.PlaceService) *Controller {
	return &Controller{placeSvc: placeSvc}
}
