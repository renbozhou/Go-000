package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main(){
	fmt.Println("tcp client starting ")
	conn, err := net.Dial("tcp", "127.0.0.1:8009")
	if err != nil {
		fmt.Println("Connect to TCP server failed ,err:",err)
		return
	}
	fmt.Println("input your name:")
	nameInput := bufio.NewReader(os.Stdin)
	if err != nil {
		fmt.Println("Read from console failed,err:")
		return
	}
	name,err :=  nameInput.ReadString('\n')
	name = strings.Replace(name, "\n", "", -1)
	fmt.Println("开始聊天:")
	inputReader := bufio.NewReader(os.Stdin)
	buf := make([]byte, 1024)
	for  {
		input,err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("Read from console failed,err:",err)
			return
		}

		str := strings.TrimSpace(input)
		if str == "quit"{
			break
		}else {
			input = name+"::"+input
		}


		_, err = conn.Write([]byte(input))
		if err != nil{
			fmt.Println("Write failed,err:",err)
			break
		}

		cnt, err := conn.Read(buf)

		if err != nil {
			fmt.Printf("客户端读取数据失败 %s\n", err)
			continue
		}

		//回显服务器端回传的信息
		fmt.Print("接收：" + string(buf[0:cnt]))
	}
}

