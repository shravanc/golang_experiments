package main

import (
	"fmt"
  "time"
  "encoding/json"
   "log"
)

type Node struct {
	Next  *Node
	Value interface{}
}

//Creating Struct For link list
type LinkedList struct {
	Length int // will be used for counting the Length
	Head   *Node
	Tail   *Node
}


// to push a new node  at the end of the link list
func (l *LinkedList) Push(val interface{}) {
	node := &Node{Value: val}
	if l.Head == nil {
		l.Head = node
	} else {
		l.Tail.Next = node
	}
	l.Tail = node
	l.Length = l.Length + 1
}

// this will pop the last element in the link list
func (l *LinkedList) Pop() {
	node := l.Head
	if l.Length == 1 {
		l.Head = nil
	} else {
		for i := 1; i < l.Length-1; i++ {
			node = node.Next
		}
    node.Next = node.Next.Next
	}
	l.Length = l.Length - 1
}

func producers(l *LinkedList, c1, c2, c4 chan string, number int){
  for i:=1; i< number+1 ;i++{
    l.Push(i)
    c1 <- "produced"
    c4 <- "produced"
    <-c2
  }
  close(c4)
  close(c1)
}

func consumer(l *LinkedList, c1, c2, c3 chan string){
  for _ = range c1{
    l.Pop()
    c2 <- "consumed"
  }
  close(c2)
  c3 <- "completed"
}



func print_data(l *LinkedList, c4 chan string){
  data, err := json.MarshalIndent(l, "", "  ")

  for _ = range c4{
    fmt.Println("*********************")
    data, err = json.MarshalIndent(l, "", "  ")
    if err != nil {
      log.Fatal(err)
    }
    fmt.Printf("%s\n", data)
  }
}

func main() {
	var l LinkedList
  var number int
  var c1 = make(chan string, 4)
  var c2 = make(chan string, 4)
  var c3 = make(chan string)
  var c4 = make(chan string)

	fmt.Println("Enter the number of producers you want: ")
	fmt.Scanf("%d", &number)

  start := time.Now()

  go producers(&l, c1, c2, c4, number)
  go consumer(&l, c1, c2, c3)
  go print_data(&l, c4)

  <-c3
  elapsed := time.Since(start)
  fmt.Println("Time taken is: ", elapsed)
}
