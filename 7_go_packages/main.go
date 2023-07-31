package main

// //1 time
// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	t := time.Now()
// 	fmt.Println(t)
// 	//format変換
// 	fmt.Println(t.Format(time.RFC3339))
// 	fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour())
// }

// //2 regexp(正規表現)

import (
	"fmt"
	"regexp"
)

func main() {
	//第一引数に正規表現、第二引数に対象となる文字列を書き、その正規表現にマッチしているのか判定する
	match, _ := regexp.MatchString("a([a-z]+)e", "apple")
	fmt.Println(match)

	//正規表現(判定)の関数化(最適化)
	r := regexp.MustCompile("a([a-z]+)e")
	ms := r.MatchString("apple")
	fmt.Println(ms)

	// s := "/views/test"
	r2 := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	//第一引数に引っかかった文字列全部を持ってきてくれる
	fs := r2.FindString("/view/test")
	fmt.Println(fs)

	//引っかかった部分の中でも一部を取り出したいときに使う
	// /で分かれたパターンでスライスに格納される
	//URLの状況によって処理を変えたいときなどに使う
	fss := r2.FindStringSubmatch("/view/test")
	fmt.Println(fss, fss[0], fss[1], fss[2])
	fss = r2.FindStringSubmatch("/edit/test")
	fmt.Println(fss, fss[0], fss[1], fss[2])
	fss = r2.FindStringSubmatch("/save/test")
	fmt.Println(fss, fss[0], fss[1], fss[2])
}

// //3 sort

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	i := []int{5, 3, 2, 8, 7}
// 	s := []string{"d", "a", "f"}
// 	p := []struct {
// 		Name string
// 		Age int
// 	}{
// 		{"Nancy", 20},
// 		{"Vera", 40},
// 		{"Mike", 30},
// 		{"Bob", 50},
// 	}
// 	fmt.Println(i, s, p)
// 	sort.Ints(i)
// 	sort.Strings(s)
// 	sort.Slice(p, func(i, j int) bool { return p[i].Name < p[j].Name})
// 	sort.Slice(p, func(i, j int) bool { return p[i].Age < p[j].Age})
// 	fmt.Println(i, s, p)
// }

// //4 iota

// import "fmt"

// //自動的にconstの連番を作ってくれる
// const (
// 	//二つ目以降は省略も可
// 	c1 = iota
// 	c2
// 	c3
// )

// const (
// 	_ = iota
// 	KB int = 1 << (10 * iota)
// 	MB
// 	GB
// )

// func main() {
// 	fmt.Println(c1, c2, c3)
// 	fmt.Println(KB, MB, GB)
// }

// //5 context
//go routine1をキャンセル、タイムアウトするときに使う

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// //context.Context⇒contextの型定義（これをすることによってcontextを使っていることを明示的に関数に示す）
// func longProcess (ctx context.Context, ch chan string) {
// 	fmt.Println("run")
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("finish")
// 	ch <- "result"
// }

// func main() {
// 	ch := make(chan string)
// 	//contextを生成
// 	ctx := context.Background()
// 	//contextにタイムアウトをつけて再定義,それに紐づいたcancel関数を生成
// 	//動作開始後に上限タイムを超えたらctxが強制終了
// 	ctx, cancel := context.WithTimeout(ctx, 1 * time.Second)

// 	//timeoutなどを特に決めていなくて形式的にcontextをgoroutineに渡したいときはTODO()で初期化するとよい
// 	// ctx := context.TODO()

// 	//cancelが呼ばれたら強制終了
// 	defer cancel()

// 	//go routineに渡すことによってctxが動作開始する
// 	go longProcess(ctx, ch)

// 	CTXLOOP:
// 			for {
// 				select {
// 					//ctxが強制終了するとDone()が動作する
// 				case <- ctx.Done():
// 					fmt.Println(ctx.Err())
// 					break CTXLOOP
// 				case <- ch:
// 					fmt.Println("success")
// 					break CTXLOOP
// 				}
// 			}
// 	fmt.Println("###########")
// }

// //6 ioutil
//ioに特化している（ネットワーク関係）

// import (
// 	"bytes"
// 	"fmt"
// 	"io/ioutil"
// 	// "log"
// )

// func main() {
// 	// //main.goのファイルのないようをget
// 	// content, err := ioutil.ReadFile("main.go")
// 	// if err != nil{
// 	// 	log.Fatalln(err)
// 	// }
// 	// fmt.Println(string(content))

// 	// //main.goのファイルの内容をコピーした別ファイルの生成
// 	// if err := ioutil.WriteFile("ioutil_temp.go", content, 0666); err != nil{
// 	// 	log.Fatalln(err)
// 	// }

// 	//ネットワークから来た情報を一時的に記憶する記憶領域
// 	r := bytes.NewBuffer([]byte("abc"))
// 	content, _ := ioutil.ReadAll(r)
// 	fmt.Println(string(content))
// }