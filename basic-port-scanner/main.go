package main

import(
"net"
"fmt"
"time")

func main(){

	for i:=1;i<=100;i++{
		scanPort(i)
	}
	
}

func scanPort(port int){
addr := fmt.Sprintf("scanme.nmap.org:%v",port)
conn,err := net.DialTimeout("tcp",addr,2*time.Second)
	if err==nil{
		fmt.Printf("The port %v is open",port)
		conn.Close()
	}
}