package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	//t := http.DefaultTransport.(*http.Transport).Clone()
	//t.DisableKeepAlives = true
	//t.IdleConnTimeout = time.Millisecond * 100
	// c := &http.Client{
	// 	Transport: t,
	// }

	for i := 0; i < 100; i++ {
		//res, err := c.Post("http://localhost:8090/hello", "text/json", bytes.NewReader(nil))
		res, err := http.Post("http://localhost:8090/hello", "text/json", bytes.NewReader(nil))
		//res, err := http.Get(url)

		if err != nil {
			log.Printf("error sending: %s", err)
			continue
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Printf("error reading: %s", err)
			continue
		}

		fmt.Printf("iteration #%0d read: %s", i, string(body))

		time.Sleep(time.Millisecond * 200)
	}
}
