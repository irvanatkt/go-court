package service

import (
	"log"

	"court.com/src/internal/entity/dto"
	"court.com/src/internal/repo"
)

type PlaceService interface {
	GetPlaceById(id int64) dto.PlaceDtl
}

type PlaceServiceImpl struct {
	r repo.PlaceRepo
}

func (p *PlaceServiceImpl) GetPlaceById(id int64) dto.PlaceDtl {
	r, err := p.r.GetPlaceById(id)
	if err != nil {
		log.Println(err)
		return dto.PlaceDtl{}
	}
	return r.ToDTO()
}

func New(pRepo repo.PlaceRepo) PlaceService {
	return &PlaceServiceImpl{pRepo}
}
