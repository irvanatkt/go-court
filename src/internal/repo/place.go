package repo

import "go.mongodb.org/mongo-driver/mongo"

type PlaceRepo interface {
	GetPlaceById(id int)
}

type PlaceRepoImpl struct {
	mongoCli *mongo.Client
}

func (p *PlaceRepoImpl) GetPlaceById(id int) {
	// p.mongoCli.Database()

}

func New(cli *mongo.Client) PlaceRepo {
	return &PlaceRepoImpl{mongoCli: cli}
}
