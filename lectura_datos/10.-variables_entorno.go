package main

import(
  "os"
  "fmt"
)

func getEnvOrDefault(variableName, defaultValor string) string{
  if val := os.Getenv(variableName); val == ""{
    return defaultValor
  }else{
    return val
  }
}

func main(){
  entorno := getEnvOrDefault("NombreVariable", "no existe :(")
  os.Setenv("NombreVariable", "Valor")
  os.Unsetenv("NombreVariable")

  fmt.Println(entorno)
}