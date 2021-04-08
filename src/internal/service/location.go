package service

import (
	"context"
	"log"

	"court.com/src/internal/entity/dto"
	"court.com/src/internal/repo"
)

type Location interface {
	GetGymnasiumByID(ctx context.Context, ID int64) dto.Gymnasium
}

type LocationImpl struct {
	repo repo.LocationRepo
}

func NewLocationSvc(pRepo repo.LocationRepo) Location {
	return &LocationImpl{pRepo}
}

func (l *LocationImpl) GetGymnasiumByID(ctx context.Context, ID int64) dto.Gymnasium {
	gym, err := l.repo.GetGymnasiumByID(ctx, ID)
	if err != nil {
		log.Println(err)
		return dto.Gymnasium{}
	}
	return gym.ToDTO()
}
