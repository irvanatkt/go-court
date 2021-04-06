package dao

import "court.com/src/internal/entity/dto"

type PlaceDtl struct {
	ID int64
}

func (p PlaceDtl) ToDTO() dto.PlaceDtl {
	return dto.PlaceDtl{ID: p.ID}
}
