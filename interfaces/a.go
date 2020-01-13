package main

import("fmt")

type I interface{
  sample()
}

type shr_Struct struct{
  x int
}

type poo_struct struct{
  x int
}

func (ms shr_Struct) sample(){
  fmt.Println(ms.x)
}

func (ps poo_struct) sample(){
  fmt.Println(ps.x *2)
}


func mediator(i I){
  i.sample()
}

func main(){

  var s1 = shr_Struct{1}
  var p1 = poo_struct{2}

  var ifc I
  fmt.Println(ifc == nil)

  mediator(s1)
  mediator(p1)
}
