package main

import(
    "fmt"
    "encoding/json"
    )

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

func push_me(n **Node, x int){
  push(*n, x)
}

func print_list(n *Node){
  for p:=n; p!=nil;p=p.next{
    fmt.Println(p.value)
  }
}



func main(){
  var n1 *Node;
  push_me(&n1, 1)
  push_me(&n1, 2)
  avl,_ := json.MarshalIndent(n1, "", "   ")
  fmt.Println(string(avl))

}
