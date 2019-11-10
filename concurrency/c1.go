package main

import ("fmt"
        "time")

func first_function( c chan string){
  fmt.Println("1---Going to sleep")
  time.Sleep(2 * time.Second)
  fmt.Println("1---Woke up. Sending it now")
  c <- "yo passing it from first routine. Can you receive it from the second"
  fmt.Println("1---Channel updated. Can you get it buddy??")

  time.Sleep(2*time.Second)
  fmt.Println("1---Woke up again")
  for msg := range c{
    fmt.Println("1---I got your message as well buddy. Thank you so much. I will exit now. My job is done.")
    fmt.Println(msg)
    break
  }
  fmt.Println("Done with the first job as well")
}

func second_function(c chan string){
  fmt.Println("2---Here I am from the second function!!!")
  for msg := range c{
    fmt.Println("111")
    fmt.Println(msg)
    break
  }
  fmt.Println("2---I received your message. There you go!!")

  fmt.Println("Okay Take it. Ihave done my first work. Thank you for your help. This is my deduction from the first job. Now its your job.")
  c <- "Reporting about the first job!!"

  fmt.Println("Exiting the second function")

}

func main(){
  var c = make(chan string)

  go first_function(c)
  go second_function(c)

  fmt.Scanln()
  fmt.Println("finishing")
}
