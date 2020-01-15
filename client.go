package main

import (
	"io"
	"net"
	"os"
)

const message = `GET / HTTP/1.1
Host: localhost:8080
User-Agent: Mozila
Cookie: ck=xxxxxx
Connection: close

`

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	conn.Write([]byte(message))
	io.Copy(os.Stdout, conn)
}
