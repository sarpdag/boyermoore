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
Bruteforce/shortEarlySub-12       138ns ± 2%
Bruteforce/shortLateSub-12       6.77µs ± 2%
Bruteforce/longSub-12            7.11µs ± 9%
Bruteforce/longNotFound-12       8.89µs ± 8%
Bruteforce/endOfString-12        9.04µs ± 3%
Bruteforce/begOfString-12        4.77ns ± 3%
Bruteforce/shortMid-12           4.41µs ± 2%
StringsIndex/shortEarlySub-12    63.4ns ± 3%
StringsIndex/shortLateSub-12      214ns ± 2%
StringsIndex/longSub-12          3.79µs ± 4%
StringsIndex/longNotFound-12     2.27µs ± 1%
StringsIndex/endOfString-12      4.05µs ± 2%
StringsIndex/begOfString-12      7.13ns ± 3%
StringsIndex/shortMid-12         2.02µs ± 3%
BM/shortEarlySub-12               374ns ± 7%
BM/shortLateSub-12               9.51µs ± 1%
BM/longSub-12                    1.00µs ± 9%
BM/longNotFound-12                723ns ± 7%
BM/endOfString-12                3.07µs ± 8%
BM/begOfString-12                 199ns ± 6%
BM/shortMid-12                   3.29µs ± 4%
BMPregenerated/shortEarlySub-12   186ns ± 2%
BMPregenerated/shortLateSub-12   9.54µs ± 5%
BMPregenerated/longSub-12         744ns ± 2%
BMPregenerated/longNotFound-12    444ns ± 2%
BMPregenerated/endOfString-12    2.53µs ± 2%
BMPregenerated/begOfString-12    6.28ns ± 3%
BMPregenerated/shortMid-12       2.90µs ± 3%
```
