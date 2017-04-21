package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	//	"io"
	//	"buffer"
	"io/ioutil"
	//	"os"
	//"encoding/hex"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("Gateway is running.")
}
func (this *MainController) Post() {
	f, h, err := this.GetFile("chunk")
	defer f.Close()
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println(" =========================")
		fmt.Println(h.Filename)
		fmt.Println(h.Header)
		fd, _ := h.Open()
		fmt.Println(&fd)
		fmt.Println(" =========================")
		b, _ := ioutil.ReadAll(fd)
		//fmt.Println(string(b))
		fmt.Println("bufferString")
		//		split := bytes.NewBuffer([]byte("\r\n\r\n"))
		p := bytes.Split(b, []byte("\r\n\r\n"))
		header := string(p[0])
		payloadbyte := p[1]
		fmt.Println(" =========================header")
		fmt.Println(header)
		fmt.Println(" =========================payload")
		payload := bytes.Split(payloadbyte, []byte("\r\n"))
		for _, record := range payload {
			//fmt.Println(record)
			fmt.Println(hex.Dump(record))
		}

		fmt.Println(" =========================")
	}
}

func main() {
	beego.Router("/", &MainController{})
	beego.Run()
}
