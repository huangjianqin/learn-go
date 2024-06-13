package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	server, err := net.Listen("tcp", ":8888")
	if err == nil {
		defer server.Close()
		go listen(server)
	} else {
		fmt.Printf("server start error %v \r\n", err)
		os.Exit(1)
	}

	serverAddr := ":8888"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", serverAddr)
	if err != nil {
		fmt.Printf("server address parse error %v \r\n", err)
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Printf("client connect error %v \r\n", err)
		os.Exit(1)
	}

	go handleResponse(conn)
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		// 读取用户输入
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read user input err: ", err)
		} else {
			input = strings.Replace(input, "\n", "", -1)
			if input == "exit" {
				break
			}
			cWrite(conn, []byte(input))
		}
	}
}

func listen(server net.Listener) {
	for {
		conn, err := server.Accept()
		if err == nil {
			fmt.Printf("accept connection %v \r\n", conn.RemoteAddr())
			go handleRequest(conn)
		} else {
			fmt.Printf("accept connection error: %v \r\n", err)
		}
	}
}

func handleRequest(c net.Conn) {
	for {
		buf := make([]byte, 2048)
		cnt, err := c.Read(buf)
		if err == nil {
			content := string(buf[:cnt])
			fmt.Printf("server receive request from connection(%v) %v \r\n", c.RemoteAddr(), content)
			_, err := c.Write([]byte("Thanks got " + content))
			if err != nil {
				fmt.Printf("server response to connection(%v) error %v \r\n", c.RemoteAddr(), err)
			}
		} else {
			if err == io.EOF {
				fmt.Printf("connection(%v) closed \r\n", c.RemoteAddr())
				c.Close()
				return
			}
			fmt.Printf("server read request from connection(%v) error %v \r\n", c.RemoteAddr(), err)
		}
	}
}

func handleResponse(c net.Conn) {
	for {
		read := cRead(c)
		if read == net.ErrClosed {
			fmt.Printf("remote server(%v) closed \r\n", c.RemoteAddr())
			return
		}
		fmt.Printf("client receive response %v \r\n", read)
	}
}

func cRead(conn net.Conn) any {
	buf := make([]byte, 2048)
	cnt, err := conn.Read(buf)
	if err == nil {
		return string(buf[:cnt])
	} else {
		fmt.Printf("connection(%v) read error %v \r\n", conn.LocalAddr(), err)
		return err
	}
}

func cWrite(conn net.Conn, bytes []byte) {
	_, err := conn.Write(bytes)
	if err != nil {
		fmt.Printf("connection(%v) write to remote error %v \r\n", conn.RemoteAddr(), err)
	}
}
