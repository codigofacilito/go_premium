package models

import "fmt"

type User struct{
  Username string
  Password string
}

func CreateUser(username, password string) User{
  user := User{Username: username, Password: password}
  user.EncryptPassword(password)
  return user
}

func (this *User) EncryptPassword(password string){
  this.Password = "password encriptada " + password
}
