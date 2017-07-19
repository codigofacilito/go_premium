package main

import "fmt"
import "models"

func main(){
  user := models.CreateUser("Eduardo", "Password123")
  fmt.Println(user)

}