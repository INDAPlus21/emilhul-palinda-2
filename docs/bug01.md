# Bug01
## Buggy Code

```go
package main

import "fmt"

// I want this program to print "Hello world!", but it doesn't work.
func main() {
	ch := make(chan string)
	ch <- "Hello world!"
	fmt.Println(<-ch)
}
```
## Problem

The problem is that the program reaches a deadlock. In other words a state where all parts of the program are waiting on others to complete. The program stops at `ch <- "Hello world!"` and waits for someone to read the input. But since the program is waiting it will never reach the next line where it would have read the input.

## Fixed Code

```go
package main

import "fmt"

func main() {
	ch := make(chan string)
	go helperFunc(ch)
	fmt.Println(<-ch)
}

func helperFunc(ch chan<- string) {
	ch <- "Hello world!"
}
```
## Solution

The solution to the problem is to send and recieve on diffrent threads. I decieded to do this by moving the sender into it's own function. The new function can then be called with as a goroutine meaning it runs on a seperate thread. This allows the main thread to continue to the print statment.

(Only use case for channels to begin with are multi-threaded programs, if we wanted it to be single-threaded we could have just printed "Hello world! or stored it in a variable and printed that variable) 