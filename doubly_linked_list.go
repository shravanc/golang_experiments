package main

import("fmt")

type node struct{
  value int
  next *node
  previous *node
}


func push(n *node, i int){
  for{
    if n.next == nil{
      new_node := node{value: i, previous: n}
      n.next = &new_node
      break
    }else{
      n = n.next
    }
  }
}

func print_list(n *node){
  for n1:=n; n1!=nil;n1=n1.next{
    fmt.Println(n1.value)
  }
}

func main(){
  var n1 = node{value: 1}

  push(&n1, 44)
  push(&n1, 4)

  fmt.Println(n1.next.value)
  fmt.Println(n1.next.next.previous.value)

  print_list(&n1)
}
