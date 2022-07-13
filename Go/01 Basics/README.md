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



