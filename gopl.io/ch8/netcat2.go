package main

import (
	"net"
	"log"
	"os"
	"io"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopyForNetcat2(os.Stdout, conn)
	mustCopyForNetcat2(conn, os.Stdin)
}

func mustCopyForNetcat2(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
