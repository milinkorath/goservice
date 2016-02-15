package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
)

func main() {

	if ln, err := net.Listen("tcp", ":8080"); err == nil {
		for {
			if conn, err := ln.Accept(); err == nil {

				reader := bufio.NewReader(conn)
				if req, err := http.ReadRequest(reader); err == nil {
					fmt.Println(req.Body)

				}

			}
		}
	}

}
