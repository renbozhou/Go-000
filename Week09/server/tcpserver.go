package main

import (
	"fmt"
	"net"
	"sync"
)
var connChanMap sync.Map
func main()  {

	fmt.Println("tcp listener starting ")

	ln,err:=net.Listen("tcp","0.0.0.0:8009")

	if err!=nil{
		fmt.Println("Listen tcp server failed, err:", err)
		return
	}

	for{
		conn,err := ln.Accept()
		if err != nil {
			fmt.Println("Listen.Accept failed, err:",err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Println("conn process ")
	for{
		var buf [128]byte
		n,err:= conn.Read(buf[:])
		if err != nil {
			fmt.Println("Read from tcp server failed,err:",err)
			removeConnFromMap(conn)
			break
		}
		data := string(buf[:n])


		sendMsgToChan(data,conn)
		fmt.Printf("Recived from client,data:%s\n",data)
	}
}

func removeConnFromMap(conn net.Conn) {
	 connChanMap.Delete(conn)
}

func sendMsgToChan(data string, conn net.Conn) {

	if msgChan,ok := connChanMap.Load(conn); ok {
		  msgChan.(chan string) <- data
	}else {
		msgChan := make(chan string,100)
		connChanMap.Store(conn,msgChan)
		msgChan <- data
	}

	sendMsgToAllConn()
}

func sendMsgToAllConn() {
	connChanMap.Range(func(k, v interface{}) bool {

		fmt.Println("iterate:", k, v)

		go sentMsgToAllClient(v.(chan string),k.(net.Conn))
		return true
	})
}

func sentMsgToAllClient(v chan string, conn net.Conn) {
	for{
		msg := <-v
		fmt.Println("需要发送的msg: ",  msg)
		connChanMap.Range(func(k, v interface{}) bool {

			fmt.Println("iterate:", k, v)
			if (conn!=k) {
				k.(net.Conn).Write([]byte(msg))
			}
			return true
		})
	}

}


