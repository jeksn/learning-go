# Maps

In Go, a map is a built-in data structure that allows you to store key-value pairs. It is similar to a dictionary or hash table in other programming languages. 

To declare a map variable, you use the `map` keyword followed by the types of the keys and values in square brackets. For example:

```go
var myMap map[string]int
```

This declares a map with keys of type `string` and values of type `int`.

To initialize a map, you can use the `make` function, like so:

```go
myMap := make(map[string]int)
```

To add or update a value in a map, you can simply assign a value to a key:

```go
myMap["foo"] = 42
```

To retrieve a value from a map, you can use the key in square brackets:

```go
fmt.Println(myMap["foo"]) // prints 42
```

If a key is not present in the map, the value returned will be the zero value for the value type (in this case, 0 for `int`).

You can also check if a key is present in a map using the following syntax:

```go
value, ok := myMap["foo"]
if ok {
    fmt.Println(value)
} else {
    fmt.Println("Key not found")
}
```

This assigns the value associated with the key to the `value` variable, and sets `ok` to `true` if the key is present in the map, or `false` otherwise.