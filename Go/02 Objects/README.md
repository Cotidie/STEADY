# Objects

## Struct
* if no value is assigned, the field will be initiated with its zero value
* plus(+) sign in %+v shows field names with values.
* embedding a struct into another is possible as a syntatic sugar
* Use a pointer to pass by reference
```go
// Defining a struct
type person struct {
    name string
    contact            // = contact contact
}
type contact struct {
    phone string
    email string
}

// Declaring a struct
me := {
    name: "Wonseok",
    contact: contact{
        phone: "010-4797-1428",
        email: "daily142857@gmail.com"
    }
}
me.contact.phone = "669-338-6638"

// normal form
mePointer := &me
mePointer.updateName("Finn")
// pointer shortcut
me.updateName("Finn")

// to actually change the value, receiver should be a pointer
func (p *person) updateName(newName string) {
    (*person).name = newName
}
```
## Map
```go
// Initializing a map
color := map[string][int] {
    "red": 1,
    "green": 2,
    "blue": 3,
}
// Just declaring a map
colors := make(map[string]string)
var colors map[string][string]

// Write and Read
colors["red"] = 1
colors["red"]       // = colors.red

// Delete an entry
delete(colors, "red")

// Iterating: replace with _ if not want to use that variable
for key, value := range colors {
    // key, value
}
```
## Interface