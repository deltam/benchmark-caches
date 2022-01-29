package main

import (
	"testing"
	"time"

	cache4 "github.com/akyoto/cache"
	cache2 "github.com/muesli/cache2go"
	cache3 "github.com/patrickmn/go-cache"
	cache1 "github.com/rif/cache2go"
)

func BenchmarkSet(b *testing.B) {
	b.Run("rif/cache2go", func(b *testing.B) {
		c := cache1.New(10, time.Second)

		b.StartTimer()
		for i := 0; i < b.N; i++ {
			c.Set("hoge", 123)
		}
	})

	b.Run("muesli/cache2go", func(b *testing.B) {
		c := cache2.Cache("mycache")

		b.StartTimer()
		for i := 0; i < b.N; i++ {
			c.Add("hoge", time.Second, 123)
		}
	})

	b.Run("patrickmn/go-cache", func(b *testing.B) {
		c := cache3.New(time.Second, time.Minute)

		b.StartTimer()
		for i := 0; i < b.N; i++ {
			c.Set("hoge", 123, cache3.DefaultExpiration)
		}
	})

	b.Run("akyoto/cache", func(b *testing.B) {
		c := cache4.New(time.Second)

		b.StartTimer()
		for i := 0; i < b.N; i++ {
			c.Set("hoge", 123, time.Second)
		}
	})
}

func BenchmarkGet(b *testing.B) {
	b.Run("rif/cache2go", func(b *testing.B) {
		c := cache1.New(10, time.Second)

		c.Set("hoge", 123)
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_, ok := c.Get("hoge")
			if !ok {
				c.Set("hoge", 123)
			}
		}
	})

	b.Run("muesli/cache2go", func(b *testing.B) {
		c := cache2.Cache("mycache")

		c.Add("hoge", time.Second, 123)
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_, err := c.Value("hoge")
			if err != nil {
				c.Add("hoge", time.Second, 123)
			}
		}
	})

	b.Run("patrickmn/go-cache", func(b *testing.B) {
		c := cache3.New(time.Second, time.Minute)
		c.Set("hoge", 123, cache3.DefaultExpiration)

		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_, ok := c.Get("hoge")
			if !ok {
				c.Set("hoge", 123, cache3.DefaultExpiration)
			}
		}
	})

	b.Run("akyoto/cache", func(b *testing.B) {
		c := cache4.New(time.Second)
		c.Set("hoge", 123, time.Second)

		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_, ok := c.Get("hoge")
			if !ok {
				c.Set("hoge", 123, time.Second)
			}
		}
	})
}
