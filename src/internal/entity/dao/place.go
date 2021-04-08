package dao

import "court.com/src/internal/entity/dto"

type PlaceDtl struct {
	ID        int64   `json:"id"`
	Latitude  float64 `json:"lat" bson:"lat"`
	Longitude float64 `json:"long" bson:"long"`
}

func (p PlaceDtl) ToDTO() dto.PlaceDtl {
	return dto.PlaceDtl{ID: p.ID}
}
