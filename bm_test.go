package boyermoore

import (
	"strings"
	"testing"
)

func IndexBruteforceByte(s, substr string) int {
	lensub := len(substr)
	lens := len(s)
	for i := 0; i < lens-lensub+1; i++ {
		found := true
		for j := 0; j < lensub; j++ {
			if s[i+j] == substr[j] {
				continue
			}
			found = false
			break
		}
		if found {
			return i
		}
	}
	return -1
}

var table = []struct {
	name   string
	substr string
}{
	{"shortEarlySub", shortEarlySub},
	{"shortLateSub", shortLateSub},
	{"longSub", longSub},
	{"longNotFound", longNotFound},
	{"endOfString", endOfString},
	{"begOfString", begOfString},
	{"shortMid", shortMid},
}

func TestIndex(t *testing.T) {
	testCases(t, Index)
}

func TestIndexBruteforceByte(t *testing.T) {
	testCases(t, IndexBruteforceByte)
}
func TestStringsIndex(t *testing.T) {
	testCases(t, strings.Index)
}

func testCases(t *testing.T, f func(string, string) int) {
	for _, tc := range table {
		expected := strings.Index(longText, tc.substr)
		calculated := f(longText, tc.substr)
		if calculated != expected {
			t.Errorf("%q Expected %d calculated %d", tc.substr, expected, calculated)
		}
	}
}

func BenchmarkBruteforce(b *testing.B) {
	for _, tc := range table {
		b.Run(tc.name, func(bs *testing.B) {
			for i := 0; i < bs.N; i++ {
				IndexBruteforceByte(longText, tc.substr)
			}
		})
	}
}

func BenchmarkStringsIndex(b *testing.B) {
	for _, tc := range table {
		b.Run(tc.name, func(bs *testing.B) {
			for i := 0; i < bs.N; i++ {
				strings.Index(longText, tc.substr)
			}
		})
	}
}

func BenchmarkBM(b *testing.B) {
	for _, tc := range table {
		b.Run(tc.name, func(bs *testing.B) {
			for i := 0; i < bs.N; i++ {
				Index(longText, tc.substr)
			}
		})
	}
}

func BenchmarkBMPregenerated(b *testing.B) {
	for _, tc := range table {
		b.Run(tc.name, func(bs *testing.B) {
			d := CalculateSlideTable(tc.substr)
			for i := 0; i < bs.N; i++ {
				IndexWithTable(&d, longText, tc.substr)
			}
		})
	}
}

const (
	endOfString   = "standard API)."
	begOfString   = "Package"
	shortMid      = "shared" // first occurrence somewhere in the middle
	shortEarlySub = "the"
	shortLateSub  = "API"
	longSub       = "the designers explicitly argue against assertions and pointer arithmetic, while defending the choice to omit type"
	longNotFound  = "lkafkas3 j3fkjf3kljs; dlksajfsdaojsf3lkajd klvsj3 sakkdjfakdlsjfkasj ksadj klajsdsfklasjfk"

	// source: https://en.wikipedia.org/wiki/Go_(programming_language)
	longText = `Package
In Go's package system, each package has a path (e.g., "compress/bzip2" or "golang.org/x/net/html") and a name (e.g., bzip2 or html). References to other packages' definitions must always be prefixed with the other package's name, and only the capitalized names from other packages are accessible: io.Reader is public but bzip2.reader is not.[74] The go get command can retrieve packages stored in a remote repository [75] and developers are encouraged to develop packages inside a base path corresponding to a source repository (such as example.com/user_name/package_name) to reduce the likelihood of name collision with future additions to the standard library or other external libraries.[76]
Proposals exist to introduce a proper package management solution for Go similar to CPAN for Perl or Rust's cargo system or Node's npm system.[77
Concurrency: goroutines and channels[edit]
The Go language has built-in facilities, as well as library support, for writing concurrent programs. Concurrency refers not only to CPU parallelism, but also to asynchrony: letting slow operations like a database or network read run while the program does other work, as is common in event-based servers.[78]
The primary concurrency construct is the goroutine, a type of light-weight process. A function call prefixed with the go keyword starts a function in a new goroutine. The language specification does not specify how goroutines should be implemented, but current implementations multiplex a Go process's goroutines onto a smaller set of operating-system threads, similar to the scheduling performed in Erlang.[79]:10
While a standard library package featuring most of the classical concurrency control structures (mutex locks, etc.) is available,[79]:151–152 idiomatic concurrent programs instead prefer channels, which provide send messages between goroutines.[80] Optional buffers store messages in FIFO order[62]:43 and allow sending goroutines to proceed before their messages are received.[citation needed]

Channels are typed, so that a channel of type chan T can only be used to transfer messages of type T. Special syntax is used to operate on them; <-ch is an expression that causes the executing goroutine to block until a value comes in over the channel ch, while ch <- x sends the value x (possibly blocking until another goroutine receives the value). The built-in switch-like select statement can be used to implement non-blocking communication on multiple channels; see below for an example. Go has a memory model describing how goroutines must use channels or other operations to safely share data.[81]
The existence of channels sets Go apart from actor model-style concurrent languages like Erlang, where messages are addressed directly to actors (corresponding to goroutines). The actor style can be simulated in Go by maintaining a one-to-one correspondence between goroutines and channels, but the language allows multiple goroutines to share a channel or a single goroutine to send and receive on multiple channels.[79]:147

From these tools one can build concurrent constructs like worker pools, pipelines (in which, say, a file is decompressed and parsed as it downloads), background calls with timeout, "fan-out" parallel calls to a set of services, and others.[82] Channels have also found uses further from the usual notion of interprocess communication, like serving as a concurrency-safe list of recycled buffers,[83] implementing coroutines (which helped inspire the name goroutine),[84] and implementing iterators.[85]

Concurrency-related structural conventions of Go (channels and alternative channel inputs) are derived from Tony Hoare's communicating sequential processes model. Unlike previous concurrent programming languages such as Occam or Limbo (a language on which Go co-designer Rob Pike worked),[86] Go does not provide any built-in notion of safe or verifiable concurrency.[87] While the communicating-processes model is favored in Go, it is not the only one: all goroutines in a program share a single address space. This means that mutable objects and pointers can be shared between goroutines; see § Lack of race condition safety, below.[citation needed]

Suitability for parallel programming[edit]
Although Go's concurrency features are not aimed primarily at parallel processing,[78] they can be used to program shared-memory multi-processor machines. Various studies have been done into the effectiveness of this approach.[88] One of these studies compared the size (in lines of code) and speed of programs written by a seasoned programmer not familiar with the language and corrections to these programs by a Go expert (from Google's development team), doing the same for Chapel, Cilk and Intel TBB. The study found that the non-expert tended to write divide-and-conquer algorithms with one go statement per recursion, while the expert wrote distribute-work-synchronize programs using one goroutine per processor. The expert's programs were usually faster, but also longer.[89]

Lack of race condition safety[edit]
There are no restrictions on how goroutines access shared data, making race conditions possible. Specifically, unless a program explicitly synchronizes via channels or other means, writes from one goroutine might be partly, entirely, or not at all visible to another, often with no guarantees about ordering of writes.[87] Furthermore, Go's internal data structures like interface values, slice headers, hash tables, and string headers are not immune to race conditions, so type and memory safety can be violated in multithreaded programs that modify shared instances of those types without synchronization.[90][91] Instead of language support, safe concurrent programming thus relies on conventions; for example, Chisnall recommends an idiom called "aliases xor mutable", meaning that passing a mutable value (or pointer) over a channel signals a transfer of ownership over the value to its receiver.[79]:155

Binaries[edit]
The linker in the gc toolchain creates statically linked binaries by default, therefore all Go binaries include the Go runtime.[92][93]
Omissions[edit]
Go deliberately omits certain features common in other languages, including (implementation) inheritance, generic programming, assertions,[e] pointer arithmetic,[d] implicit type conversions, untagged unions,[f] and tagged unions.[g] The designers added only those facilities that all three agreed on.[96]

Of the omitted language features, the designers explicitly argue against assertions and pointer arithmetic, while defending the choice to omit type inheritance as giving a more useful language, encouraging instead the use of interfaces to achieve dynamic dispatch[h] and composition to reuse code. Composition and delegation are in fact largely automated by struct embedding; according to researchers Schmager et al., this feature "has many of the drawbacks of inheritance: it affects the public interface of objects, it is not fine-grained (i.e, no method-level control over embedding), methods of embedded objects cannot be hidden, and it is static", making it "not obvious" whether programmers will overuse it to the extent that programmers in other languages are reputed to overuse inheritance.[61]

The designers express an openness to generic programming and note that built-in functions are in fact type-generic, but these are treated as special cases; Pike calls this a weakness that may at some point be changed.[53] The Google team built at least one compiler for an experimental Go dialect with generics, but did not release it.[97] They are also open to standardizing ways to apply code generation.[98] In June 2020, a new draft design document[99] was published, which would add the necessary syntax to Go for declaring generic functions and types. A code translation tool go2go was provided to allow users to try out the new syntax, along with a generics-enabled version of the online Go Playground.[100]
Initially omitted, the exception-like panic/recover mechanism was eventually added, which the Go authors advise using for unrecoverable errors such as those that should halt an entire program or server request, or as a shortcut to propagate errors up the stack within a package (but not across package boundaries; there, error returns are the standard API).`
)
