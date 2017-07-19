package main

import(
  "os"
  "fmt"
)
func main(){
  size := len(os.Args)
  fmt.Println(size)
  
  if len(os.Args) == 1{
    fmt.Println("Es necesario colocar más de un argumento!")
  }else{
    argumentos := os.Args[1:]
    for _, arg := range argumentos{
      fmt.Println(arg)
    }
  }
}