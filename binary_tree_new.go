package main
import("fmt")

type tree struct{
  data int
  lchild *tree
  rchild *tree
}


func my_insert(t *tree, data int){
  if t.lchild == nil{
    t.lchild = &tree{data: data}
    return
  }else if t.rchild == nil{
    t.rchild = &tree{data: data}
    return
  }

  if child == "left"{
    my_insert(t.rchild, data, "right")
  }else{
    my_insert(t.lchild, data, "left")
  }
}


func insert(t *tree, data int){
  if data < t.data{
    if t.lchild == nil{
        t.lchild = &tree{data: data}
        return
    }else{
      insert(t.lchild, data)
    }
  }else{
    if t.rchild == nil{
      t.rchild = &tree{data: data}
      return
    }else{
      insert(t.rchild, data)
    }
  }
}

func print_tree(t *tree){
  if t.lchild != nil{
    fmt.Println("going left", t.data)
    print_tree(t.lchild)
  }
  fmt.Println(t.data)
  if t.rchild != nil{
    fmt.Println("going right", t.data)
    print_tree(t.rchild)
  }
}


func main(){
  var t = tree{data: 8}
  my_insert(&t, 2, "")
  my_insert(&t, 3, "")
  my_insert(&t, 5, "")
  my_insert(&t, 6, "")
  my_insert(&t, 0, "")
  print_tree(&t)
}
