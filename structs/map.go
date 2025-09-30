package main

import "fmt"

func main(){
  m := make(map[string]int)

  m["sam"]=40
  m["elon"]=54

  m["jeff"]=60

  fmt.Println(m)

  m["jeff"]=61
  fmt.Println("After updating  map : ", m)

  fmt.Println("sam's age : " , m["sam"])

  age, ok := m["sundar"]
    if ok {
        fmt.Println("sundar's age:", age)
    } else {
        fmt.Println("sundar not found")
    }

  delete(m, "jeff")
  fmt.Println("After deleting map", m)

  fmt.Println("All entries:")
    for name, age := range m {
        fmt.Println(name, "=>", age)
    }

  
