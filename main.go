// Go requires every variable you declare to be used somewhere in your code. If you comment out all uses of a variable, you'll get an error until you use it again or remove the declaration.

package main

import "fmt"

func add(x int, y int) int{
  return x+y
  }

func main(){
  sum := add(5,3)
  fmt.Println(sum)
  
  var score int = 98;
  if score > 50 {
		fmt.Println("Passed!")
	} else {
		fmt.Println("Try again.")
	}

  }
