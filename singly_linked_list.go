package main

import("fmt")

type Node struct{
  value int
  next *Node
}

func push(n *Node, x int){

  for{
    if n.next == nil{
      new_node := Node{value: x}
      n.next = &new_node
      break
    }else{
      n = n.next
    }
  }


}

func print_list(n *Node){
  for p:=n; p!=nil;p=p.next{
    fmt.Println(p.value)
  }
}



func main(){
  var n1 Node;
  push(&n1, 1)
  push(&n1, 2)
  print_list(&n1)

}
