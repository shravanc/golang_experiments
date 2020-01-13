package main

import("fmt")

type Person struct{
  Name string
  Age int
}


func (p Person)String() string{
  return fmt.Sprintf("%v age: %v\n", p.Name, p.Age)
}


func main(){
  var p1 = Person{"shravan", 28}
  var p2 = Person{"Pooran", 31}

  fmt.Println(p1, p2)

}
