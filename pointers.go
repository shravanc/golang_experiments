package main

import ("fmt")

type hello struct{
  name string
}

func my_pointer(h **hello){

  //p := *h
  fmt.Println((*h).name)
  fmt.Println("called")
}


func main(){
  var n *hello
  n = &hello{name: "asd"}
  my_pointer(&n)
  fmt.Println("done")
}
