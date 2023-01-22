package app

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

// EasyGoContext request -> mid with RsCtx -> handler(ctx, req, res)
type EasyGoContext struct {
	ctx context.Context
	env string // init
}

func (c EasyGoContext) GetEnv() string {
	return c.env
}

func (c EasyGoContext) Context() context.Context {
	return c.ctx
}

var DBClient = &sqlx.DB{}

var RedisClient = &redis.Client{}
