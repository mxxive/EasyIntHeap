# EasyIntHeap
Easy to use integer heap for Go

# Example
```go
package main

import (
	"fmt"

	"github.com/mxxive/intheap"
)

func main() {
	heap := intheap.NewIntHeap()
	heap.Push(3)
	heap.Push(4)
	heap.Push(1)
	heap.Push(2)
	for i := 0; i < 4; i++ {
		fmt.Print(heap.Pop(), " ")
	}
}
```
## Result
```
1 2 3 4 
```
