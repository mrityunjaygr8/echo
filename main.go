package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
  fmt.Println("yo")

  l, err := net.Listen("tcp", "127.0.0.1:3333")
  if err != nil {
    log.Fatal(err)
  }

  defer l.Close()

  for {
    conn, err := l.Accept()
    if err != nil {
      fmt.Println("error occurred: ", err)
    }

    buf := make([]byte, 0, 4096)
    for {
      tmp := make([]byte, 256)
      n, err := conn.Read(tmp)
      if err != nil {
        if err != io.EOF {
          fmt.Println("error has occurred")
        }
        break
      }
      buf = append(buf, tmp[:n]...)
      if n < 256 {
        break
      }
    }
    conn.Close()
    fmt.Println(string(buf))
  }

}
