package main

import "fmt"

func main() {
    
  var users = make(map[int]string)
  users[0] = "eduardo_gpg"
  user := users[0]
  fmt.Println(user)

  username, ok := users[0]
  if ok{
    fmt.Println("el usuario existe", username)
  }else{
    fmt.Println("El usuario no existe")
  }

  courses := map[string]string{ "gos": "go web",
    "py": "Python 3",
  }
  
  fmt.Println(courses)
  if el, ok := courses["py"]; ok {
    fmt.Println(el)
  }

}

