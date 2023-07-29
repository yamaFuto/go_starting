package main

//1 method pointer

// import "fmt"

// type Vertex struct{
// 	X, Y int
// }

// //Vertex struct に関数を結び付けている
// func (v Vertex) Area() int{
// 	return v.X * v.Y
// }

// //pointer レシーバー　メンバの値を変えたいときに使う
// //レシーバの【値型 ⇔ ポインタ型】間の変換はコンパイラが暗黙的におこなってくれる
// func (v *Vertex) Scale(i int) {
// 	v.X = v.X * i
// 	v.Y = v.Y * i
// }

// func Area(v Vertex) int {
// 	return v.X * v.Y
// }

// func main() {
// 	v := Vertex{3, 4}
// 	fmt.Println(Area(v))
// 	v.Scale(10)
// 	fmt.Println(v.Area())
// }

//2 construct

// import "fmt"

// type Vertex struct {
// 	x, y int
// }

// func (v Vertex) Area() int{
// 	return v.x * v.y
// }

// func (v *Vertex) Scale(i int) {
// 	v.x = v.x * i
// 	v.y = v.y * i
// }

// func Area(v Vertex) int {
// 	return v.x * v.y
// }

// //メンバが小文字だと外から直接呼び出せないため関数を問うして代入構築を行う
// func New(x, y int) *Vertex{
// 	return &Vertex{x, y}
// }

// func main() {
// 	v := New(3, 4)
// 	v.Scale(10)
// 	fmt.Println(v.Area())
// }

//3 embedded

// import "fmt"

// type Vertex struct {
// 	x, y int
// }

// func (v Vertex) Area() int{
// 	return v.x * v.y
// }

// func (v *Vertex) Scale(i int) {
// 	v.x = v.x * i
// 	v.y = v.y * i
// }

// type Vertex3D struct {
// 	//embedded
// 	Vertex
// 	z int
// }

// func (v Vertex3D) Area3D() int{
// 	return v.x * v.y * v.z
// }

// func (v *Vertex3D) Scale3D(i int) {
// 	v.x = v.x * i
// 	v.y = v.y * i
// 	v.z = v.z * i
// }

// func Area(v Vertex) int {
// 	return v.x * v.y
// }

// //メンバが小文字だと外から直接呼び出せないため関数を問うして代入構築を行う
// //embeddedされたストラクトの生成
// func New(x, y, z int) *Vertex3D{
// 	return &Vertex3D{Vertex{x, y}, z}
// }

// func main() {
// 	v := New(3, 4, 5)
// 	v.Scale(10)
// 	fmt.Println(v.Area())
// 	fmt.Println(v.Area3D())
// }

//4 non-struct type

// import "fmt"

// //別名をつけて特別なメソッドを持たせることもできる
// type MyInt int

// func (i MyInt) Double() int{
// 	//MyIntをキャストして返り値のintに合わせなくてはならない
// 	return int(i * 2)
// }

// func main() {
// 	myInt := MyInt(10)
// 	fmt.Println(myInt.Double())
// }

//5 interface

// import "fmt"

// type Human interface {
// 	Say() string
// }

// type Person struct {
// 	Name string
// }

// type Dog struct {
// 	Name string
// }

// //実装した中身を操作するので*が必要
// func (p *Person) Say() string {
// 	p.Name = "Mr." + p.Name
// 	fmt.Println(p.Name)
// 	return p.Name
// }

// //interfaceのだっぐタイピング⇒その関数を持っていないstructは受け付けない
// func DriveCar(human Human) {
// 	if human.Say() == "Mr.Mike" {
// 		fmt.Println("Run")
// 	} else {
// 		fmt.Println("Get out")
// 	}
// }

// func main() {
// 	//interfaceの中にstructを入れたらinterfaceの関数を持たなくてはいけない interface以外の関数は除外される
//	// ストラクトのメソッドと違ってポインタ変換されないために初期化するときに&が必要になる
// 	var mike Human = &Person{"Mike"}
// 	var x Human = &Person{"x"}

// 	// var dog Dog = Dog{"dog"}
// 	DriveCar(mike)
// 	DriveCar(x)
// 	// DriveCar(dog)
// }

//6 type assertion

// import "fmt"

// //interface{}⇒どんな方でも受け付ける
// func do(i interface{}) {
// 	//interface()はどんな方でもあるのでtype assertionをする
// 	// ii := i.(int)
// 	// ii *= 2
// 	// fmt.Println(ii)
// 	// ss := i.(string)
// 	// fmt.Println(ss + "!")

// 	//switchとtypeはセット
// 	//type assertionしている
// 	switch v := i.(type) {
// 	case int:
// 		fmt.Println(v * 2)
// 	case string:
// 		fmt.Println(v + "!")
// 	default:
// 		fmt.Printf("I don't know %T\n", v)
// 	}
// }

// func main() {
// 	// var i interface{} = 10
// 	// do(i)
// 	do("mike")
// 	do(10)
// 	do(true)

// 	//interfaceではないtypeからほかのtypeにconversionしている
// 	var i int = 10
// 	ii := float64(10)
// 	fmt.Println(i, ii)
// }

//7 stringer

// import "fmt"

// type Person struct {
// 	Name string
// 	Age int
// }

// //stringerはstruct表示方法を変更する特別な関数
// func (p Person) String() string {
// 	//printfは引数をstring型で返す
// 	return fmt.Sprintf("My name is %v", p.Name)
// }

// func main() {
// 	mike := Person{"Mike", 22}
// //fmtがString()⇒string()を持っているとstringerInterfaceと認識してfmtがString関数を呼び出す
// 	fmt.Println(mike)
// }

//8 error

// import "fmt"

// type UserNotFound struct {
// 	Username string
// }

// func (e *UserNotFound) Error() string{
// 	return fmt.Sprintf("User not found: %v", e.Username)
// }

// func myFunc() error {
// 	//Something wrong
// 	ok := false
// 	if ok {
// 		return nil
// 	}

// 	//Error関数が設定されているために、	初期化後にErrorを返す
// 	//error同士を比較するために区別としてpointerを返している
// 	return &UserNotFound{Username: "mike"}
// }

// func main() {
// 	if err := myFunc(); err != nil {
// 		fmt.Println(err)
// 	}
// }