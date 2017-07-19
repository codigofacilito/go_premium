package main

import(
  "flag"
  "fmt"
)

func main(){
  var puerto = flag.Int("puerto", 5000, "Establece el puerto!")
  var debug = flag.Bool("debug", false, "Establece el modo debug!")

  flag.Parse()
  fmt.Println(*puerto)
  fmt.Println(*debug)
}

//go run 7.-banderas.go --help