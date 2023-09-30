package main

//1 if

// import "fmt"

// func by2(num int) string{
// 	if num % 2 == 0{
// 		return "ok"
// 	} else {
// 		return "no"
// 	}
// }

// func main() {
// 	result := by2(10)

// 	if result == "ok" {
// 		fmt.Println("great")
// 	}

// 	//上記の分を一行で書いている
// 	//result2はif分の中でだけ完結している
// 	if result2 := by2(10); result2 == "ok" {
// 		fmt.Println("great 2")
// 	}

// 	num := 6
// 	if num % 2 == 0 {
// 		fmt.Println("by 2")
// 	} else if num % 3 == 0{
// 		fmt.Println("by 3")
// 	} else {
// 		fmt.Println("else")
// 	}
// 	x, y := 11, 12
// 	if x == 10 && y == 10{
// 		fmt.Println("&&")
// 	}

// 	if x == 10 || y == 10 {
// 		fmt.Println("||")
// 	}
// }

//2 for

// import "fmt"

// func main() {
// 	for i := 0; i < 10; i++ {
// 		if i == 3{
// 			fmt.Println("continue")
// 			continue
// 		}

// 		if i > 5 {
// 			fmt.Println("break")
// 			break
// 		}
// 		fmt.Println(i)
// 	}

// 	sum := 1

// 	//;;は省略できる
// 	for ; sum < 10; {
// 		sum += sum
// 		fmt.Println(sum)
// 	}
// 	fmt.Println(sum)

// 	//無限ループ
// 	// for {
// 	// 	fmt.Println("hello")
// 	// }
// }

//3 range

// import "fmt"

// func main() {
// 	l := []string{"python", "go", "java"}

// 	for i := range l {
// 		fmt.Println(i, l[i])
// 	}

// 	for i, v := range l {
// 		fmt.Println(i, v)
// 	}

// 	for _, v := range l {
// 		fmt.Println(v)
// 	}

// 	m := map[string]int{"apple": 100, "banana": 200}

// 	for k, v := range m {
// 		fmt.Println(k, v)
// 	}

// 	//keyはvalueを_する必要がない
// 	for k := range m {
// 		fmt.Println(k)
// 	}

// 	for _, v := range m {
// 		fmt.Println(v)
// 	}
// }

//4 switch

// import (
// 	"fmt"
// 	"time"
// )

// func getOsName() string {
// 	return "ldjfa"
// }

// func main() {
// 	// os := getOsName()
// 	//一文にしてかける
// 	switch os := getOsName(); os {
// 	case "mac":
// 		fmt.Println("Mac")
// 	case "windows":
// 		fmt.Println("Windows")
// 	default:
// 		fmt.Println("Default",os)
// 	}

// 	t := time.Now()
// 	fmt.Println(t.Hour())
// 	//switch変数を書かなくてもよい
// 	switch {
// 		case t.Hour() < 12:
// 			fmt.Println("Morning")
// 		case t.Hour() < 17:
// 			fmt.Println("Afternoon")
// 		}
// }

//5 defer

// import (
// 	"fmt"
// 	"os"
// )

// func foo() {
// 	defer fmt.Println("world foo")

// 	fmt.Println("hello foo")
// }

// func main() {
// 	//defer⇒親関数が実行された後に実行するという意味
// 	defer fmt.Println("world")
// 	fmt.Println("hello")

// 	//hello foo, world foo, hello, world
// 	defer fmt.Println("world")
// 	foo()
// 	fmt.Println("hello")

// 	//run, success, 3, 2, 1, stackの形でdeferが返ってくる
// 	fmt.Println("run")
// 	defer fmt.Println(1)
// 	defer fmt.Println(2)
// 	defer fmt.Println(3)
// 	fmt.Println("success")

// 	//fileのclose処理
// 	file, _ := os.Open("./main.go")
// 	defer file.Close()
// 	data := make([]byte, 100)
// 	file.Read(data)
// 	fmt.Println(string(data))
// }

//6 logging

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// )

// //logをファイルに書き込む
// func LoggingSettings(logFile string) {
// 	//ファイルをオープン　なければ作成,第二引数でOpen方法を記述
// 	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
// 	//画面上に出るエラーを出力する、ログファイルに書き込む
// 	//出力先を決めている
// 	multiLogFile := io.MultiWriter(os.Stdout, logfile)
// 	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
// 	//実行
// 	log.SetOutput(multiLogFile)
// }

// func main() {
// 	LoggingSettings("test.log")
// 	_, err := os.Open("lkfjdlfa")
// 	if err != nil {
// 		log.Fatalln("Exit")
// 	}
// 	log.Println("logging")
// 	log.Printf("%T %v", "test", "test")

// 	//この時点でexitする
// 	log.Fatalf("%T %v", "test", "test")
// 	log.Fatalln("error")
// 	fmt.Println("ok")
// }

//7 errorhandling

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// func main() {
// 	file, err := os.Open("./main.go")
// 	if err != nil{
// 		log.Fatalln("Error")
// 	}
// 	defer file.Close()
// 	data := make([]byte, 100)

// 	//一つでも初期化できていると使える errは上書きされている
// 	count, err := file.Read(data)
// 	if err != nil{
// 		log.Fatalln("Error")
// 	}
// 	fmt.Println(count, string(data))

// 	//error一つしか返さないときは使える
// 	if err = os.Chdir("test"); err != nil{
// 		log.Fatalln("Error!")
// 	}
// }

//8 panic recover

// import "fmt"

// func thirdPartyConnectDB() {
// 	//自ら例外を投げ、強制終了
// 	panic("Unable to connect database")
// }

// func save(){
// 	defer func() {
// 		//例外を受け取りプログラムを再起動する
// 		s := recover()
// 		fmt.Println(s)
// 		}()
//		//これをdeferの上に書くとdefer関数が読み込まれる前に強制終了されてしまう
// 		thirdPartyConnectDB()
// }

// func main() {
// 	save()
// 	fmt.Println("OK")
// }