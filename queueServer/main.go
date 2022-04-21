package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var letters = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

func RandomString(length int, chars []byte) string {
	clen := len(chars)
	b := make([]byte, length)
	r := make([]byte, length) // storage for random bytes.

	// if seed is not set it is 1, which makes for not
	// so random numbers
	rand.Seed(time.Now().UnixNano())
	rand.Read(r)

	for i, rb := range r {
		c := int(rb)
		b[i] = chars[c%clen]
	}

	return string(b)
}

func main() {
	fmt.Println("Starting API server")
	const port = 8001

	http.HandleFunc("/upload", func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			var _, _, _ = req.FormFile("imageupload")
			// var file, fileHeader, err = req.FormFile("imageupload")
			var filename = RandomString(10, letters)
			filename = fmt.Sprintf("%v.png", filename)
			fmt.Println(filename)

			// os.WriteFile(filename, file)
		} else {
			fmt.Println("Else")
		}

	})

	fmt.Printf("Listening on http://localhost:%v/\n", port)
	var addr = fmt.Sprintf(":%v", port)
	log.Fatal(http.ListenAndServe(addr, nil))
}
