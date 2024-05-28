# lastrune
## Instructions
Write a function that returns the last rune of a string.

## Expected function
```bash
func LastRune(s string) rune {

}
```
## Usage
Here is a possible program to test your function :
```bash
package main

import (
	"github.com/01-edu/z01"

	"piscine"
)

func main() {
	z01.PrintRune(piscine.LastRune("Hello!"))
	z01.PrintRune(piscine.LastRune("Salut!"))
	z01.PrintRune(piscine.LastRune("Ola!"))
	z01.PrintRune('\n')
}
```
And its output :
```bash
$ go run .
!!!
$
```