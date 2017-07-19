package main

import(
  "flag"
  "fmt"
)

var puerto int
var debug bool
var ip string

func init(){
  flag.IntVar(&puerto, "puerto", 5000, "Establece el puerto!")
  flag.BoolVar(&debug, "debug", false, "Establece el modo debug!")
  flag.StringVar(&ip, "ip", "127.0.0.1", "Establece la Ip!")
  flag.Parse()
}

func main(){
  fmt.Println(puerto)
  fmt.Println(debug)
  fmt.Println(ip)

  argumentos := flag.Args()
  fmt.Println(argumentos)
}

//go run 7.-banderas.go --help