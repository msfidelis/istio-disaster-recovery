package memory_cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var instance *cache.Cache

// Memory Cache Singleton - Used to store key / values into container memory
func GetInstance() *cache.Cache {
	if instance == nil {
		instance = cache.New(5*time.Minute, 10*time.Minute)
	}
	return instance
}
