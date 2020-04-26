package datasource

import (
	"github.com/bbcloudGroup/gothic/database"
	"github.com/go-redis/redis/v7"
)

type Cache struct {
	*redis.Client
}

func NewCache() Cache {
	return Cache{database.R()}
}
