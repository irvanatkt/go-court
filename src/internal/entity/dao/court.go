package dao

import "court.com/src/internal/entity/dto"

type Gymnasium struct {
	ID         int64      `json:"id"`
	Name       string     `json:"name"`
	Courts     []CourtDtl `json:"courts"`
	BusinessID int64      `json:"business_id"`
	Place      PlaceDtl   `json:"place_dtl"`
	Capacity   int        `json:"capacity"`
}

type CourtDtl struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func (g Gymnasium) ToDTO() dto.Gymnasium {
	dtoPlace := dto.PlaceDtl{
		ID:        g.Place.ID,
		Latitude:  g.Place.Latitude,
		Longitude: g.Place.Longitude,
	}
	var dtoCourts []dto.CourtDtl
	for _, v := range g.Courts {
		dtoCourts = append(dtoCourts, dto.CourtDtl{ID: v.ID, Name: v.Name, Type: v.Type})
	}
	return dto.Gymnasium{
		ID:         g.ID,
		Name:       g.Name,
		BusinessID: g.BusinessID,
		Capacity:   g.Capacity,
		Place:      dtoPlace,
		Courts:     dtoCourts,
	}
}
