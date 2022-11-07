// 14.    Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var (
		i int
		s string
		b bool
		c chan int
	)

	fmt.Println(reflect.TypeOf(interface{}(i)))
	fmt.Println(reflect.TypeOf(interface{}(s)))
	fmt.Println(reflect.TypeOf(interface{}(b)))
	fmt.Println(reflect.TypeOf(interface{}(c)))

}
