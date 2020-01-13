package main

import("fmt")






func main(){
  var i interface{} = "hello"

  _, ok := i.(string)
  fmt.Println(ok)


  _, ok= i.(int)
  fmt.Println(ok)


  _, ok = i.(float64)
  fmt.Println(ok)
}
