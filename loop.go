package main

import ("fmt")

func main(){
  var a = []int{1,2,3,4,5}

  //Looping over array
  for index, value := range a{
    fmt.Println(index, value)
  }

  //looping for n times
  end := 5
  for i:=0; i<end; i++{
    fmt.Println(i)
  }

  //looping for ever
  break_me_at_zero := 10
  for {
    fmt.Println("Infite Loop-->", break_me_at_zero)
    if break_me_at_zero == 0{
      break
    }
    break_me_at_zero -= 1
  }

}
