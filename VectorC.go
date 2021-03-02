package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

type Timestamp struct {
	firstprocess  int
	secondprocess int
	thirdprocess  int
	fourthprocess int
}

var chngs1 = "Change for frstuser"
var chngs2 = "Change for scnduser"
var chngs3 = "Change for thrduser"
var chngs4 = "Change for frthuser"
var fdata string

func main() {
	timestamp := Timestamp{0, 0, 0, 0}
	timestampChannel := make(chan *Timestamp)

	go writefrstuserchngs(&timestamp, timestampChannel)
	time.Sleep(time.Second * 3)
	go writescnduserchngs(&timestamp, timestampChannel)
	time.Sleep(time.Second * 4)
	go writethrduserchngs(&timestamp, timestampChannel)
	time.Sleep(time.Second * 5)
	go writefrthuserchngs(&timestamp, timestampChannel)
	time.Sleep(time.Second * 7)

}

func writefrstuserchngs(timestamp *Timestamp, timestampChannel chan *Timestamp) {
	timestamp.firstprocess = timestamp.firstprocess + 1
	fmt.Println("\n The timestamp recorded after event 1 occurred in the first process  ", timestamp)
	go makePrimarychngs(1, timestamp, timestampChannel)

}

func writescnduserchngs(timestamp *Timestamp, timestampChannel chan *Timestamp) {
	timestamp.secondprocess = timestamp.secondprocess + 1

	_ = <-timestampChannel
	fmt.Println("\n event 1 of second process receiving a message from event 2 of first process at given timestamp ", timestamp)
	go makePrimarychngs(2, timestamp, timestampChannel)

}
func writethrduserchngs(timestamp *Timestamp, timestampChannel chan *Timestamp) {
	timestamp.thirdprocess = timestamp.thirdprocess + 1

	_ = <-timestampChannel
	fmt.Println("\n event 1 of third process is recieving a message from event 2 of second process at given timestamp  ", timestamp)
	go makePrimarychngs(1, timestamp, timestampChannel)
}
func writefrthuserchngs(timestamp *Timestamp, timestampChannel chan *Timestamp) {
	timestamp.fourthprocess = timestamp.fourthprocess + 1

	_ = <-timestampChannel
	fmt.Println("\n event 1 of fourth process is recieving a message from event 3 of third process at given timestamp ", timestamp)

}

func makePrimarychngs(id int, timestamp *Timestamp, timestampChannel chan *Timestamp) {

	if id == 1 {
		timestamp.firstprocess += 1
		writeFile(chngs1)
		fmt.Println("\n the time after event 2 in first process occurred at ", timestamp)
	} else if id == 2 {
		timestamp.secondprocess += 1
		writeFile(chngs2)
		fmt.Println("\n  the time after event 2 in second process occurred at ", timestamp)

	}

	timestampChannel <- timestamp

}

func ReadFile(filename string) {

	fmt.Printf("\n \n Reading the file in the Golang and giving next timestamp \n")

	f, err := os.Open("VectorClock.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	data := ""
	for scanner.Scan() {

		data += scanner.Text()
	}
	fdata = data

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	f.Close()

}

func writeFile(data string) {

	ReadFile("VectorClock.txt")

	chngs := fdata + "\n" + data

	f, err := os.Create("VectorClock.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	f.WriteString(chngs)

}
