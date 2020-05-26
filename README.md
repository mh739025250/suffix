# suffix

This package implements a suffix tree, uses Ukkonen algorithm.

A suffix tree can be built in O(n) operations, find a key in O(k) operations and count a key in O(k) operations.

## example

A simple use case:

```go
package main

import (
"fmt"
"github.com/mh739025250/suffix"
)

func main() {
    content := "abcabxabcd"
    tree := suffix.NewSuffixTree()
    tree.BuildFromStr(content)
    fmt.Println(tree.CountStr("ab"))
    fmt.Println(tree.ExistStr("ac"))
}
```