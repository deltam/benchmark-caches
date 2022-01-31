package main

import (
	"fmt"
	"testing"
	"time"

	cache4 "github.com/akyoto/cache"
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
