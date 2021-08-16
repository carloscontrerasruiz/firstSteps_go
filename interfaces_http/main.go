package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		panic("Error " + err.Error())
	}

	/*bs := make([]byte, 999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))*/
	lw := LogWriter{}
	//io.Copy(os.Stdout, resp.Body)
	io.Copy(lw, resp.Body)
}

type LogWriter struct{}

func (LogWriter) Write(bs []byte) (int, error) {
	//fmt.Println(string(bs))
	fmt.Println(len(bs))
	return len(bs), nil
}
