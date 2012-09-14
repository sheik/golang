package main

import (
	"fmt"
	"net"
	"time"
)

const (
	master_hostname = "localhost"
	master_port     = ":8080"
)

func transmit(payload string) {
	conn, err := net.Dial("udp", master_hostname+master_port)
	defer conn.Close()

	if err != nil {
		panic(err)
	}

	conn.Write([]byte(payload))
}

func server(done chan<- int) {
	laddr, err := net.ResolveUDPAddr("udp", master_port)
	if err != nil {
		panic(err)
	}
	buffer := make([]byte, 1024)
	for {
		conn, err := net.ListenUDP("udp", laddr)
		if err != nil {
			panic(err)
		}

		for {
			n, err := conn.Read(buffer)
			if err != nil {
				fmt.Println(n)
				panic(err)
			}
			fmt.Printf("%s", buffer[:n])
			if string(buffer[:n]) == "QUIT" {
				conn.Close()
				done <- 0
				return
			}
		}

		conn.Close()
	}
}

func main() {
	done := make(chan int)
	go server(done)
	time.Sleep(1 * time.Second)
	test := ""
	for i := 0; i < 2000; i++ {
		test += "A"
	}
	transmit(test)
	time.Sleep(1 * time.Second)
	transmit("QUIT")
	<-done
}
