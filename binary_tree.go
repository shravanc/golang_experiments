package main
import("fmt")

type tree struct{
  data int
  lchild *tree
  rchild *tree
}

func insert(t *tree, data int){
  for {
    if t.lchild == nil{
      child := tree{data: data}
      t.lchild = &child
      break
    }
    if t.rchild == nil{
      child := tree{data: data}
      t.rchild = &child
      break
    }
    t = t.lchild
  }
}

func print_tree(t *tree){


  if t.lchild == nil{

    if t.rchild ==nil{
      fmt.Println(t.data)
      return
    }
  }else{
    print_tree(t.lchild)
  }


  if t.rchild == nil{

    if t.rchild == nil{
      fmt.Println(t.data)
      return
    }
  }else{
    print_tree(t.rchild)
  }
  fmt.Println(t.data)
  return
}

func main(){
  var t = tree{data: 1}
  fmt.Println(t)
  insert(&t, 2)
  insert(&t, 3)
  insert(&t, 4)
  insert(&t, 5)
  insert(&t, 6)
  insert(&t, 7)
  fmt.Println(t)

  //fmt.Println(t.data)
  print_tree(&t)
}
