package main
import("fmt")

type circular struct{
  data int
  next *circular
}

func insert(c *circular, data int){

  node := circular{data: data}
  if c.next == nil{
    c.next = &node
    node.next = c
    return
  }

  first_node := c
  for {
    if c.next == first_node{
      c.next = &node
      node.next = first_node
      break
    }else{
      c = c.next
    }
  }

}

func print_list(c *circular){
  temp := c
  for c1:=c; c1!=nil; c1=c1.next{
    if c1.next == nil{
      fmt.Println(c1.data)
      break
    }
    if c1.next == temp{
      fmt.Println(c1.data)
      break
    }else{
    }
    fmt.Println(c1.data)
  }
}

func main(){
  node := circular{data: 1}
  insert(&node, 2)
  insert(&node, 3)
  insert(&node, 4)
  insert(&node, 5)
  print_list(&node)
}
