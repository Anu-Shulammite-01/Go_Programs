package main

import (
	"fmt"
	"net"
)

func main(){
	ServerAddr,err:=net.ResolveUDPAddr("udp",":8081")
	errorHandle(err)
	Conn,err:=net.ListenUDP("udp", ServerAddr)
	errorHandle(err)
	defer Conn.Close()

	buf:=make([]byte,1024)
	for{
		length,addr,err:=Conn.ReadFromUDP(buf)
		errorHandle(err)
		fmt.Println("Recieved",string(buf[0:length]),"from",addr)
	}

}

func errorHandle(err error){
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
}
