package main
import("fmt"
       "math")


type geometry interface{
  area() float64
  perim() float64
}

type rect struct{
  l float64
  b float64
}

type circle struct{
  r float64
}


func (r rect) area() float64{
  return r.l*r.b
}

func (c circle) area() float64{
  return  (math.Pi * c.r * c.r)
}

func (r rect) perim() float64{
  return (2 * r.l + 2* r.b)
}

func (c circle) perim() float64{
  return (2*math.Pi*c.r)
}

func measure(g geometry){
  fmt.Println( g.area() )
  fmt.Println( g.perim() )
}


func main(){
  fmt.Println("Hello wassup")
  var a = rect{l: 1, b: 2}
  var b = circle{r: 2}

  measure(a)
  measure(b)
}
