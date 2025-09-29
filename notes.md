1. If you want to print without a newline, use fmt.Print().
2. If you want to format output (like with variables), use fmt.Printf().
3. Go requires every variable you declare to be used somewhere in your code. If you comment out all uses of a variable, you'll get an error until you use it again or remove the declaration.
4. To declare a variable in if initializer, we use n:= 10; like this **TYPE INFERENCE**
5. We can use blank identifier _ at many places to make it a garbage value like in for loop if we did not want to decalare index we make it like for _,

 ```For _, v := range nums {
    fmt.Println(v) // Only use value, ignore index
   }```
