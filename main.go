package main

import (
	"article/internal/gateway"
	"article/pkg/config"
	"article/pkg/elasticsearch"
	"article/pkg/email"
	"article/pkg/mysql"
	"article/pkg/redis"
	"article/pkg/rules"
	"context"
	"fmt"
)

func main() {
	conf, err := config.NewGlobalConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	rdb := redis.NewClient(conf)
	db, err := mysql.NewClient(conf)
	if err != nil {
		fmt.Println(err)
		return
	}

	es, err := elasticsearch.NewClient(context.Background(), conf, "article")
	if err != nil {
		fmt.Println(err)
		return
	}

	emailServer := email.NewServer(conf)
	enforcer := rules.NewEnforcer(db)

	base := gateway.NewBaseApi(db, es, rdb, conf, emailServer, enforcer)
	r := gateway.NewRouter(base)
	if err := r.Run(); err != nil {
		fmt.Println("err:", err)
		return
	}
}
