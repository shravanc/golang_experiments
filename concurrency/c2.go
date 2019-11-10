package main

import ("fmt"
        "time"
        "math/rand")

func foobar(msg string, c chan string){
  for i:=0;;i++{
    c<-fmt.Sprintf("%s %d", msg, i)
    time.Sleep(time.Duration(rand.Intn(1e3))*time.Millisecond)
  }
}

func main(){
  c := make(chan string)
  go foobar("foo!!", c)
  for i:=0; i<10; i++{
    fmt.Printf("You say-->%q \n", <-c)
  }

  fmt.Println("You are boring am leaving")
}
