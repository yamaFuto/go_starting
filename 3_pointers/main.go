package main

//1 pointers

// import "fmt"

// func one(x *int) {
// 	*x = 1
// }

// func main() {
// 	var n int = 100
// 	one(&n)
// 	fmt.Println(n)

// 	//nのアドレス
// 	fmt.Println(&n)

// 	var p *int = &n

// 	//pのアドレス
// 	fmt.Println(p)

// 	//p
// 	fmt.Println(*p)
// }

//2 new make

// import "fmt"

// func main() {

// 	//pointer以外を返す
// 	s := make([]int, 0)
// 	fmt.Printf("%T\n", s)

// 	m := make(map[string]int)
// 	fmt.Printf("%T\n", m)

// 	ch := make(chan int)
// 	fmt.Printf("%T\n", ch)

// 	//pointerを返す
// 	//値を何も入れない状態でメモリにポインタが入る領域を確保する
// 	var p *int = new(int)
// 	fmt.Printf("%T\n", p)

// 	var st = new(struct{})
// 	fmt.Printf("%T\n", st)

// 	// //値を格納せずにメモリを確保する
// 	// //default 0
// 	// var p *int = new(int)
// 	// fmt.Println(p)
// 	// fmt.Println(*p)
// 	// *p++
// 	// fmt.Println(*p)

// 	// //メモリを確保していないためnilが帰ってくる
// 	// var p2 *int
// 	// fmt.Println(p2)
// }

//3 struct

import "fmt"

type Vertex struct {
	//小文字だとプライベートな変数になってしまってパッケーの外から呼び出せなくなってしまう
	X, Y int
	S    string
}
type Test struct {
	X, Y int
	S    *Vertex
}

func changeVertex(v Vertex) {
	v.X = 1000
}

//*がついているとstructと理解して参照先に置き換えてくれる
func changeVertex2(v *Vertex) {
	v.X = 1000
}

func main() {
	v := Vertex{1, 2, "test"}
	fmt.Printf("%T\n %v\n", v, v)
	changeVertex(v)
	fmt.Println(v)

	v2 := &Vertex{1, 2, "test"}
	changeVertex2(v2)
	fmt.Println(v2)

	v3 := Test{1, 2, &Vertex{3, 4, "String"}}
	//structの中のstructはポインタ型だとアドレスを返す
	fmt.Println(v3.S.X)

	// v := Vertex{X: 1, Y:2}
	// fmt.Println(v)
	// fmt.Println(v.X, v.Y)
	// v.X = 100
	// fmt.Println(v.X, v.Y)

	// //{1 0 }
	// v2 := Vertex{X: 1}
	// fmt.Println(v2)

	// //{1 2 "test"}
	// v3 := Vertex{1, 2, "test"}
	// fmt.Println(v3)

	// //main.Vertex {0 0 }
	// v4 := Vertex{}
	// fmt.Printf("%T %v\n", v4, v4)

	// 	//main.Vertex {0 0 }
	// 普通の型と違ってnilにならない
	// var v5 Vertex
	// fmt.Printf("%T %v\n", v5, v5)

	//  struct型はポインタでも初期値はアドレスではなく値を返す
	// 	//*main.Vertex {0 0 }
	// v6 := new(Vertex)
	// fmt.Printf("%T %v\n", v6, v6)

	// //*main.Vertex {0 0 }
	// v7 := &Vertex{}
	// fmt.Printf("%T %v\n", v7, v7)

	//推奨
	// s := make([]int, 0)

	// s := []int{}
}
