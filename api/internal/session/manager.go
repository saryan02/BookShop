package session

import (
	"github.com/alexedwards/scs/goredisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/redis/go-redis/v9"
)

func NewManager(redisClient *redis.Client) *scs.SessionManager {
	manager := scs.New()
	manager.Store = goredisstore.New(redisClient)

	return manager
}
