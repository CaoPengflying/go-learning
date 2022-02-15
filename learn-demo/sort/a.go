package main

import (
	"fmt"
	"time"
)
func main() {
	//a := 0
	//fmt.Scan(&a)
	//fmt.Printf("%d\n", a)
	fmt.Printf("Hello World!\n");
	//课程表 t1

	//学生表 t2

	//学生课程表 t3

	//查询同时在上2门课及以上的学生
	// select * from t2 where sid in (select sid from t3 group by sid having count(*) > 1)
	ch1 := make(chan int)
	ch2 := make( chan int)
	go func(){
		for i:=0; i<100; i+=2 {
			<- ch1
			fmt.Println(i)
			ch2 <- i
		}
	}()

	go func(){
		for i:=1; i<100; i+=2 {
			<- ch2
			fmt.Println(i)
			ch1 <- i
		}
	}()
	ch1 <- 0
	time.Sleep(10*time.Second)


}

type CQueue struct {
	vaules []int
}


func Constructor() CQueue {
	return CQueue{vaules:make([]int,0)}
}


func (this *CQueue) AppendTail(value int)  {
	this.vaules = append(this.vaules,value)
}


func (this *CQueue) DeleteHead() int {
	value := this.vaules[0]
	this.vaules = this.vaules[1:]
	return value
}
