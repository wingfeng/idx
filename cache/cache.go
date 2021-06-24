package cache

import (
	"context"
	"sync"
	"time"
)

type ICacheProvider interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{}, duration time.Duration)
}

type CacheHelper struct {
	Provider ICacheProvider
}

func (ch *CacheHelper) Get(context context.Context, key string, GetFunc func() interface{}, duration time.Duration) (interface{}, error) {
	if ch.Provider != nil {
		result, exist := ch.Provider.Get(key)
		if exist && result != nil {
			return result, nil
		}
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	c := make(chan interface{})
	go func(k string) {
		obj := GetFunc()
		c <- obj
	}(key)
	obj := <-c
	ch.Provider.Set(key, obj, duration)
	return obj, nil
}
