package dto

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
