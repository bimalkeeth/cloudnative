package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	chock:=make(chan string,8)
	redVelvet:=make(chan string,8)
	go cakeMaker("chocolate",4,chock)
	go cakeMaker("velvet",4,redVelvet)

	moreChocolote:=true
	moreRedVelvet:=true
    var cake string

    for moreChocolote || moreRedVelvet{
    	select{
    	    case cake,moreChocolote = <- chock :
    	    	if moreChocolote{
    	    		fmt.Printf("got the cake from the first factory : %s\n",cake)
				}
		case cake,moreRedVelvet =<-redVelvet:
			if moreRedVelvet{
				fmt.Printf("got the cake from the second factory : %s\n",cake)
			}
        case <-time.After(500 * time.Microsecond):
        	fmt.Println("time out")
        	return
		}
	}

	
}

func cakeMaker(kind string,number int,to chan<- string){

	rand.Seed(time.Now().Unix())
	for i:=0;i<number;i++{
		time.Sleep(time.Duration(rand.Intn(500)) * time.Microsecond)
		to<-kind
	}
	close(to)
}