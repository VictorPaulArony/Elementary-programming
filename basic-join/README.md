# basicjoin
## Instructions
Write a function that returns a concatenated string from the 'strings' passed as arguments.

## Expected function
```go
func BasicJoin(elems []string) string {

}
```

## Usage
Here is a possible program to test your function :
```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	elems := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(piscine.BasicJoin(elems))
}
```
```bash
And its output :

$ go run .
Hello! How are you?
$
```