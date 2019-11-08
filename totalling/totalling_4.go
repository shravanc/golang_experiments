package main

import ("fmt"
        "sync"
       )


func sum(min int, max int, wg *sync.WaitGroup, c chan int){
  total := 0
  for i:=min; i<=max; i++{
    total += i
  }
  c <- total
  wg.Done()
}


func main(){
  max := 500
  min := 1

  intervals := 100
  iterations := max / intervals

  var c = make(chan int, iterations)
  var wg sync.WaitGroup

  total := 0
  for i:=0; i<iterations; i++{
    wg.Add(1)
    max = min + intervals -1
    go sum(min, max, &wg, c)
    min = max + 1
    total += <-c
  }

  wg.Wait()
  fmt.Println(total)

}
