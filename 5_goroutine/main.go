package main

//1 goroutine

// import (
// 	"fmt"
// 	"sync"
// 	// "time"
// )

// func goroutine(s string, wg *sync.WaitGroup) {
// 	for i := 0; i < 5; i++ {
// 		// time.Sleep(100 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// 	//Add()分Doneを呼ばなくてはいけない
// 	wg.Done()
// }

// func normal(s string) {
// 	for i := 0; i < 5; i++ {
// 		// time.Sleep(100 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup
// 	//並列処理が一つあることを伝える
// 	wg.Add(1)
// 	//goroutineは並列機能で後の動作と同時に動く
// 	//goroutineはthreadを作る段階でプログラムが終了すると動作せずに終わることががある
// 	//WaitGroupはコピーしてはいけないことになっている
// 	go goroutine("world", &wg)
// 	normal("hello")
// 	// time.Sleep(2000 * time.Millisecond)

// 	//wgがDoneするまで待ってください
// 	wg.Wait()
// }

//2 channel

// import "fmt"

// func goroutine1(s []int, c chan int) {
// 	sum := 0
// 	for _, v := range s{
// 		sum += v
// 	}
// 	//channelに送信している
// 	c <- sum
// }

// func goroutine2(s []int, c chan int) {
// 	sum := 0
// 	for _, v := range s{
// 		sum += v
// 	}
// 	//channelに送信している
// 	c <- sum
// }

// func main() {
// 	s := []int{1, 2, 3, 4, 5}
// 	c := make(chan int) // queueのようなイメージ
// 	go goroutine1(s, c)
// 	go goroutine2(s, c)
// 	//chennelを受け取る
// 	//cが入ってくるまでプログラムはずっとここで待っている
// 	x := <-c
// 	fmt.Println(x)
// 	y := <-c
// 	fmt.Println(y)
// }

//3 buffered channel

// import "fmt"

// func main() {
// 	ch := make(chan int, 2)
// 	ch <- 100
// 	fmt.Println(len(ch))
// 	ch <- 200
// 	fmt.Println(len(ch))
// 	//chennelはcloseしないとmakeの第二引数で設定した一以上に繰り返してしまう可能性がある
// 	close(ch)

// 	for c := range ch{
// 		fmt.Println(c)
// 	}
// }

//4 channel range close

// import "fmt"

// //その都度channelを返す
// func goroutine1(s []int, c chan int) {
// 	sum := 0
// 	for _, v := range s{
// 		sum += v
// 		c <- sum
// 	}
// 	close(c)
// }

// func main() {
// 	s := []int{1, 2, 3, 4, 5}
// 	c := make(chan int, len(s))
// 	go goroutine1(s, c)
// 	for i := range c{
// 		fmt.Println(i)
// 	}
// }

//5 producer consumer

// import (
// 	"fmt"
// 	"sync"
// )

// func producer(ch chan int, i int) {
// 	ch <- i * 2
// }

// func consumer(ch chan int, wg *sync.WaitGroup) {
// 	for i := range ch {
// 		func () {
// 			//fmtが何らかの形で中断してしまってもdeferでDone処理は呼ばれる
// 			defer wg.Done()
// 			fmt.Println("process", i * 1000)
// 		}()
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup
// 	ch := make(chan int)
	
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go producer(ch, i)
// 	}
	
// 	go consumer(ch, &wg)
	
// 	wg.Wait()

// 	//並列で動いているconsumerを停止させるためのclose()
// 	close(ch)
// }

// 6 fan-out fan-in

// import "fmt"

// func producer(first chan int) {
// 	defer close(first)
// 	for i := 0; i < 10; i++ {
// 		first <- i
// 	}
// }

// func multi2(first chan int, second chan int) {
// 	defer close(second)
// 	for i := range first {
// 		second <- i * 2
// 	}
// }

// //このように書くことで受信送信を明示的にしてわかりやすくする
// func multi4(second <-chan int, third chan<- int) {
// 	defer close(third)
// 	for i := range second {
// 		third <- i * 4
// 	}
// }

// func main() {
// 	first := make(chan int)
// 	second := make(chan int)
// 	third := make(chan int)

// 	go producer(first)
// 	go multi2(first, second)
// 	go multi4(second, third)
// 	for result := range third {
// 		fmt.Println(result)
// 	}
// }

// 7 select channel

// import (
// 	"fmt"
// 	"time"
// )

// func goroutine1(ch chan string) {
// 	for {
// 		ch <- "packet from 1"
// 		time.Sleep(1 * time.Second)
// 	}
// }

// func goroutine2(ch chan int) {
// 	for {
// 		ch <- 100
// 		time.Sleep(1 * time.Second)
// 	}
// }

// func main() {
// 	c1 := make(chan string)
// 	c2 := make(chan int)
// 	go goroutine1(c1)
// 	go goroutine2(c2)

// 	//channelで待機するのを防ぐためにselectで分岐している
// 	for {
// 		select {
// 		case msg1 := <-c1:
// 			fmt.Println(msg1)
// 		case msg2 := <-c2:
// 			fmt.Println(msg2)
// 		}
// 	}
// }

//8 default selection for break

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	tick := time.Tick(100 * time.Millisecond)
// 	boom := time.After(500 * time.Millisecond)
// OuterLoop:
// 	for {
// 		select {
// 			//case t := <-tickだが、必要なければt := は省略できる
// 		case <-tick:
// 			fmt.Println("tick.")
// 		case <-boom:
// 			fmt.Println("BOOM!")
// 			break OuterLoop
// 			// break⇒通常のbreakだとselect文を出るだけになってしまう
// 		//上記のcase以外のものがかえってくるとdefaultを返す
// 		default:
// 			fmt.Println("    .")
// 			time.Sleep(50 * time.Millisecond)
// 		}
// 	}
// 	fmt.Println("########")
// }

//9 sync.Mutex

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// type Counter struct {
// 	v map[string]int
// 	mux sync.Mutex
// }

// //同時に同じmapについて捜査してもlockをかけることによって操作を重複させない
// func (c *Counter) Inc(key string) {
// 	c.mux.Lock()
// 	defer c.mux.Unlock()
// 	c.v[key]++
// }

// func (c *Counter) Value(key string) int {
// 	c.mux.Lock()
// 	defer c.mux.Unlock()
// 	return c.v[key]
// }

// func main() {
// 	// c := make(map[string]int)
// 	c := Counter{v: make(map[string]int)}
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			// c["key"] += 1
// 			c.Inc("Key")
// 		}
// 	}()

// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			// c["key"] += 1
// 			c.Inc("Key")
// 		}
// 	}()
// 	time.Sleep(1 * time.Second)
// 	fmt.Println(c.v, c.Value("Key"))
// }
