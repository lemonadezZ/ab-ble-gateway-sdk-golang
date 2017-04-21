package main

import (
	//	"bytes"
	"encoding/hex"
	"fmt"
)

func main() {
	b := []byte("\r\n\r\n")
	fmt.Println(hex.DecodeString("feff"))
	fmt.Println(hex.Dump(b))
}
