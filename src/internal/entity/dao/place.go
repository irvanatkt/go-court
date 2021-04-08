package dao

import "court.com/src/internal/entity/dto"

type PlaceDtl struct {
	ID        int64   `json:"id"`
	Latitude  float32 `json:"lat"`
	Longitude float32 `json:"long"`
}

func (p PlaceDtl) ToDTO() dto.PlaceDtl {
	return dto.PlaceDtl{ID: p.ID}
}
