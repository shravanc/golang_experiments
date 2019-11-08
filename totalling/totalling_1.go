package main

import ("fmt")

func main(){
  max := 500
  min := 0
  sum := 0

  for i:=min; i<=max; i++{
    sum += i
  }
  fmt.Println(sum)
}

