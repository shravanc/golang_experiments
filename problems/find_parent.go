/*
Input:
data = [
        [3, 2],
        [5, 3],
        [4, 5],
        [2, 0]
       ]

Explanation:
3 is child of parent 2

Output:
result = [2,3,5,4]


*/

package main
import ("fmt")

func construct_map(data [4][2]int) map[int]int{

  var hash = make(map[int]int)
  for _, val1 :=  range data{
    hash[val1[1]] = val1[0]
  }
  return hash
}


func get_result(hash map[int]int) []int{
  var result []int
  parent := 0
  for i:=0; i<len(hash); i++{
    result = append(result, hash[parent])
    parent = result[len(result)-1]
  }
  return result
}

func main(){

  var data = [4][2]int{ {3, 2}, {5, 3}, {4, 5}, {2, 0} }

  hash := construct_map(data)
  result := get_result(hash)


  fmt.Println(hash)
  fmt.Println(result)
}
