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
$ go test -bench=. -count=5 > bench.txt
$ benchstat bench.txt

name                             time/op
Bruteforce/shortEarlySub-12       140ns ± 3%
Bruteforce/shortLateSub-12       7.16µs ± 5%
Bruteforce/longSub-12            7.38µs ± 3%
Bruteforce/longNotFound-12       8.06µs ± 5%
Bruteforce/endOfString-12        8.82µs ± 3%
Bruteforce/begOfString-12        5.00ns ± 3%
Bruteforce/shortMid-12           4.24µs ± 3%
StringsIndex/shortEarlySub-12    61.9ns ± 4%
StringsIndex/shortLateSub-12      198ns ± 1%
StringsIndex/longSub-12          3.61µs ± 4%
StringsIndex/longNotFound-12     2.11µs ± 1%
StringsIndex/endOfString-12      3.78µs ± 3%
StringsIndex/begOfString-12      6.49ns ± 1%
StringsIndex/shortMid-12         1.93µs ± 1%
BM/shortEarlySub-12               280ns ±25%
BM/shortLateSub-12               3.31µs ± 2%
BM/longSub-12                    1.08µs ± 6%
BM/longNotFound-12                518ns ± 7%
BM/endOfString-12                1.85µs ± 6%
BM/begOfString-12                 200ns ± 9%
BM/shortMid-12                   1.38µs ± 2%
BMPregenerated/shortEarlySub-12  66.4ns ± 2%
BMPregenerated/shortLateSub-12   3.18µs ± 6%
BMPregenerated/longSub-12         718ns ± 8%
BMPregenerated/longNotFound-12    236ns ± 5%
BMPregenerated/endOfString-12    1.55µs ± 4%
BMPregenerated/begOfString-12    6.52ns ± 6%
BMPregenerated/shortMid-12       1.12µs ± 5%
```
