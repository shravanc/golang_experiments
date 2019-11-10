package main

import("fmt")


type Node struct{
  data int
  next *Node
  previous *Node
}

func producer(node *Node, data int){

  for {
    if node.next == nil{
      new_node := Node{data: data, previous: node}
      node.next = &new_node
      break
    }else{
      node = node.next
    }
  }


}

func consumer(node *Node){
  for {
    if node.next == nil{
      prev_node := node.previous
      prev_node.next = nil
      fmt.Printf("Consumed %d\n", node.data)
      break
    }else{
      node = node.next
    }
  }
}

func print_list(node *Node){
  for n:=node; n!=nil; n=n.next{
    fmt.Println(n.data)
  }
}


func main(){
  var node = Node{data: 1}
  defer print_list(&node)

  producer(&node, 2)
  producer(&node, 3)
  consumer(&node)
  producer(&node, 4)
  producer(&node, 5)
  print_list(&node)
  consumer(&node)
  consumer(&node)
}
