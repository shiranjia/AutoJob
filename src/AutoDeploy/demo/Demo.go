package demo

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"sync"
	"time"
)

func Route() {
	c1 := make(chan int)
	c2 := make(chan int)

	go sum(9000, c2)
	go sum(10000, c1)
	var wait sync.WaitGroup
	wait.Add(1)
	count := 0
	go func() {
	loop:
		for {
			fmt.Println("===loop")
			select {
			case v1 := <-c1:
				fmt.Println("c1------------------------------------------------------------------------------", v1)
			case v2 := <-c2:
				fmt.Println("c2-------------------------------------------------------------------------------", v2)
			case <-time.After(1 * time.Microsecond):
				{
					fmt.Println("time out:", count)
					close(c1)
					close(c2)
					break loop
				}
			default:
				count++
				fmt.Println("---", count)
			}
		}
		file, _ := os.Create("E:/" + strconv.Itoa(count) + ".txt")
		defer file.Close()
		sre := "aaa" + strconv.Itoa(count)
		file.Write([]byte(sre))
		wait.Done()
	}()
	//wait.Wait()
	fm(10)
	fmt.Println("done")
}

func Reflect() {
	p := persion{"asfef", 21}
	var i interface{}
	var m me = p
	i = m
	defer func() {
		if err := recover(); err != nil {
			log.Println("err:", err)
		}
	}()
	if value, ok := i.(a); ok {
		fmt.Println(value)
	}
	switch i.(type) {
	case string:
		fmt.Println("chan")
	case persion:
		fmt.Println("persion")

	}
	var a float32 = 23.34
	t := reflect.ValueOf(&a)
	fmt.Println(t.Type())
	//fmt.Println(t.Elem())
	t.Elem().SetFloat(7.2)
	fmt.Println(a)
}

func fm(c int) {
	for i := 0; i < c; i++ {
		time.Sleep(1 * time.Second)
	}
	fmt.Println("over")
}
func sum(a int, c chan int) {
	sum := 0
	for i := 0; i <= a; i++ {
		sum += i
	}
	file, _ := os.Create("E:/" + strconv.Itoa(sum) + ".txt")
	//file,err := os.Open("E:/a.txt");
	defer file.Close()
	/*if err != nil{
		log.Println(err)
	}*/
	sre := "aaa" + strconv.Itoa(sum)
	file.Write([]byte(sre))
	fmt.Println("sum:", sum)
	//time.Sleep(3 * time.Second)
	c <- sum
}

type persion struct {
	name string
	age  int
}

func (p persion) setName(n string) {
	p.name = n
}

type me interface {
	setName(n string)
}

type a string

func testStruts() {
	var p persion
	p.name = "jiashiran"
	p.age = 10
	p.setName("asfa")
	fmt.Println(p)
}
