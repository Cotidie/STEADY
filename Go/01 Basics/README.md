# Go Basics

- [Package](#package)
- [Import](#import)
- [Variables](#variables)
- [Functions](#functions)
- [Slice and Loop](#slice-and-loop)

## Basics
```go
package main

// use linebreaks to import multiple packages
import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println("Hello, world!")
}
```
### Package
 Package refers to a workspace or a project. 
* **Executable**: Generate a .exe file we can run, by 'go build command'
    * package name must be **'main'**
    * must have 'main' function
* **Reusable**: Packages used as dependencies.
    * anything but name 'main' is a reusable package.
  
| Q. How to test reusable packages

### Import
| Standard Packages: https://pkg.go.dev/std

import standard or reusable packages. The root path is where the main package is located at.

| Q. absolute and relative import?

### Variables
Go is referred to a static typed language. The variable name is followed by its type declaration.  
  * **baisc types**: bool, string, int, float64, array, map
  * types can be inferred from its declaration
  * global variables can be declared, but cannot be initialized
  
```go
var name type = value
card := value  // type inferrence

// can be assigned later
var deckSize int
deckSize = 52
```

### Functions
- Arguments are basically passed by value
- Functions can have multiple returns
```go
func main() {
    card := newCard()
    fmt.Println(card)
}

// function with multiple returns
func deal(d deck, size int) (deck, deck) {
    return deck[:size], deck[size:]
}
```

### Slice and Loop
Go provides slices as a default array type.
* **Slice**: an array can grow and shrink (ref. vector in c++)
* ```range``` keyword gets an index and an iterator from an iterable
```go
func main() {
    cards := []type{"a", "b", "c"}
    cards = append(cards, "d")

    // if index is not in use, replace it with '_'
    for i, card := range cards {
        fmt.Println(i, card)
    }

    // slice: [start:end)
    sliced := cards[0:2] // = cards[:2]
}
```

### Class
```go
// Declare a new type deck
type deck []string

// Receiver: d is an instance of the type deck
func (d deck) print() {
    for i, card := range deck {
        fmt.PrintlnI(i, card)
    }
}
```


## Errors
### Unnamed Name
* Cause: a go file in the same package wasn't imported
* Solution: initialize go modules first
```go
go mod init
go mod tidy
```

