package main

import (
	"fmt"
	"net"
	//"bufio"
	//"strings"
	"os"
)

var cipheredText string

//SlaveChan is lalallala
var SlaveChan chan net.Conn

//TextChan is bababbaba
var TextChan chan string

var slaves []net.Conn
var freeSlaves []int
var counter int
var slaveCounter int
var letterCounter1 int
var letterCounter2 int
var letterCounter3 int
var letterCounter4 int
var ClientConn net.Conn
//var clientAlready bool
//var slaveConn []net.Conn

func handleSlaveConnection(TextChan chan string) {
	//counter:=0
	letters := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
  text := ""
	//if cipheredText == "" {
		//text = <-TextChan
	//} else {
	for  {
		if cipheredText != ""{
			break;
		}
}
		text = cipheredText
		//cipheredText=""
	//}
	for {
		/*if (counter==0 && letterCounter1==0 && letterCounter2==0 && letterCounter3==0 && letterCounter4==0) {
			if cipheredText == "" {
				text = <-TextChan
			} else {
				text = cipheredText
			}
		}*/
		if counter == 7454981 {
			for i := 0; i < len(slaves); i++ {
				slaves[i].Write([]byte("StopIt"))
			}
			break
		}
		if freeSlaves[slaveCounter] == 0 {
			//fmt.Println(text, " ", slaveCounter, " ")
			slaves[slaveCounter].Write([]byte(text))
			fmt.Print("Text Sent to slave ", counter)
			if counter == 0 {
				slaves[slaveCounter].Write([]byte("\n"))
			} else if counter > 0 && counter <= 52 {
				slaves[slaveCounter].Write([]byte(letters[letterCounter1]))
				fmt.Println(" ", letters[letterCounter1])
			} else if counter > 52 && counter <= 2756 {
				slaves[slaveCounter].Write([]byte(letters[letterCounter1] + letters[letterCounter2]))
				fmt.Println(" ", letters[letterCounter1]+letters[letterCounter2])
			} else if counter > 2756 && counter <= 143364 {
				slaves[slaveCounter].Write([]byte(letters[letterCounter1] + letters[letterCounter2] + letters[letterCounter3]))
				//fmt.Print(" ", letters[letterCounter1]+letters[letterCounter2])
			} else if counter > 143364 && counter <= 7454980 {
				slaves[slaveCounter].Write([]byte(letters[letterCounter1] + letters[letterCounter2] + letters[letterCounter3] + letters[letterCounter4]))
				//fmt.Print(" ", letters[letterCounter1]+letters[letterCounter2])
			}
			freeSlaves[slaveCounter] = 1
			go receiveSlaveResponse(slaves[slaveCounter], slaveCounter)

			letterCounter1++
			if counter > 52 && counter%52 == 0 {
				letterCounter2++
			}
			if counter > 2756 && counter%2756 == 0 {
				letterCounter3++
			}
			if counter > 143364 && counter%143364 == 0 {
				letterCounter4++
			}
			counter++
			slaveCounter++

			if letterCounter1 == 52 {
				letterCounter1 = 0
			}
			if letterCounter2 == 52 {
				letterCounter2 = 0
			}
			if letterCounter3 == 52 {
				letterCounter3 = 0
			}
			if letterCounter4 == 52 {
				letterCounter4 = 0
			}
			if slaveCounter == (len(slaves)) {
				slaveCounter = 0
			}
		}
	}
}

func receiveSlaveResponse(c net.Conn, index int) {

	buf := make([]byte, 32)
	n, err := c.Read(buf)
	if err != nil || n == 0 {
		c.Close()
	}

	result := string(buf[0:n])
	//fmt.Println("rsult returned from slave: ", result)
	if result == "unsuccess" {
		freeSlaves[index] = 0
	} else {
		fmt.Println("Found")
		fmt.Println("Password is : "+result)
		counter = 7454981
		ClientConn.Write([]byte(result))
		ClientConn.Close()
		//clientAlready=false
	}
	//fmt.Println("cyphered Text recieved from Server: ", Text)

}
func handleSlaveRequests(SlaveChan chan net.Conn, TextChan chan string) {

	ln, err := net.Listen("tcp", ":5001") //port for slaves
	if err != nil {
		fmt.Println(err)
	}

	for {
		slaveConn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("A Slave has been connected to server")
		fmt.Printf("New Slave: %v\n", slaveConn)
		slaves = append(slaves, slaveConn)
		freeSlaves = append(freeSlaves, 0)
		fmt.Println("Total slaves in channel")
		fmt.Println(len(slaves))
		if len(slaves) == 1 {
			go handleSlaveConnection(TextChan)
		}
	}
}
func handleClientConnection(c net.Conn) {
	buf := make([]byte, 4096)
	n, err := c.Read(buf)
	if err != nil || n == 0 {
		c.Close()
	}
	cipheredText = string(buf[0:n])
}
func main() {
	counter = 7454981
//	clientAlready:=false
	SlaveChan = make(chan net.Conn)
	TextChan = make(chan string)
	go handleSlaveRequests(SlaveChan, TextChan)
	//for client
	port := "2001"
	if len(os.Args) == 2 {
		port = os.Args[1]
	}
	connec := ":"+ port
	Cln, Cerr := net.Listen("tcp", connec)
	if Cerr != nil {
		fmt.Println(Cerr)
	}
	slaveCounter = 0
	for {
		/*for {
			if clientAlready==false{
				break
				}
		}*/
		ClientConn, _ = Cln.Accept()
		fmt.Println("A new client has added")
		handleClientConnection(ClientConn)
	//	clientAlready=true
		counter = 0
		letterCounter1 = 0
		letterCounter2 = 0
		letterCounter3 = 0
		letterCounter4 = 0
		//TextChan <- cipheredText
		fmt.Println("Received Ciphered Text: " + cipheredText)
		cipheredText=""
		//fmt.Println("Received Ciphered Text from channel: " + <-TextChan)
		/*for {
			if counter == 7454981 {
				counter = 0
				slaveCounter = 0
				letterCounter1 = 0
				letterCounter2 = 0
				letterCounter3 = 0
				letterCounter4 = 0
				TextChan <- cipheredText
				fmt.Println("Received Ciphered Text from channel: " + <-TextChan)
				break
			}
		}*/
	}
}
