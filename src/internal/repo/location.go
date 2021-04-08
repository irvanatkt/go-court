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

type LocationRepo interface {
	GetPlaceById(ctx context.Context, id int64) (*dao.PlaceDtl, error)
	GetGymnasiumByID(ctx context.Context, id int64) (*dao.Gymnasium, error)
}

type LocationRepoImpl struct {
	mongoCli *mongo.Client
	redisCli *redis.Client
}

func New(cli *mongo.Client, redisCli *redis.Client) LocationRepo {
	return &LocationRepoImpl{mongoCli: cli, redisCli: redisCli}
}

func (p *LocationRepoImpl) GetPlaceById(ctx context.Context, id int64) (*dao.PlaceDtl, error) {
	// get from cache
	result := p.getCachePlaceById(ctx, id)
	if result != nil {
		return result, nil
	}

	// get from db if empty
	courts := p.mongoCli.Database(constants.CourtDB).Collection(constants.CourtsCollection)
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

func (p *LocationRepoImpl) GetGymnasiumByID(ctx context.Context, id int64) (*dao.Gymnasium, error) {
	// get from cache
	result := p.getCacheGymnasiumById(ctx, id)
	if result != nil {
		return result, nil
	}

	// get from db if empty
	courts := p.mongoCli.Database(constants.CourtDB).Collection(constants.GymCollection)
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
	err = p.setCacheGymnasiumById(ctx, *result)

	return result, nil
}

/*
 private methods
*/

func (p *LocationRepoImpl) getFromRedis(ctx context.Context, key string) (string, error) {
	log.Printf("GET key %s from cache\n", key)
	result := p.redisCli.Get(ctx, key)
	if result.Err() == redis.Nil {
		return "", fmt.Errorf("empty cache")
	}
	res, err := result.Result()
	if err != nil {
		log.Println(err)
		return "", err
	}
	return res, nil
}

func (p *LocationRepoImpl) getCachePlaceById(ctx context.Context, id int64) *dao.PlaceDtl {
	log.Printf("Get place ID:%d from cache\n", id)
	res, err := p.getFromRedis(ctx, fmt.Sprintf(constants.CachePlaceID, id))
	if err != nil {
		return nil
	}

	var dtl dao.PlaceDtl
	err = json.Unmarshal([]byte(res), &dtl)
	if err != nil {
		return nil
	}
	return &dtl
}

func (p *LocationRepoImpl) getCacheGymnasiumById(ctx context.Context, id int64) *dao.Gymnasium {
	log.Printf("Get gymnasium ID:%d from cache\n", id)
	res, err := p.getFromRedis(ctx, fmt.Sprintf(constants.CacheGymID, id))
	if err != nil {
		return nil
	}

	var dtl dao.Gymnasium
	err = json.Unmarshal([]byte(res), &dtl)
	if err != nil {
		return nil
	}
	return &dtl
}

func (p *LocationRepoImpl) setCache(ctx context.Context, key string, value []byte, TTL int) error {
	stat := p.redisCli.Set(ctx, key, []byte(value), time.Duration(TTL)*time.Second)
	if stat.Err() != nil {
		return stat.Err()
	}
	return nil
}

func (p *LocationRepoImpl) setCachePlaceById(ctx context.Context, dtl dao.PlaceDtl) error {
	log.Printf("Set place ID:%d to cache\n", dtl.ID)
	byteVal, err := json.Marshal(dtl)
	if err != nil {
		return err
	}
	return p.setCache(ctx, fmt.Sprintf(constants.CachePlaceID, dtl.ID), byteVal, constants.CachePlaceIDTTL)
}

func (p *LocationRepoImpl) setCacheGymnasiumById(ctx context.Context, dtl dao.Gymnasium) error {
	log.Printf("Set gym ID:%d to cache\n", dtl.ID)
	byteVal, err := json.Marshal(dtl)
	if err != nil {
		return err
	}
	return p.setCache(ctx, fmt.Sprintf(constants.CacheGymID, dtl.ID), byteVal, constants.CachePlaceIDTTL)
}
