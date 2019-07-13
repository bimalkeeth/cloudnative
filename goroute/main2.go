package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {

	buffer:=make(chan string,1024)
	numWorker:=10000
    var wg sync.WaitGroup
	for i:=0;i<numWorker;i++{
		go WebGetWorker(buffer,&wg)
	}
    urls:=[6]string{"http://example.com","http://packtpub.com","http://reddit.com","http://twitter.com","http://facebook.com","https://en.wikipedia.org/wiki/Main_Page"}

    for i:=0;i<100;i++{
		for _,item :=range urls{
			wg.Add(1)
			buffer<- item
		}
	}
	wg.Wait()
}

func WebGetWorker(in <-chan string,wg *sync.WaitGroup){

	for{
		url:=<-in
		res,err:=http.Get(url)
		if err!=nil{
			fmt.Println(err.Error())
		}else{
			fmt.Printf("Get %s : %d \n",url,res.StatusCode)
		}
        wg.Done()
	}
}