package main

import ("fmt"
        "time"
        "sync")

func one(c chan string, wg *sync.WaitGroup){
  c <- "Hello its me buddy, Can you reach to this message?"
  wg.Done()
}

func two(c chan string, wg *sync.WaitGroup){
  msg := <- c
  fmt.Println(msg)
  fmt.Println("See I told you I can find your message!!!")
  wg.Done()
}


func main(){
  var c = make(chan string)
  var wg sync.WaitGroup

  wg.Add(1)
  go one(c, &wg)
  time.Sleep(1 * time.Second)


  wg.Add(1)
  go two(c, &wg)

  wg.Wait()
  fmt.Println("Finishing without a glitch")
}
