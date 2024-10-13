package gateway

import (
	"article/pkg/config"
	"article/pkg/elasticsearch"
	"article/pkg/email"
	"article/pkg/redis"
	"context"
	"gorm.io/gorm"
)

type BaseApi struct {
	ctx   context.Context
	db    *gorm.DB
	es    *elasticsearch.Elasticsearch
	rdb   *redis.Client
	conf  *config.GlobalConfig
	email *email.Server
}

func NewBaseApi(db *gorm.DB, es *elasticsearch.Elasticsearch, rdb *redis.Client, conf *config.GlobalConfig, email *email.Server) *BaseApi {
	return &BaseApi{
		ctx:   context.Background(),
		db:    db,
		es:    es,
		rdb:   rdb,
		conf:  conf,
		email: email,
	}
}
