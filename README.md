# boyermoore

Implementation of Boyer-Moore fast string search algorithm in Go, according to my undestanding of [https://www.cs.utexas.edu/users/moore/best-ideas/string-searching/](https://www.cs.utexas.edu/users/moore/best-ideas/string-searching/). 

If you need to search big chunks of string (or []byte) a lot, you can benefit from this package or Boyer-Moore algorithm.

In my test cases long search terms can be 10 times faster than `strings` package. Looks like strings package is currently using [Rabin Karp](https://en.wikipedia.org/wiki/Rabin%E2%80%93Karp_algorithm) algorithm for longer sub string lookups.

For more details you can check source code it is around 60 lines.


## Usage/Examples

```go
package main

import(
    "fmt"
    "github.com/sarpdag/boyermoore"
) 

func main() {
    src := `Some long string which you need to check. Bla bla ... `
    subStr := `you need to check`
    if pos := boyermoore.Index(src, subStr); pos > -1 {
        fmt.Println("Found in position: ", pos)
    }
}
```

If you need to check the same substring against multiple sources better to pre calculate the sliding table.

```go
package main

import(
    "fmt"
    "github.com/sarpdag/boyermoore"
) 

func main() {
    src := []string{
		`Some long string which you need to check. Bla bla ... `,
		`Another source string ,,faksl fdklssdlkfsjaklsfjklsjfalks`,
	}
	subStr := `you need to check`
	table := boyermoore.CalculateSlideTable(substr)
	if pos := boyermoore.IndexWithTable(&table, src, subStr); pos > -1 {
		fmt.Println("Found in position: ", pos)
	}
}
```

## Benchmarks

```
$ go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/sarpdag/boyermoore
cpu: Intel(R) Core(TM) i9-8950HK CPU @ 2.90GHz

BenchmarkBruteforce/shortEarlySub-12         	 8050245	       135.3 ns/op
BenchmarkBruteforce/shortLateSub-12          	  175002	      6497 ns/op
BenchmarkBruteforce/longSub-12               	  165993	      7233 ns/op
BenchmarkBruteforce/longNotFound-12          	  148197	      7816 ns/op
BenchmarkBruteforce/endOfString-12           	  133054	      8655 ns/op
BenchmarkBruteforce/begOfString-12           	254873073	         4.755 ns/op
BenchmarkBruteforce/shortMid-12              	  270502	      4167 ns/op

BenchmarkStringsIndex/shortEarlySub-12       	19605667	        58.64 ns/op
BenchmarkStringsIndex/shortLateSub-12        	 6066878	       192.3 ns/op
BenchmarkStringsIndex/longSub-12             	  332367	      3323 ns/op
BenchmarkStringsIndex/longNotFound-12        	  542186	      2085 ns/op
BenchmarkStringsIndex/endOfString-12         	  304563	      3913 ns/op
BenchmarkStringsIndex/begOfString-12         	174452732	         6.994 ns/op
BenchmarkStringsIndex/shortMid-12            	  562904	      2065 ns/op

BenchmarkBM/shortEarlySub-12                 	 4530505	       258.3 ns/op
BenchmarkBM/shortLateSub-12                  	  350739	      3269 ns/op
BenchmarkBM/longSub-12                       	 1266339	       935.3 ns/op
BenchmarkBM/longNotFound-12                  	 2445733	       470.6 ns/op
BenchmarkBM/endOfString-12                   	  637333	      1804 ns/op
BenchmarkBM/begOfString-12                   	 6669206	       201.0 ns/op
BenchmarkBM/shortMid-12                      	  874108	      1360 ns/op

BenchmarkBMPregenerated/shortEarlySub-12     	17737701	        67.51 ns/op
BenchmarkBMPregenerated/shortLateSub-12      	  373986	      3179 ns/op
BenchmarkBMPregenerated/longSub-12           	 1729155	       680.2 ns/op
BenchmarkBMPregenerated/longNotFound-12      	 5010907	       234.2 ns/op
BenchmarkBMPregenerated/endOfString-12       	  768459	      1571 ns/op
BenchmarkBMPregenerated/begOfString-12       	181285944	         6.363 ns/op
BenchmarkBMPregenerated/shortMid-12          	  983270	      1214 ns/op
PASS
ok  	github.com/sarpdag/boyermoore	39.736s
```
