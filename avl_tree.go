package main
import("fmt")

type Avl struct{
  data int
  l_child *Avl
  r_child *Avl
  parent *Avl
  r_height int
  l_height int
  longest_path int
}


func analyse(avl *Avl, data int){
  p1 := avl.parent
  p2 := p1.parent

  p3 := p2.parent

}


func update_parent(avl *Avl, count){
  child = avl
  for{
    if avl.parent ==nil{
      data = avl.height - count
      data = data * data
      if data >i= 4{
        analyse(child)
      }
    }
  }
}


func insert(avl *Avl, data int, ){
  if data < avl.data{
    fmt.Println("Less than satisfied", data)
    if avl.l_child == nil{
      child := Avl{data: data, parent: avl}
      avl.l_child = &child
      return
    }else{
      insert(avl.l_child, data)
    }
  }else{
    fmt.Println("Greater than satisfied", data)
    if avl.r_child == nil{
      child := Avl{data: data}
      avl.r_child = &child
    }else{
      insert(avl.r_child, data)
    }
  }
}

func print_tree(t *Avl){


  if t.l_child == nil{

    if t.r_child ==nil{
      fmt.Println(t.data)
      return
    }
  }else{
    print_tree(t.l_child)
  }


  if t.r_child == nil{

    if t.r_child == nil{
      fmt.Println(t.data)
      return
    }
  }else{
    print_tree(t.r_child)
  }
  fmt.Println(t.data)
  return
}

func main(){
  var a = Avl{data: 41, r_height: 0, l_height: 0}
  insert(&a, 20)
  insert(&a, 55)
  insert(&a, 11)
  insert(&a, 50)
  insert(&a, 65)
  insert(&a, 29)
  insert(&a, 26)
  insert(&a, 23)
  print_tree(&a)
}
