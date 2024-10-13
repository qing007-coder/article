package gateway

import (
	"article/pkg/config"
	"article/pkg/elasticsearch"
	"article/pkg/redis"
	"gorm.io/gorm"
)

type BaseApi struct {
	db   *gorm.DB
	es   *elasticsearch.Elasticsearch
	rdb  *redis.Client
	conf *config.GlobalConfig
}

func NewBaseApi(db *gorm.DB, es *elasticsearch.Elasticsearch, rdb *redis.Client, conf *config.GlobalConfig) *BaseApi {
	return &BaseApi{
		db:   db,
		es:   es,
		rdb:  rdb,
		conf: conf,
	}
}
