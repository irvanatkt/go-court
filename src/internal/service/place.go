package service

import "court.com/src/internal/repo"

type PlaceService interface {
	GetPlaceById(id int)
}

type PlaceServiceImpl struct {
	r repo.PlaceRepo
}

func (p *PlaceServiceImpl) GetPlaceById(id int) {
	p.r.GetPlaceById(id)
}

func New(pRepo repo.PlaceRepo) PlaceService {
	return &PlaceServiceImpl{pRepo}
}
