# Chapter 5

## Arrays, Slices and Maps

### Arrays

An array is a numbered sequence of elements of a single type with a fixed length. In Go, they can be declared like this:

```go
var x [5]float64
```

`x` is an example of an array which is composed of 5 `float64`s.

In order to fill up each element we can do the following:

```go
x[0] = 98
x[1] = 93
x[2] = 77
x[3] = 82
x[4] = 83
```

There are several ways to iterate over an array:

```go
var total float64 = 0
// 1. hard code the array length (can be problematic, when the array length changes)
for i := 0; i < 5; i++ {
    total += x[i]
}
// 2. using len(x), in order to dynamically get the array length
for i := 0; i < len(x); i++ {
    total += x[i]
}
// 3. using special form of the for loop, but we don't use both variables (that's a no-no for the Go compiler)
for i, value := range x {
    // value is the same as x[i]
    total += value
}
// 4. use underscore to discard one of the variables (simple & elegant)
for _, value := range x {
    total += value
}
// we need type conversion for len, as it's of type int
fmt.Println(total / float64(len(x)))
```

As an alternative to the creation of arrays, Go provides a shorter syntax:

```go
// one line
x := [5]float64{98, 93, 77, 82, 83}
// OR multiple lines
x := [5]float64{
    98,
    93,
    77,
    82,
    83,
}
```

Notice the extra trailing comma after `83`. This is required by Go and it allows us to easily remove an element from the array by commenting out the line:

```go
x := [4]float64{
    98,
    93,
    77,
    82,
    // 83,
}
```

You can also have the compiler count the number of elements for you with `...`:

```go
x := [...]int{1, 2, 3, 4, 5}
```

If you specify the index with `:`, the elements in between will be zeroed:

```go
x := [...]int{1, 3: 4, 5}
// [1 0 0 4 5]
```

These examples also illustrate a major issue with arrays - their length is fixed and part of the array's type name. In order to add/remove one or more items, we must change the type as well. Go's solution to this problem is to use a different type - slice.

### Slices

A slice is a segment of an array. Just like arrays, slices are indexable and have a length. Unlike arrays, however, the length is allowed to change. Here's an example of a slice:

```go
var x []float64
```

Notice that now we don't write the length between the brackets. Unlike arrays, slices are typed only by the elements they contain (not the number of elements). An initialized slice, like in this case, equals to `nil` and has length `0`.

To create a slice with non-zero length, use the built-in `make` function:

```go
x := make([]float64, 5)
```

Here we created a slice of type `float64` of length 5 (initially zero-valued). By default, a new slice's capacity is equal to its length. If we know the slice is going to grow ahead of time, it's possible to pass a capacity explicitly as an additional (third) parameter to `make`.

We can set and get just like with arrays. Additionally, `len()` returns the length of the slice as expected:

```go
s := make([]string, 3)
s[0] = "a"
s[1] = "b"
s[2] = "c"
fmt.Println("set:", s)
fmt.Println("get:", s[2])
```

Another way to create slices is to use the `[low : high]` expression (also known as "slice" operator):

```go
arr := [5]float64{1, 2, 3, 4, 5}
x := arr[0:5]
```

"low" is the index of where to start the slice and "high" is the index of where to end it, but not including the index itself. In the example above, this creates a slice of elements `arr[0]`, `arr[1]`, `arr[2]`, `arr[3]`, and `arr[4]`.

For convenience, it's also allowed to omit low, high, or even both:

- `arr[:5]` is the same as `arr[0:5]`
- `arr[0:]` is the same as `arr[0:len(arr)]`
- `arr[:]` is the same as `arr[0:len(arr)]`

In addition, Go includes two built-in functions to assist with slices: `append` and `copy`

#### Append

`append` adds elements onto the end of a slice. If there is sufficient capacity in the underlying array, the element is placed after the last element and the length is incremented. Otherwise, a new array is created, all of the existing elements are copied over, the new element is added onto the end, and the new slice is returned.

It returns a slice containing one or more new values. Note that we need to accept a return value from `append` as we may get a new slice value.

```go
slice1 := []int{1, 2, 3}
slice2 := append(slice, 4, 5)
fmt.Println(slice1, slice2)
```

Here, `append` creates a new slice by taking an existing slice (as first argument) and appending all the following arguments to it.

#### Copy

`copy` takes two arguments: `dst` (destination) and `src` (source). All of the entries in `src` are copied into `dst` overwriting whatever is there. If the lengths of the two slices are not the same, the smaller of the two will be used.

```go
slice1 := []int{1, 2, 3}
slice2 := make([]int, 2)
copy(slice2, slice1)
fmt.Println(slice1, slice2)
// [1 2 3] [1 2]
```

The contents of `slice1` are copied into `slice2`, but because `slice2` has room for only two elements, only the first two elements of `slice1` are copied.

Slices are typically used to represent lists of items, especially when you need to access the nth item quickly. But what if you want to look up an entry by something other than an integer? What if you need to look up a player on a team by last name? For such cases, Go has another built-in type - map.

### Maps

A *map* is an unordered collection of key-value pairs (aka associative arrays, hash tables, or dictionaries). Maps are used to look up a value by its associated key.

Here is are 2 examples of a map in Go:

```go
// map[KEY_TYPE]VALUE_TYPE
var x map[string]int

y := make(map[string]int)
```



## Problems

### 1. How do you access the 4th element of an array or slice?

Both arrays and slices access their elements in the same manner:

```go
arr := [5]int{1, 2, 3, 4, 5}
s := []int{1, 2, 3, 4, 5}
fmt.Println("The 4th item of arr is", arr[3])
fmt.Println("The 4th item of s is", s[3])
```

### 2. What is the length of a slice created using `make([]int, 3, 9)`?

The length of the slace is 3, whereas its capacity is 9.

> Note: The capacity of the slice is the maximum number of elements that the slice can hold without allocating more memory. It can be checked with the built-in `cap()` function.

### 3. Given the array `x := [6]string{"a","b","c","d","e","f"}` What would `x[2:5]` give you?

`x[2:5]` would return `[c, d, e]`.

### 4. Write a program that finds the smallest number in this list

```go
x := []int{
    48,96,86,68,
    57,82,63,70,
    37,34,83,27,
    19,97, 9,17,
}
```

```go
func task_4() {
    x := []int{
        48, 96, 86, 68,
        57, 82, 63, 70,
        37, 34, 83, 27,
        19, 97, 9, 17,
    }
    // This is the manual approach
    // alternatively, as of go 1.21, we could use slices.Min()
    var smallest_number int
    for index, value := range x {
        if index == 0 || value < smallest_number {
            smallest_number = value
        }
    }
    fmt.Println("The smallest number in x is", smallest_number)
}
```
