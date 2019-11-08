package main

import (
      "fmt"
      "sync"
      )

func allow_me_to_exit(wg *sync.WaitGroup){
  fmt.Println("Okay!! since you have asked, Why not...")
  wg.Done()
}

func main(){
  var wg sync.WaitGroup

  wg.Add(1)

  fmt.Println("Allow to exit Sir?")
  go allow_me_to_exit(&wg)

  wg.Wait()

  fmt.Println("Completed")
}
