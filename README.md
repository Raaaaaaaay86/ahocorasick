# ahocorasick

A Go implementation of the Aho-Corasick multiple pattern string search algorithm.

## Features
- Supports searching for multiple keywords simultaneously
- Trie structure with automatic fail link construction
- Find all matches or the longest match
- Unicode support

## Installation

```
go get github.com/raaaaaaaay86/ahocorasick
```

## Usage Example

```go
package main

import (
	"fmt"
	"github.com/raaaaaaaay86/ahocorasick"
)

func main() {
	dictionary := []string{"hello", "world", "hell", "helloworld", "hellow"}
	trie := ahocorasick.NewTrie(dictionary)
	matches := trie.FindAllMatches("helloworld")
	for _, m := range matches {
		fmt.Println(m)
	}
}
```

## Testing

```
go test -v
```

## File Structure
- `trie.go`: Main trie structure and construction logic
- `node.go`: Trie node and matching logic
- `trie_test.go`: Unit tests

## Reference
- [Aho-Corasick algorithm - Wikipedia](https://en.wikipedia.org/wiki/Aho%E2%80%93Corasick_algorithm)
