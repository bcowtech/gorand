# gorand

gorand is a thread safe based on rand.  It was originally implemented in golang by gosl. Due to lightweight requirements, we only copy the some rnd function from gosl.

## Install

```console
go get -u -v github.com/bcowtech/gorand
```

## Usage

Let's start with a trivial example:

```go

package main

import (
	"math/rand"

	"github.com/bcowtech/gorand"
)

func main() {
	r := gorand.New(rand.NewSource(9527))
	r.Seed(9527)

	println("When seed is 9527. Int63r(0,10) return", r.Int63r(0, 10), ".")
}


```

Output:

```console

go run ./example/app.go
When seed is 9527. Int63r(0,10) return 2 .

```
  
## Benckmark

```console

go test -benchmem -bench .
goos: darwin
goarch: amd64
pkg: github.com/bcowtech/gorand
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_rand_Int63Threadsafe-12               79590735                17.03 ns/op            0 B/op          0 allocs/op
Benchmark_mt_Int63Threadsafe-12                 71069598                16.89 ns/op            0 B/op          0 allocs/op
Benchmark_rand_Int63ThreadsafeParallel-12       21650484                54.18 ns/op            0 B/op          0 allocs/op
Benchmark_rand_Int63Unthreadsafe-12             73113511                17.21 ns/op            0 B/op          0 allocs/op
Benchmark_mt_Int63ThreadsafeParallel-12         21074505                56.56 ns/op            0 B/op          0 allocs/op
Benchmark_rand_New-12                             144631              8839 ns/op               0 B/op          0 allocs/op
Benchmark_mt_New-12                             24594357                43.74 ns/op            0 B/op          0 allocs/op
Benchmark_rand_Intn1000-12                      57917715                21.31 ns/op            0 B/op          0 allocs/op
Benchmark_mt_Intn1000-12                        44490828                24.04 ns/op            0 B/op          0 allocs/op
Benchmark_rand_Int63n1000-12                    42276000                23.83 ns/op            0 B/op          0 allocs/op
Benchmark_mt_Int63n1000-12                      43607821                27.17 ns/op            0 B/op          0 allocs/op
Benchmark_rand_Int31n1000-12                    59763703                17.87 ns/op            0 B/op          0 allocs/op
Benchmark_mt_Int31n1000-12                      51766426                20.94 ns/op            0 B/op          0 allocs/op
Benchmark_rand_Float32-12                       75257120                15.27 ns/op            0 B/op          0 allocs/op
Benchmark_mt_Float32-12                         68844014                17.92 ns/op            0 B/op          0 allocs/op
Benchmark_rand_Float64-12                       79798725                17.66 ns/op            0 B/op          0 allocs/op
Benchmark_mt_Float64-12                         67590370                18.81 ns/op            0 B/op          0 allocs/op
Benchmark_rand_Perm3-12                         16987875                63.99 ns/op           24 B/op          1 allocs/op
Benchmark_mt_Perm3-12                           17462666                67.02 ns/op           24 B/op          1 allocs/op
Benchmark_rand_Perm30-12                         2985162               457.4 ns/op           240 B/op          1 allocs/op
Benchmark_mt_Perm30-12                           2220142               503.3 ns/op           240 B/op          1 allocs/op
Benchmark_rand_Perm30ViaShuffle-12               4152525               291.3 ns/op             0 B/op          0 allocs/op
Benchmark_mt_Perm30ViaShuffle-12                 4030701               301.6 ns/op             0 B/op          0 allocs/op
Benchmark_rand_ShuffleOverhead-12                3166647               397.6 ns/op             0 B/op          0 allocs/op
Benchmark_mt_ShuffleOverhead-12                  2536288               478.8 ns/op             0 B/op          0 allocs/op
Benchmark_rand_Read3-12                         63220056                19.79 ns/op            0 B/op          0 allocs/op
Benchmark_mt_Read3-12                           52822533                21.12 ns/op            0 B/op          0 allocs/op
Benchmark_rand_Read64-12                        21584400                54.98 ns/op            0 B/op          0 allocs/op
Benchmark_mt_Read64-12                          12575113                98.20 ns/op            0 B/op          0 allocs/op
Benchmark_rand_Read1000-12                       1887732               634.0 ns/op             0 B/op          0 allocs/op
Benchmark_mt_Read1000-12                         1000000              1167 ns/op               0 B/op          0 allocs/op
Benchmark_rand_Int63r1000-12                    73973775                14.93 ns/op            0 B/op          0 allocs/op
Benchmark_mt_Int63r1000-12                      70679886                17.49 ns/op            0 B/op          0 allocs/op
Benchmark_rand_Int63s30-12                       1815048               650.3 ns/op             0 B/op          0 allocs/op
Benchmark_mt_Int63s30-12                         1587296               688.3 ns/op             0 B/op          0 allocs/op
Benchmark_rand_Int63Shuffle30-12                 1915461               608.8 ns/op             0 B/op          0 allocs/op
Benchmark_mt_Int63huffle30-12                    1722516               681.5 ns/op             0 B/op          0 allocs/op
Benchmark_rand_Uint32r1000-12                   79178548                14.99 ns/op            0 B/op          0 allocs/op
Benchmark_mt_Uint32r1000-12                     68490308                17.30 ns/op            0 B/op          0 allocs/op
Benchmark_rand_Uint32s30-12                      2414181               498.2 ns/op             0 B/op          0 allocs/op
Benchmark_mt_Uint32s30-12                        1947705               566.4 ns/op             0 B/op          0 allocs/op
Benchmark_rand_Uint32Shuffle30-12                1924041               621.2 ns/op             0 B/op          0 allocs/op
Benchmark_mt_Uint32Shuffle30-12                  1791526               642.3 ns/op             0 B/op          0 allocs/op
Benchmark_mt_Uint64r1000-12                     70699820                16.60 ns/op            0 B/op          0 allocs/op
Benchmark_rand_Uint64r1000-12                   75709482                14.53 ns/op            0 B/op          0 allocs/op
Benchmark_mt_Uint64s30-12                        1980199               609.0 ns/op             0 B/op          0 allocs/op
Benchmark_rand_Uint64s30-12                      2236620               545.5 ns/op             0 B/op          0 allocs/op
Benchmark_mt_Uint64Shuffle30-12                  1834626               659.2 ns/op             0 B/op          0 allocs/op
Benchmark_rand_Uint64Shuffle30-12                2062753               571.5 ns/op             0 B/op          0 allocs/op
Benchmark_mt_Int31r1000-12                      68594617                16.54 ns/op            0 B/op          0 allocs/op
Benchmark_rand_Int31r1000-12                    80427056                14.71 ns/op            0 B/op          0 allocs/op
Benchmark_rand_Int31s30-12                       2416244               503.4 ns/op             0 B/op          0 allocs/op
Benchmark_mt_Int3130-12                          2053809               596.7 ns/op             0 B/op          0 allocs/op
Benchmark_rand_Int31Shuffle30-12                 1637503               723.9 ns/op             0 B/op          0 allocs/op
Benchmark_mt_Int31Shuffle30-12                   1704332               652.6 ns/op             0 B/op          0 allocs/op
Benchmark_rand_Intr1000-12                      83608130                14.97 ns/op            0 B/op          0 allocs/op
Benchmark_mt_Intr1000-12                        71573498                16.81 ns/op            0 B/op          0 allocs/op
Benchmark_rand_Ints30-12                         1870984               635.2 ns/op             0 B/op          0 allocs/op
Benchmark_mt_Ints30-12                           1742503               697.4 ns/op             0 B/op          0 allocs/op
Benchmark_rand_IntShuffle30-12                   1955840               596.3 ns/op             0 B/op          0 allocs/op
Benchmark_mt_IntShuffle30-12                     1803513               826.2 ns/op             0 B/op          0 allocs/op
Benchmark_rand_Float64r1000-12                  80616586                15.06 ns/op            0 B/op          0 allocs/op
Benchmark_mt_Float64r1000-12                    70612638                17.16 ns/op            0 B/op          0 allocs/op
Benchmark_rand_Float64s30-12                     2433694               464.1 ns/op             0 B/op          0 allocs/op
Benchmark_mt_Float64s0-12                        2272856               527.9 ns/op             0 B/op          0 allocs/op
Benchmark_rand_Float64Shuffle30-12               1960246               614.8 ns/op             0 B/op          0 allocs/op
Benchmark_mt_Float64Shuffle30-12                 1875510               669.9 ns/op             0 B/op          0 allocs/op
Benchmark_rand_Float32r1000-12                  75454182                16.00 ns/op            0 B/op          0 allocs/op
Benchmark_mt_Float32r1000-12                    69269961                17.38 ns/op            0 B/op          0 allocs/op
Benchmark_rand_Float32s30-12                     2427676               497.1 ns/op             0 B/op          0 allocs/op
Benchmark_mt_Float32s0-12                        2178597               571.5 ns/op             0 B/op          0 allocs/op
Benchmark_rand_Float32Shuffle30-12               1970047               620.0 ns/op             0 B/op          0 allocs/op
Benchmark_mt_Float32Shuffle30-12                 1794363               670.5 ns/op             0 B/op          0 allocs/op
Benchmark_rand_FlipCoin-12                      39582832                30.22 ns/op            0 B/op          0 allocs/op
Benchmark_mt_FlipCoin-12                        34671224                34.22 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/bcowtech/gorand      115.529s

```
