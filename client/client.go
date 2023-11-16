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
	t := http.DefaultTransport.(*http.Transport).Clone()
	//t.DisableKeepAlives = true // no EOF with this
	//t.IdleConnTimeout = time.Millisecond * 100 // or no EOF with this
	c := &http.Client{
		Transport: t,
	}

	for i := 0; i < 100; i++ {
		req, _ := http.NewRequest(http.MethodPost, "http://localhost:8090/hello", bytes.NewReader(nil))
		//req.Header.Set("X-Idempotency-Key", "true") // no EOF

		req.Close = true // no EOF
		res, err := c.Do(req)

		//res, err := c.Get("http://localhost:8090/hello") // no EOF (GET is idempotent)

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

		time.Sleep(time.Millisecond * 200) // to cause EOF
	}
}
