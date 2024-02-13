package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	k := "Deep Rock Galactic" //Key value
	fmt.Println("Keyword: " + k)
	s := "I am inside the cave"                       //Message to send
	fmt.Println("Message sent to the receiver: " + s) //Receiver's message
	p := k + s                                        //Payload
	fmt.Println("Key+Payload: " + p)
	h := sha256.New()
	h.Write([]byte(p))
	bs := h.Sum(nil)
	fmt.Printf("Hash value sent to the receiver: %x\n", bs) //Hash value

}
