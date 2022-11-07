// 7.	Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Data struct {
	mx sync.Mutex
	M  map[string]int
}

func NewData() *Data {
	return &Data{
		M: make(map[string]int),
	}
}

func (c *Data) Get(key string) (int, bool) {
	c.mx.Lock()
	defer c.mx.Unlock()
	val, ok := c.M[key]
	return val, ok
}

func (c *Data) Put(key string, value int) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.M[key] = value
}

func MTest(p string, d *Data) {
	for i := 0; i < 100; i++ {
		d.Put(p+"_"+strconv.Itoa(i), i)
	}
}

func main() {

	// sync.Mutex
	dataMutex := NewData()

	MTest("prefix_1", dataMutex)
	MTest("prefix_2", dataMutex)

	time.Sleep(10 + time.Second)

	for k, v := range dataMutex.M {
		fmt.Println(k, ":", v)

	}

	// sync.Map
	var dataSM sync.Map
	go func(p string) {
		for i := 0; i < 100; i++ {
			dataSM.Store(p+"_"+strconv.Itoa(i), i)
		}

	}("prefix_1_SM")

	go func(p string) {
		for i := 0; i < 100; i++ {
			dataSM.Store(p+"_"+strconv.Itoa(i), i)
		}
	}("prefix_2_SM")

	time.Sleep(10 + time.Second)

	dataSM.Range(func(k, v interface{}) bool {
		fmt.Println(k, ":", v)
		return true
	})
}
