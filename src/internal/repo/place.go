package repo

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"court.com/src/constants"
	"court.com/src/internal/entity/dao"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PlaceRepo interface {
	GetPlaceById(ctx context.Context, id int64) (*dao.PlaceDtl, error)
}

type PlaceRepoImpl struct {
	mongoCli *mongo.Client
	redisCli *redis.Client
}

func New(cli *mongo.Client, redisCli *redis.Client) PlaceRepo {
	return &PlaceRepoImpl{mongoCli: cli, redisCli: redisCli}
}

func (p *PlaceRepoImpl) GetPlaceById(ctx context.Context, id int64) (*dao.PlaceDtl, error) {
	// get from cache
	result := p.getCachePlaceById(ctx, id)
	if result != nil {
		return result, nil
	}

	// get from db if empty
	courts := p.mongoCli.Database("court").Collection("courts")
	qId := bson.D{{"id", id}}
	err := courts.FindOne(ctx, qId).Decode(&result)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			log.Println("not found")
			return nil, err
		}
		return nil, err
	}

	// set to cache
	err = p.setCachePlaceById(ctx, *result)

	return result, nil
}

func (p *PlaceRepoImpl) getCachePlaceById(ctx context.Context, id int64) *dao.PlaceDtl {
	log.Printf("Get ID:%d from cache\n", id)
	result := p.redisCli.Get(ctx, fmt.Sprintf(constants.CachePlaceID, id))
	if result.Err() == redis.Nil {
		return nil
	}
	res, err := result.Result()
	if err != nil {
		log.Println(err)
		return nil
	}

	var dtl dao.PlaceDtl
	err = json.Unmarshal([]byte(res), &dtl)
	if err != nil {
		return nil
	}
	return &dtl
}

func (p *PlaceRepoImpl) setCachePlaceById(ctx context.Context, dtl dao.PlaceDtl) error {
	log.Printf("Set ID:%d to cache\n", dtl.ID)
	byteVal, err := json.Marshal(dtl)
	if err != nil {
		return err
	}
	stat := p.redisCli.Set(ctx, fmt.Sprintf(constants.CachePlaceID, dtl.ID), byteVal, time.Duration(constants.CachePlaceIDTTL*time.Second))
	log.Println(stat)
	return nil
}
