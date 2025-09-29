# Go Slices: Quick Notes

## What is a Slice?
- A **slice** is a dynamic, flexible view into the elements of an array.
- Slices are more common than arrays in Go for most tasks.

## Declaring and Initializing Slices
```go
var nums []int // Declares an empty slice of ints
nums := []int{1, 2, 3, 4} // Initializes a slice with values
```

## Accessing Elements
```go
fmt.Println(nums[0]) // Access first element (index 0)
```

## Appending Elements
```go
nums = append(nums, 5) // Adds 5 to the end
```

## Getting Length
```go
fmt.Println(len(nums)) // Number of elements in the slice
```

## Slicing a Slice
```go
sub := nums[1:3] // Elements at index 1 and 2 (end is exclusive)
```

## Removing an Element
```go
// Remove element at index 2
nums = append(nums[:2], nums[3:]...)
```

## Iterating Over a Slice
```go
for i, v := range nums {
    fmt.Println("Index:", i, "Value:", v)
}
```

## Copying a Slice
```go
src := []int{1, 2, 3}
dst := make([]int, len(src))
copy(dst, src)
```

## Key Points
- Slices are **zero-indexed**.
- Slices can grow and shrink with `append` and slicing.
- Use `len(slice)` for length.
- Use `range` for easy iteration.
- Removing elements requires combining slicing and `append`.
- Slices are references to underlying arrays—modifying one can affect others if they share the same backing array.

***
*Copy and use these notes as a quick reference for Go slices and their operations!*
