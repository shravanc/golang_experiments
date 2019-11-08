package main

import (
      "fmt"
      "sync"
      )

func  count(min int , max int, wg *sync.WaitGroup, c chan int){
  sum := 0
  for i:=min; i<=max; i++{
    sum += i
  }
  c <- sum
  wg.Done()
}


func main(){
  var wg sync.WaitGroup
  var c = make(chan int )


  max := 500
  min := 1

  interval := 100
  iteration := max / interval

  total := 0
  for i:=0; i<iteration; i++{
    wg.Add(1)
    max = min + interval - 1
    go count(min, max, &wg, c)
    min = max + 1
    total = total + <-c
  }

  wg.Wait()
  fmt.Println(total)
}
