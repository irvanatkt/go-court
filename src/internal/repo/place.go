package repo

import (
	"context"
	"log"

	"court.com/src/internal/entity/dao"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PlaceRepo interface {
	GetPlaceById(id int64) (*dao.PlaceDtl, error)
}

type PlaceRepoImpl struct {
	mongoCli *mongo.Client
}

func (p *PlaceRepoImpl) GetPlaceById(id int64) (*dao.PlaceDtl, error) {
	var result dao.PlaceDtl
	courts := p.mongoCli.Database("court").Collection("courts")
	qId := bson.D{{"id", id}}
	err := courts.FindOne(context.Background(), qId).Decode(&result)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			log.Println("not found")
			return nil, err
		}
		return nil, err
	}
	return &result, nil
}

func New(cli *mongo.Client) PlaceRepo {
	return &PlaceRepoImpl{mongoCli: cli}
}
