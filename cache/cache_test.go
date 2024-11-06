package cache

import (
	"context"
	"log/slog"
	"testing"
	"time"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/stretchr/testify/assert"

	_ "github.com/wingfeng/idx/conf"
	"github.com/wingfeng/idx/models"
	"github.com/wingfeng/idx/utils"
	"gorm.io/datatypes"
)

func TestCacheSet(t *testing.T) {
	// err := Set(context.Background(), "test", time.Minute, "test")
	// assert.NoError(t, err)
	// s := ""
	// err = Get(context.Background(), "test", &s)
	// assert.NoError(t, err)
	// assert.Equal(t, "test", s)

}
func TestGetOrLoad(t *testing.T) {
	ctx := context.Background()
	key := "test:test:1838872840128958464"
	c := setAdapter(ctx)
	u := &models.User{

		UserName: "test",
	}
	ttl := time.Minute * 5
	u.Claims = datatypes.JSON([]byte(`{"test":"test"}`))
	u.Id = utils.GeneratID()
	loadInFunc := false
	loadFunc := func(ctx context.Context) (interface{}, error) {
		t.Log("load from function")
		loadInFunc = true

		return u, nil

	}
	re := &models.User{}
	r, err := c.GetOrSetFunc(ctx, key, loadFunc, ttl)
	r.Struct(re)
	assert.NoError(t, err)
	assert.Equal(t, u, re)
	assert.True(t, loadInFunc)
	loadInFunc = false
	//for i := 0; i < 10; i++ {
	r, err = c.GetOrSetFunc(ctx, key, loadFunc, ttl)
	r.Struct(re)
	assert.False(t, loadInFunc)
	assert.Equal(t, u, re)
	assert.NoError(t, err)
	//}

}
func TestDelete(t *testing.T) {
	// s := ""
	// err := GetOrLoad(context.Background(), "test", load, &s, time.Minute)
	// assert.NoError(t, err)
	// assert.Equal(t, "hello world", s)
	// assert.NoError(t, Delete(context.Background(), "test"))
	// err = Get(context.Background(), "test", &s)
	// assert.Error(t, err)
	// assert.Nil(t, s)
}

func load(ctx context.Context, key interface{}) (interface{}, error) {
	slog.Info("load from function")
	return "hello world", nil

}

func setAdapter(ctx context.Context) *gcache.Cache {
	var (
		err error

		cache       = gcache.New()
		redisConfig = &gredis.Config{
			Address: "127.0.0.1:6379",
			Db:      0,
		}
	)
	// Create redis client object.
	redis, err := gredis.New(redisConfig)
	if err != nil {
		panic(err)
	}
	// Create redis cache adapter and set it to cache object.
	cache.SetAdapter(gcache.NewAdapterRedis(redis))
	return cache
	// // Set and Get using cache object.
	// err = cache.Set(ctx, cacheKey, cacheValue, time.Second)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(cache.MustGet(ctx, cacheKey).String())

	// // Get using redis client.
	// fmt.Println(redis.MustDo(ctx, "GET", cacheKey).String())
	// cache.GetOrSetFunc(ctx, cacheKey, func(ctx context.Context, key interface{}) (interface{}, error) {
	// 	return "hello", nil
	// })
}
