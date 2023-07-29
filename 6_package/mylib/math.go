package mylib

import "fmt"

//6 Person2 doc(godocの説明文になる)
type Person2 struct {
	// Name
	Name string
	//Age
	Age int
}

func (p *Person2) Say() {
	fmt.Println("Person2")
}

// gofmtでフォーマットをきれいにすることができる
func Average(s []int) int {
	total := 0
	for _, i := range s {
		total += i
	}
	return int(total / len(s))
}