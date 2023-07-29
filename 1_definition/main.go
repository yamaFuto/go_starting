package main

//1.hello world
// import "fmt"
//init ⇒　main ⇒　その他の関数の順に呼ばれる
// func init() {
//     fmt.Println("init");
// }

// func bazz() {
//     fmt.Println("Bazz")
// }

// func main() {
//     // name := "Go Developers"
//     // fmt.Println("Azure for", name)
//     // fmt.Println("a");
//     // bazz();

// }

//2.import
//複数のimport を一行で書ける
//os/user⇒osというパッケージの中のuserというライブラリ
// import (
//     "fmt"
//     "os/user"
//     "time"
// )

// func main() {
//     fmt.Println("Hello world", time.Now())
//     fmt.Println(user.Current())
// }

//3　変数
// import "fmt"

// func main() {
//     // var i int = 1
//     // var f64 float64 = 1.2
//     // var s string = "test"
//     // var t, f bool = true, false

//     //まとめられる、値を入れないと初期値が入る
//     //再宣言できない
//     var (
//         i int = 1
//         f64 float64 = 1.2
//         s string = "test"
//         t, f bool = true, false
//     )

//     //短くかける（関数内でしか書けないやり方や⇒main関数の外ではできない）
//     //再宣言ができない
//     xi := 1
//     xf64 := 1.2
//     xs := "test"
//     xt, xf := true, false
//     fmt.Println(i, f64, s, t, f);
//     fmt.Println(xi, xf64, xs, xt, xf);

//     //型調査 \n改行
//     fmt.Printf("%T\n", xf64);
// }

//3 const
// import "fmt"

// const Pi = 3.14

// const (
//     Username = "test_user"
//     Password = "test_pass"
// )

// func main() {
//     fmt.Println(Pi, Username, Password)
// }

//4　数値
// import "fmt";

// func main() {
//     // var (
//     //     u8 uint8      = 255
//     //     i8 int8       = 127
//     //     f32 float32   = 0.2
//     //     c64 complex64 = -5 + 121
//     // )
//     // fmt.Println(u8, i8, f32, c64)
//     // fmt.Printf("type=%T value=%v", u8, u8)

//     x := 1 + 1
//      x++
//      x--
// fmt.Println(1 << 2) 0100
//     fmt.Println(x)
//     fmt.Println("1 + 1 =", 1+1)
// }

//5　文字列
// import (
//     "fmt"
//     "strings"
// )

// func main() {
//     fmt.Println("HEllo World")
//     fmt.Println("HEllo" + "World")
//     fmt.Println(string("Hello World"[0]))
//     var s string = "Hello World"

//     //sをコピーして置き換えている
//     s = strings.Replace(s, "H", "X", 1)
//     fmt.Println(s)
//     //true
//     fmt.Println(strings.Contains(s, "World"))
//     //そのまま改行も表現する
//     fmt.Println(`Test
//             Test
//     Test`)

//     //""の中で”を表現する
//     fmt.Println("\"")
//     fmt.Println(`"`)
// }

//6　bool

// import "fmt"

// func main() {
//     t, f := true, false
//     //%tでbool型ではないと表示できないようにしている
//     fmt.Printf("%T %v %t\n", t, t, t)
//     fmt.Printf("%T %v %t\n", f, f, f)

//     fmt.Println(false || true)
//     fmt.Println(false && true)
//     fmt.Println(!true)
// }

//7　型変換

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
//     var x int = 1
//     xx := float64(x)
//     fmt.Printf("%T %v %f\n", xx, xx, xx)

//     var y float64 = 1.2
//     yy := int(y)
//     fmt.Printf("%T %v %d\n", yy, yy, yy)

//     //string⇒intはint()ではできない
//     //integer,errorを返す
//     //_にすれば使わなくてもerrorをはかなくなる
//     var s string = "14"
//     i, _ := strconv.Atoi(s)
//     fmt.Printf("%T %v", i, i)

//     h := "Hello World"
//     fmt.Println(string(h[0]))
// }

//8　配列

// import "fmt"

// func main() {
// 	var a [2]int
// 	a[0] = 100
// 	a[1] = 200
// 	fmt.Println(a);

// 	//配列はサイズを変更できない
// 	// var b [2]int = [2]int{100, 200}
// 	// fmt.Println(b)

// 	//スライスはサイズを変更できる
// 	var b []int = []int{100, 200}
// 	b = append(b, 300)
// 	fmt.Println(b);
// }

//9　スライス

// import "fmt"

// func main() {
// 	n := []int{1, 2, 3, 4, 5, 6}
// 	fmt.Println(n)
// 	fmt.Println(n[2])
// 	fmt.Println(n[2:4])
// 	fmt.Println(n[:2])
// 	fmt.Println(n[2:])
// 	fmt.Println(n[:])

// 	n[2] = 1000
// 	fmt.Println(n)

// 	var board = [][]int{
// 		{0, 1, 2},
// 		{3, 4, 5},
// 		{6, 7, 8},
// 	}
// 	fmt.Println(board)

// 	n = append(n, 100, 200, 300, 400)
// 	fmt.Println(n)
// }

//10 make cap

// import "fmt"

// func main() {
// 	// makeでキャパシティを決めた状態で初期化している
// 	n := make([]int, 3, 5)
// 	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
// 	n = append(n, 0, 0)
// 	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
// 	n = append(n, 0, 0)
// 	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)

// 	//第３引数がないときはcapとlenが第２引数に統一される
// 	a := make([]int, 3)
// 	fmt.Printf("len=%d cap=%d value=%v\n", len(a), cap(a), a)

// 	//配列を０で初期化する方法
// 	b := make([]int, 0)

// varの場合には初期化される際に値を書かなかったらメモリを保存してくれない
// 	var c []int
// 		fmt.Printf("len=%d cap=%d value=%v\n", len(b), cap(b), b)
// 		fmt.Printf("len=%d cap=%d value=%v\n", len(c), cap(c), c)

// 		// c = make([]int, 5)
// 		c = make([]int, 0, 5)
// 		for i := 0; i < 5; i++ {
// 			c = append(c, i)
// 			fmt.Println(c)
// 		}
// 		fmt.Println(c)

// }

//11 map

// import "fmt"

// func main() {
// 	m := map[string]int{"apple": 100, "banana": 200}
// 	fmt.Println(m)
// 	fmt.Println(m["apple"])
// 	m["banana"] = 300
// 	fmt.Println(m)
// 	m["new"] = 500
// 	fmt.Println(m)

// 	//0
// 	fmt.Println(m["nothing"])

// 	//100 true 二つ目の返り値を無視することができる
// 	// v := m["apple"]
// 	v, ok := m["apple"]
// 	fmt.Println(v,ok)

// 	//0 false
// 	v2, ok2 := m["nothing"]
// 	fmt.Println(v2, ok2)

// 	//メモリ上にからのmapを作ってから値を代入することができる
// 	m2 := make(map[string]int)
// 	m2["pc"] = 5000
// 	fmt.Println(m2)

// 	//varで宣言するだけだとメモリを確保してくれないためにerrorが起こる
// 	// var m3 map[string]int
// 	// m3["pc"] = 5000
// 	// fmt.Println(m3)

// 	// //sliceの時も同様
// 	// var s []int 
// 	// if s == nil {
// 	// 	fmt.Println("Nil")
// 	// }
// }

//12 byte

// import "fmt"

// func main() {
// 	b := []byte{72,73}
// 	//[72 73]
// 	fmt.Println(b)
// 	//HI
// 	fmt.Println(string(b))

// 	c := []byte("HI")
// 	//[72 73]
// 	fmt.Println(c)
// 	//HI
// 	fmt.Println(string(c))
// }

//13 関数

// import "fmt"

// func add(x, y int) (int, int) {
// 	return x + y, x - y
// }

// func cal(price, item int) (result int) {
// 	result = price * item
// 	return
// }

// func convert(price int) float64{
// 	return float64(price)
// }

// func main() {
// 	r1, r2 := add(10, 20)
// 	fmt.Println(r1, r2)

// 	r3 := cal(100, 2)
// 	fmt.Println(r3)

// 	//変数の中に格納できる
// 	f := func(x int) {
// 		fmt.Println("inner func", x)
// 	}
// 	f(1)

// 	func(x int) {
// 		fmt.Println("inner func", x)
// 	}(1)
// }

// 14 クロージャ

// import "fmt"

// func incrementGenerator() (func() int) {
// 	x := 0
// 	return func() int {
// 		x++
// 		return x
// 	}
// }

// //基礎変数を変えながら関数を生成する事ができる
// func circleArea(pi float64) func(radius float64) float64{
// 	return func (radius float64) float64{
// 		return pi * radius * radius
// 	}
// }

// func main() {
// 	counter := incrementGenerator()
// 	fmt.Println(counter())
// 	fmt.Println(counter())
// 	fmt.Println(counter())
// 	fmt.Println(counter())

// 	c1 := circleArea(3.14)
// 	fmt.Println(c1(2))
// 	fmt.Println(c1(3))

// 	c2 := circleArea(3)
// 	fmt.Println(c2(2))
// 	fmt.Println(c2(3))
// }

//15　可変長引数

// import "fmt"

// func foo (params ...int){
// 	fmt.Println(len(params), params)
// 	for _, param := range params{
// 		fmt.Println(param)
// 	}
// }

// func main() {
// 	foo()
// 	foo(10, 20)
// 	foo(10, 20, 30)
// 	s := []int{1, 2, 3}
// 	fmt.Println(s)

// 	foo(s...)
// }