package main

import(
"net"
"fmt"
"time"
"sync")

 var wg sync.WaitGroup
func main(){
	ports := make(chan int, 100)

for i:=1;i<=100;i++{
	    go worker(ports, &wg)
	}
   
    
	
	for i:=1;i<=65535;i++{
	    wg.Add(1)
		ports <-i
	}
	
	close(ports)
	
	wg.Wait()
	
}


func scanPort(port int){
addr := fmt.Sprintf("scanme.nmap.org:%v",port)
conn,err := net.DialTimeout("tcp",addr,2*time.Second)
	if err==nil{
		fmt.Printf("The port %v is open\n",port)
		conn.Close()
	}
	
}

func worker(ports <-chan int, wg *sync.WaitGroup){
	for port := range ports{
		scanPort(port)
		wg.Done()
		
	}
}