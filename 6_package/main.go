package main

//1 ~ 4

// import (
// 	"fmt"
// 	"sample-app/6_package/mylib"
// 	"sample-app/6_package/mylib/under"
// )

// func main() {
// 	s := []int{1, 2, 3, 4, 5}
// 	fmt.Println(mylib.Average(s))

// 	mylib.Say()
// 	under.Hello()

// 	person := mylib.Person{Name: "Mike", Age: 20}
// 	fmt.Println(person)

// 	fmt.Println(mylib.Public)

// 	//できない
// 	// fmt.Println(mylib.private)
// }

//5 third package

import (
	"fmt"
	"github.com/markcheno/go-quote"
	//a を_に変えるとつかわなくてもよい
	a "github.com/markcheno/go-talib"
)

func main() {
	spy, _ := quote.NewQuoteFromYahoo("spy", "2016-01-01", "2016-04-01", quote.Daily, true)
	fmt.Print(spy.CSV())
	rsi2 := a.Rsi(spy.Close, 2)
	fmt.Println(rsi2)
}