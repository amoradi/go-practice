# NOTES

## WORKSPACES: TYPICAL SETUP
The go tool builds source packages and installs the resulting binaries to the pkg and bin directories.

- *src* contains Go source files,
- *pkg* contains package objects
- *bin* contains executable commands.

## PACKAGES
The first statement in a Go source file must be
```go
package name
```
where name is the package's default name for imports. (All files in a package must use the same name.

- Executable commands must always use package main.
- There is no requirement that package names be unique across all packages linked into a single binary, only that the import paths (their full file names) be unique.
- If you include the repository URL in the package's import path, go get will fetch, build, and install it automatically:
- executable packages are named 'main'
- reusable packages are not name 'main'

## TYPES
- string
- bool
- float64
- int
- array: fixed len, must be the same type
- slice: not fixed len, must be the same type

## VARIABLES
```go
// equivalent ways to declare and assign var's
var card string = "Ace of spades"
card := "Ace of spades"
```

## FUNCTIONS
```go
func fooBar() string {
    return "Shaken, but not stirred"
}
```

## FOR LOOP
```go
    cards := []string{"One", "Two", "Three"}
    cards = append(cards, "Four") // returns a NEW slice

    for i, card := range cards {
        fmt.Println(i, card)
    }
```

## OOP VS TYPES
- custom types
- func's with custom types as RECIEVER's

```go
// custom type
type deck []string

// reciever example
func (d deck) print() {
    for i, card := range d {
        fmt.Println(i, card)
    }
}
```

## POINTERS
Use pointers with non-reference types (these are copied when passes as func args, and thus, their original address in memory needs pointing to) :
- int
- float
- string
- bool
- struct

Reference Types include:
- slice
- map
- channel
- pointer
- function

## ERROR HANDLING

## STRUCTS
- non-reference type, i.e. a value type
-- these are copied when passed as args
-- need to point to original address in memory
- keys don't support indexing
- values can vary in type

## MAPS
- reference type
- all keys must be the same type
- all values must be the same type
- all keys are indexed
- maps are great for closely-related properties
- can add/delete members

```go
    // ways to declare and assign maps
    // 1
    colors := map[string]string{
		"red": "#FF0000",
		"green": "#4bf745",
	}

    // 2
	var colors map[string]string

    // 3
	colors := make(map[string]string)

    // maps only use bracket[syntax]
    // cannot access members with dot.syntax
    colors["white"] = "#ffffff"
    
    // delete members
	delete(colors, "white")
```

## INTERFACES

## CHANNELS

## ROUTINES

