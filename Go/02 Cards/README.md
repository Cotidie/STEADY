# Cards
From the deck of 52 cards, the program provides those functions: create a deck, print, shuffle, deal, save and load deck. 'deal' command creates a 'hand' of cards.
- **deal(deck, int)**: pick n number of cards from a deck
  - **return**: a hand of cards and a remaining deck
- **shuffle()**: randomize the order of the current deck
  - Go doesn't shuffle a slice
- **saveToFile()**: saves a list of cards to a file on the local machine
  - it should use JSON format
- **readFromFile()**: reads JSON file into deck type.



## Standard Packages
### io/ioutil
| https://pkg.go.dev/io/ioutil
- [WriteFile(filename string, data []byte, perm FileMode))](https://pkg.go.dev/io/ioutil#WriteFile)
  - Filemode follows Linux permission form, which is 3 bits(rwx) for 3 user groups (0666 for read, write permissions for all three)
  - Save path follows current working directory
- [ReadFile(filename string)](https://pkg.go.dev/io/ioutil#ReadFile)
  - if error occurs, error != nil

### strings
- [Join(a []string,sep string)](https://pkg.go.dev/strings#Join)
- [Split(s, sep string)](https://pkg.go.dev/strings#Split)

### encoding/json
- [Marhsal(v any)](https://pkg.go.dev/encoding/json#Marshal): Converts any type that implements MarshalJSON or TextMarshaler into JSON bytes. use [MarshalIndent(...)](https://pkg.go.dev/encoding/json#MarshalIndent) if needs prettier form.

### math/rand
: Before call any random functions, the seed must be changed. For this purpose, use NewSource() and New(source Source)
- [NewSource(seed int64)](https://pkg.go.dev/math/rand#NewSource): create a new seed. **time.Now().UnitNano()** is recommended.
- [New(source Source)](https://pkg.go.dev/math/rand#New)
- [Intn(n int)](https://pkg.go.dev/math/rand#Intn): returns a random int between [0, n)

### os
- [Exit(code int)](https://pkg.go.dev/os#Exit): Completely finishes the program with exit code