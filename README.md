# Benchmark Caches

私的な調査メモ

```
goos: darwin
goarch: amd64
pkg: github.com/deltam/benchmark-caches
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
BenchmarkRifCache2go/Set-8    	 2655990	       448.3 ns/op	     142 B/op	       4 allocs/op
BenchmarkRifCache2go/Get-8    	 2040366	       592.0 ns/op	     166 B/op	       6 allocs/op
BenchmarkMuesliCache2go/Set-8 	 1000000	      1187 ns/op	     330 B/op	       6 allocs/op
BenchmarkMuesliCache2go/Get-8 	 2653855	      1053 ns/op	     216 B/op	       6 allocs/op
BenchmarkPatrickmnGoCache/Set-8         	 2201108	       542.7 ns/op	      23 B/op	       1 allocs/op
BenchmarkPatrickmnGoCache/Get-8         	 2696368	       482.1 ns/op	      29 B/op	       2 allocs/op
BenchmarkAkyotoCache/Set-8              	 1000000	      1107 ns/op	     226 B/op	       7 allocs/op
BenchmarkAkyotoCache/Get-8              	 2737189	      1375 ns/op	     157 B/op	       6 allocs/op
BenchmarkRistretto/Set-8                	 1697214	       890.9 ns/op	     275 B/op	       4 allocs/op
BenchmarkRistretto/Get-8                	 2583714	       482.0 ns/op	      47 B/op	       3 allocs/op
BenchmarkNaiveMap/Set-8                 	 2855382	       664.3 ns/op	     151 B/op	       3 allocs/op
BenchmarkNaiveMap/Get-8                 	 2496924	       488.4 ns/op	      23 B/op	       1 allocs/op
PASS
ok  	github.com/deltam/benchmark-caches	78.008s
```
