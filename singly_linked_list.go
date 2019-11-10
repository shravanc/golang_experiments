package main

import("fmt")


type node struct{
  value int
  next *node
}

func push(n *node, x int){

  for{
    if n.next == nil{ 
      new_node := node{value: x}
      n.next = &new_node
      break
    }else{
      n = n.next
    }
  }


}

func print_list(n *node){
  for p:=n; p!=nil;p=p.next{
    fmt.Println(p.value)
  }
}



func main(){
  var n1 node;

  push(&n1, 1)
  push(&n1, 2)

  print_list(&n1)

}
