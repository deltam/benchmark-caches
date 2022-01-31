package main

import (
	"fmt"
	"log"
	"testing"
	"time"

	cache4 "github.com/akyoto/cache"
	cache6 "github.com/coocood/freecache"
	cache5 "github.com/dgraph-io/ristretto"
	cache2 "github.com/muesli/cache2go"
	cache3 "github.com/patrickmn/go-cache"
	cache1 "github.com/rif/cache2go"
)

func key(i int) string { return fmt.Sprintf("key-%d", i) }

func BenchmarkRifCache2go(b *testing.B) {
	c := cache1.New(10, time.Minute)

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c.Set(key(i), 123)
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, ok := c.Get(key(i))
			if !ok {
				c.Set(key(i), 123)
			}
		}
	})
}

func BenchmarkMuesliCache2go(b *testing.B) {
	c := cache2.Cache("mycache")

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c.Add(key(i), time.Minute, 123)
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := c.Value(key(i))
			if err != nil {
				c.Add(key(i), time.Minute, 123)
			}
		}
	})
}

func BenchmarkPatrickmnGoCache(b *testing.B) {
	c := cache3.New(time.Minute, time.Minute)

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c.Set(key(i), 123, cache3.DefaultExpiration)
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, ok := c.Get(key(i))
			if !ok {
				c.Set(key(i), 123, cache3.DefaultExpiration)
			}
		}
	})
}

func BenchmarkAkyotoCache(b *testing.B) {
	c := cache4.New(time.Minute)

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c.Set(key(i), 123, time.Minute)
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, ok := c.Get(key(i))
			if !ok {
				c.Set(key(i), 123, time.Minute)
			}
		}
	})
}

func BenchmarkRistretto(b *testing.B) {
	c, err := cache5.NewCache(&cache5.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		log.Fatal(err)
	}

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c.Set(key(i), 123, 1)
		}
	})

	// wait for value to pass through buffers
	c.Wait()

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			v, ok := c.Get(key(i))
			if !ok {
				_ = v
			}
		}
	})
}

func BenchmarkFreecache(b *testing.B) {
	// In bytes, where 1024 * 1024 represents a single Megabyte, and 100 * 1024*1024 represents 100 Megabytes.
	const cacheSize = 100 * 1024 * 1024

	c := cache6.NewCache(cacheSize)

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c.Set([]byte(key(i)), []byte("123"), 600)
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := c.Get([]byte(key(i)))
			if err != nil {
				log.Println(err)
			}
		}
	})
}

func BenchmarkNaiveMap(b *testing.B) {
	type item struct {
		val       interface{}
		expiredAt time.Time
	}
	c := make(map[string]item, 1000)

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c[key(i)] = item{123, time.Now().Add(time.Minute)}
		}
	})

	Get := func(key string) (interface{}, bool) {
		itm, ok := c[key]
		if !ok {
			return nil, false
		}
		if time.Now().After(itm.expiredAt) {
			delete(c, key)
			return nil, false
		}
		return itm.val, true
	}

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			v, ok := Get(key(i))
			if !ok {
				_ = v
			}
		}
	})
}
