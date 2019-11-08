package main

import ("fmt")

func count(min, max int)int{
  sum:=0
  for i:=min; i<=max;i++{
    sum += i
  }
  return sum
}


func main(){
  min := 1
  max := 500
  interval := 100
  var sums []int

  iteration := max / interval

  for i:=0; i<iteration; i++{
    max = min + interval - 1
    sum := count(min, max)
    sums = append(sums, sum)
    min = max + 1
  }

  total := 0
  for _, val := range sums{
    total = total + val
  }
  fmt.Println(total)


}
