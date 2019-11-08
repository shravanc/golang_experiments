package main

import ("fmt")

type mystruct struct{
  name string
}

func main(){
  //integer
  var i1 int;
  i2 := 1
  fmt.Println(i2, i1)

  //string
  var s1 string
  s2 := "my string"
  fmt.Println(s2, s1)

  //array of int
  var arr1 []int
  arr2 := []int{1,2,3,}
  fmt.Println(arr2, arr1)

  //struct
  var a mystruct
  a.name = "hello"
  fmt.Println(a)

  //interface
  var i interface{}
  i = 42
  fmt.Println(i)
  i = "now string assignment"
  fmt.Println(i)

  //map 
  var m1 = make(map[string]int)
  m1["key"] = 1
  fmt.Println(m1)

  var m2 = make(map[int]int)
  m2[0] = 1
  fmt.Println(m2)

}
