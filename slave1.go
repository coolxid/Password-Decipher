package main

import (
	"fmt"
	"net"
	"github.com/kahootali/go-crypt"
	//"bufio"
	//"strings"

)

//TextChan is hahaha
var TextChan chan string

func getPassword(c net.Conn) {
	//buf := make([]byte, 4096)
	//n, err := c.Read(buf)
	//if err != nil || n == 0 {
	//c.Close()

	//}
	//Text := string(buf[0:n])
	//TextChan <- Text
	//fmt.Println("cyphered Text recieved from Server: ", Text)
	for {
		buf := make([]byte, 512)
		n, err := c.Read(buf)
		if err != nil || n == 0 {
			c.Close()
		}
		Text := string(buf[0:n])

		//checking if the pw has been cracked and process we should stop!?
		if Text == "StopIt" {
			fmt.Println("ok i have stopped now, send me new work if u want")
			/*buf = []byte("\n")
			c.Read(buf)
			if err != nil || n == 0 {
				c.Close()
			}
			Text = string(buf[0:n])*/
		} else{
		fmt.Println("Ciphered Text recieved from Server: ", Text)
		counter := 0
		letters := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
			"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

		buf = []byte("\n")
		buf2 := make([]byte, 40)
		n, err = c.Read(buf2)
		if err != nil || n == 0 {
			c.Close()
		}
		startString := string(buf2[0:n])
		fmt.Println("StartString recieved from Server: ", startString)
		text := "unsuccess"
		for {
			if counter == 52 {
				break
			} else if crypt.Crypt(startString+letters[counter],"") == Text {
				fmt.Println(startString + letters[counter])
				counter++
				text = startString+letters[counter-1]
				fmt.Print("Crypted for aaaa is:  ")
				fmt.Println(crypt.Crypt("aaaa",""))
				break
			} else {
				fmt.Println(startString + letters[counter])
				counter++
			}
		}
		c.Write([]byte(text))
		buf2 = []byte("\n")
	}
	}
}

func main() {
	TextChan = make(chan string)
	conn, err := net.Dial("tcp", "127.0.0.1:5001")
	if err != nil {
		//	log.Fatal(err)
	}
	fmt.Print(" Slave connected")
	go getPassword(conn)
	text := <-TextChan
	fmt.Println("Password received : " + text)
}
