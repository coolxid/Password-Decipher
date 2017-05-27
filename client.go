package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func handleConnection(c net.Conn) {
	b := bufio.NewReader(c)
	for {
		line, err := b.ReadString('\n')
		if err != nil { // EOF, or worse
			break
		}
		c.Write([]byte(line))
	}
}

func main() {
	host := "127.0.0.1"
	port := "2001"
	//cipheredText := "Bz.wTj/mBaM"
	cipheredText := "kmscp9jrzH2"
	if len(os.Args) < 2 {
		fmt.Println("Ciphered Password not entered so decrypting kmscp9jrzH2")
		//os.Exit(1)
	} else if len(os.Args) == 2 {
		cipheredText = os.Args[1]
		host = "127.0.0.1"
		port = "2001"
	} else if len(os.Args) == 3 {
		cipheredText = os.Args[1]
		host = os.Args[2]
		port = "2001"
	} else if len(os.Args) == 4 {
		cipheredText = os.Args[1]
		host = os.Args[2]
		port = os.Args[3]
	} else {
		host = "127.0.0.1"
		port = "2001"
		cipheredText = "kmscp9jrzH2"
	}
	fmt.Print("no. of args")
	fmt.Println(len(os.Args))
	fmt.Println("text: "+os.Args[1])
	fmt.Println("host: "+os.Args[2])
	fmt.Println("port: "+os.Args[3])
	connection := string(host + ":" + string(port))
	c, err := net.Dial("tcp", connection)
	if err != nil {
		//	log.Fatal(err)
		fmt.Println(err)
	}
	c.Write([]byte(cipheredText))
	fmt.Print("Ciphered Password Sent")
	go handleConnection(c)
	buf := make([]byte, 32)
	n, err := c.Read(buf)
	if err != nil || n == 0 {
		c.Close()
	}
	result := string(buf[0:n])
	fmt.Println("\nPassword is : "+result)

}
