package main
import("fmt"
       "time")

type Node struct{
  data int
  next *Node
  prev *Node
}

func produce(node *Node, data int){
  for {
    if node.next==nil{
      new_node := Node{data: data, prev: node}
      node.next = &new_node
      break
    }else{
      node = node.next
    }
  }
}

func consume(node *Node){
  for{
    if node.next == nil{
      prev := node.prev
      prev.next = nil
      fmt.Printf("Consumed: %d\n", node.data)
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

func producer(node *Node, c chan string){
  for i:=0; i<10; i++{
    produce(node, i)
    c <-fmt.Sprintf("Produced--->%d", i)
    fmt.Printf("Producing--->%d\n", i)
    time.Sleep(time.Second*1)
  }
  close(c)
}


func main(){
  var node = Node{data: 1}
  var c = make(chan string)

  go producer(&node, c)

  for msg := range c{
    fmt.Printf("Consuming--->%q", msg)
    consume(&node)
  }

  defer print_list(&node)
}
