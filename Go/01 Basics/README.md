# Go Basics
## Basics
```go
package main

import "fmt"

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
import standard of reusable packages. The root path is where the main package is located at.

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
```go
func main() {
    card := newCard()
    fmt.Println(card)
}

// the return type follows the function name
func newCard() string {
    return "Five of Diamonds"
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

    for i, card := range cards {
        fmt.Println(i, card)
    }
}
```


