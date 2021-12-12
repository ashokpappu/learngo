package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Printf("Hello World!\n")

	sec, _ := time.ParseDuration("5s")

	sockets := make([]string, 5)
	for i := 1; i <= 1024; i++ {
		addr := fmt.Sprintf("scanme.nmap.org:%d", i)
		con, err := net.DialTimeout("tcp", addr, sec)
		if err != nil {
			//fmt.Println(err)

			continue
		}

		if con != nil {
			fmt.Println(con.LocalAddr().Network())
			sockets = append(sockets, addr)
		}

		con.Close()
	}

	for i := 0; i < len(sockets); i++ {
		fmt.Println(sockets[i])

	}
}
