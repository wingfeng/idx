package cache

import (
	"fmt"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/wingfeng/idx/conf"
)

var Service *gcache.Cache

func init() {
	option := conf.Options

	link := fmt.Sprintf("%s:%d", option.RedisHost, option.RedisPort)
	var (
		err error

		redisConfig = &gredis.Config{
			Address: link,
			Db:      option.RedisDB,
		}
	)
	Service = gcache.New()
	// Create redis client object.
	redis, err := gredis.New(redisConfig)
	if err != nil {
		panic(err)
	}
	// Create redis cache adapter and set it to cache object.
	Service.SetAdapter(gcache.NewAdapterRedis(redis))
}
