package mylib

import "fmt"

var Public string = "Public"
var private string = "private"

// 変数、struct, メンバの名前は大文字でなくてないとexportされたことにはならない
type Person struct {
	Name string
	Age int
}

func Say() {
	fmt.Println("Human")
	fmt.Println(private)
}