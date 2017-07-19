package main

import "fmt"

func reverse(slice []string){
  for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
    slice[i], slice[j] = slice[j], slice[i]
  }
  fmt.Println(slice)
}

func main() {
  
  months := []string {"Enero","Febrero","Marzo","Abril","Mayo","Junio","Julio","Agosto","Septiembre","Octubre","Noviembre","Diciembre"}
  months = append(months, "eduardo", "ismael", "n cantidad de strings")
  reverse(months)
}

