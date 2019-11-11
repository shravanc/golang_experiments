package main
import "fmt"


type valStruct struct {
    value1 int
}

type hello interface{
  inc()int
  dec()int
}

func (r *valStruct) inc() int {
    r.value1 = r.value1 + 1
    return r.value1
}
func (r valStruct) dec() int {
    r.value1 = r.value1 - 1
    return r.value1
}
func main() {
    r := valStruct{value1: 10}
    fmt.Println("value: ", r.inc())
    fmt.Println("value: ", r.inc())
    fmt.Println("value: ", r.inc())
    fmt.Println("value: ", r.inc())
    fmt.Println("value: ", r.inc())

    fmt.Println("value: ", r.dec())
    fmt.Println("value: ", r.dec())
    fmt.Println("value: ", r.dec())
    fmt.Println("value: ", r.dec())
    fmt.Println("value: ", r.dec())

    var b hello

    b = &r
    fmt.Println(b.inc())

}
