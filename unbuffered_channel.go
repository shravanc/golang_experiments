package main

import ("fmt")

func hello_channel(c chan string){
  c <- "Hello Channel"
}

func main(){
  var c = make(chan string )

  go hello_channel(c)
  msg := <-c
  fmt.Println(msg)

}
