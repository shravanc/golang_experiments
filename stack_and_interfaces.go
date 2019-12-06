/*
func () push(){

}

func () pop()string{

}

func() setSize(int){

}
*/

package main
import("fmt")


//**************************Interface***********************************
type stack interface{
  pop() string
  push(string)
  setSize(int)
  printData()
}
//**************************Interface***********************************


//**************************Structs***********************************
type array struct{
  data []string
  size int
}

type linked_list struct{
  data string
  next *linked_list
  prev *linked_list
  size int
}


type binary_tree struct{
  data int
  lchild *binary_tree
  rchild *binary_tree
}
/*
type avl_tree struct{
  data int
  lchild *avl_tree
  rchild *avl_tree
}
*/
//**************************Structs***********************************


//**************************Array***********************************
func (arr *array) push( data string){
  arr.data = append(arr.data, data)
}

func (arr *array) pop()string{
  prev_ind := 0
  for ind, val := range arr.data{
    if val == ""{
      prev_ind = ind
      break
    }else{
      prev_ind = ind
    }
  }
  fmt.Println("Popping out", arr.data[prev_ind])
  arr.data[prev_ind] = ""
  return "a"
}

func (arr *array) setSize(size int){
  arr.size = size
}

func (arr *array) printData(){
  fmt.Println("data:", arr.data)
}
//**************************Array***********************************

//**************************LinkedList***********************************
func (ll *linked_list) push( data string){
  for {
    if ll.next == nil{
      new_ll := linked_list{data: data, prev: ll}
      ll.next = &new_ll
      break
    }else{
      ll = ll.next
    }
  }
}

func (ll *linked_list) pop()string{
  for i:=0; i<ll.size; i++{
    if ll.next ==nil {
      old_node := ll.prev
      old_node.next = nil
      break
    }else{
      ll = ll.next
    }
  }
  return "hello"
}

func(ll *linked_list) setSize(size int){
  ll.size = size
}

func (ll *linked_list) printData(){
  fmt.Print("data:")
  for i:=ll; i!=nil; i=i.next{
    fmt.Print(i.data)
  }
  fmt.Println()
}
//**************************LinkedList***********************************


func producer(s stack){
  s.setSize(4)

  s.push("a")
  s.push("b")
  s.push("c")
  s.printData()
  s.pop()
  s.printData()
  s.pop()
  s.printData()

}

func main(){
  var a array;
  var b linked_list;
  producer(&a)
  producer(&b)
}


