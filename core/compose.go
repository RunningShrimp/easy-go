package core

import (
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

var DBClient = &sqlx.DB{}

var CacheClient = &redis.Client{}
