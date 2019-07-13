package main

import (
	"fmt"
	"sync"
)


func main() {

	for i:=0;i<10;i++{
		fib:=lucasoid(0,1,i)
		luc:=lucasoid(2,1,i)
		fmt.Printf("i:%d fib: %d luc: %d\n",i,fib,luc)

	}


}

func lucasoid(a,b,n int) int{
	if n==0{
		return a
	}
	if n==1{
		return b
	}
	return lucasoid(a,b,n-1) + lucasoid(a,b,n-2)
}

func increment(ptr *int, wg *sync.WaitGroup){
	i:=*ptr
	fmt.Printf("is is %d",i)
	*ptr=i+1
	wg.Done()
}