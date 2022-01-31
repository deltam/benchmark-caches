# Benchmark Caches

私的な調査メモ

```
goos: darwin
goarch: amd64
pkg: github.com/deltam/benchmark-caches
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
BenchmarkRifCache2go/Set-8    	 2642464	       457.1 ns/op	     142 B/op	       4 allocs/op
BenchmarkRifCache2go/Get-8    	 1995303	       608.1 ns/op	     166 B/op	       6 allocs/op
BenchmarkMuesliCache2go/Set-8 	 1000000	      1196 ns/op	     330 B/op	       6 allocs/op
BenchmarkMuesliCache2go/Get-8 	 2514271	      1181 ns/op	     214 B/op	       6 allocs/op
BenchmarkPatrickmnGoCache/Set-8         	 2803508	       574.6 ns/op	      93 B/op	       2 allocs/op
BenchmarkPatrickmnGoCache/Get-8         	 2736442	       460.0 ns/op	      23 B/op	       1 allocs/op
BenchmarkAkyotoCache/Set-8              	 1000000	      1147 ns/op	     226 B/op	       7 allocs/op
BenchmarkAkyotoCache/Get-8              	 2667693	      1245 ns/op	     157 B/op	       6 allocs/op
BenchmarkRistretto/Set-8                	 1733604	       900.5 ns/op	     273 B/op	       4 allocs/op
BenchmarkRistretto/Get-8                	 2402091	       517.1 ns/op	      47 B/op	       3 allocs/op
BenchmarkFreecache/Set-8                	 2209300	       661.3 ns/op	      53 B/op	       1 allocs/op
BenchmarkFreecache/Get-8                	 1717724	       728.0 ns/op	      68 B/op	       4 allocs/op
BenchmarkNaiveMap/Set-8                 	 2805727	       600.7 ns/op	     125 B/op	       2 allocs/op
BenchmarkNaiveMap/Get-8                 	 2645494	       483.1 ns/op	      23 B/op	       1 allocs/op
PASS
ok  	github.com/deltam/benchmark-caches	87.525s
```
