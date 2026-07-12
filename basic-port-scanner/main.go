package main

import(
"net"
"fmt"
"time"
"sync")

 var wg sync.WaitGroup
func main(){
   
	
	for i:=1;i<=100;i++{
	    wg.Add(1)
		go scanPort(i)
	}
	wg.Wait()
	
}

func scanPort(port int){
addr := fmt.Sprintf("scanme.nmap.org:%v",port)
conn,err := net.DialTimeout("tcp",addr,2*time.Second)
	if err==nil{
		fmt.Printf("The port %v is open\n",port)
		conn.Close()
	}
	wg.Done()
}