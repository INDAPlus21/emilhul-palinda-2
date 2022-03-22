# Bug02
## Buggy Code
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go Print(ch)
	for i := 1; i <= 11; i++ {
		ch <- i
	}
	close(ch)
}

func Print(ch <-chan int) {
	for n := range ch {
		time.Sleep(10 * time.Millisecond)
		fmt.Println(n)
	}
}
```
## Problem

The problem is that the main routine ends before the print routine has time to finish printing the last number.

## Fixed Code

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	wait := Print(ch)
	for i := 1; i <= 11; i++ {
		ch <- i
	}
	close(ch)
	<-wait
}

func Print(ch <-chan int) chan struct{} {
	wait := make(chan struct{})
	go func() {
		for n := range ch {
			time.Sleep(10 * time.Millisecond) 
			fmt.Println(n)
		}
		close(wait)
	}()
	return wait
}
```

## Solution
The solution to the problem I used is adding another channel to the program. The print routine now creates a `chan struct{}` called `wait` that stops the main routine from exiting until it's closed after the print routine is done. 