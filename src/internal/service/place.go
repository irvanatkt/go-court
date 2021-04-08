package service

import (
	"context"
	"log"

	"court.com/src/internal/entity/dto"
	"court.com/src/internal/repo"
)

type PlaceService interface {
	GetPlaceById(ctx context.Context, id int64) dto.PlaceDtl
}

type PlaceServiceImpl struct {
	r repo.LocationRepo
}

func (p *PlaceServiceImpl) GetPlaceById(ctx context.Context, id int64) dto.PlaceDtl {
	r, err := p.r.GetPlaceById(ctx, id)
	if err != nil {
		log.Println(err)
		return dto.PlaceDtl{}
	}
	return r.ToDTO()
}

func New(pRepo repo.LocationRepo) PlaceService {
	return &PlaceServiceImpl{pRepo}
}
