package mylib

//goのテストはこれを使う
//Ginko, Gomegaを使うとよい
import (
	// "fmt"
	"testing"
)

var Debug bool = true

// //6 godocの説明
// //godoでPackageのExampleとして表示(ExampleはTestに書く)
// func Example() {
// 	v := Average([]int{1, 2, 3, 4, 5})
// 	fmt.Println(v)
// }

// //PersonStructのSay()の説明文がgodocに表示される
// func ExamplePerson2_Say() {
// 	p := Person2{"Mike", 20}
// 	p.Say()
// }

// //godocでAverageのexampleとして表示
// func ExampleAverage() {
// 	v := Average([]int{1, 2, 3, 4, 5})
// 	fmt.Println(v)
// }

//3 test testはtest対象の隣にファイルを置く
//go test ./...でtest実行
func TestAverage(t *testing.T) {
	if Debug {
		//testをkipする
		t.Skip("Skip Reason")
	}
	v := Average([]int{1, 2, 3, 4, 5})
	if v != 3{
		t.Error("Expected 3, got", v)
	}
}