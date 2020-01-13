package main

import ("fmt")

type I interface{
  M()
}

type Shr struct{
  x int
}

func (s *Shr)M(){
  if s ==nil{
    fmt.Println("Sorry nothing is there in this address")
    return
  }
  fmt.Println(s.x)
}


func main(){
  var s1 *Shr
  var i I

  i = s1
  i.M()

  s2 := &Shr{4}
  i = s2
  i.M()

}
