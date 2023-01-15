package app

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

// RsContext request -> mid with RsCtx -> handler(ctx, req, res)
type RsContext struct {
	ctx context.Context
	env string // init
}

func (c RsContext) GetEnv() string {
	return c.env
}

func (c RsContext) Context() context.Context {
	return c.ctx
}

var DBClient = &sqlx.DB{}

var RedisClient = &redis.Client{}
