package main

import "fmt"

// 型スイッチ
func handle(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int: ", v)
	case float64:
		fmt.Println("float: ", v)
	case string:
		fmt.Println("string: ", v)
	default:
		fmt.Println("default: ", v)
	}
}

// 型アサーション
func typeassert(i interface{}) {
	n1, ok := i.(int)
	fmt.Println(n1, ok)
	n2, ok := i.(float64)
	fmt.Println(n2, ok)
	n3, ok := i.(string)
	fmt.Println(n3, ok)
}

func main() {
	var v interface{}

	v = 100
	fmt.Println(v)
	handle(v)
	typeassert(v)

	v = 42.42
	fmt.Println(v)
	handle(v)
	typeassert(v)

	v = "hoge"
	fmt.Println(v)
	handle(v)
	typeassert(v)
}
