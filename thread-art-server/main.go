package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func imageUploaderHandler(w http.ResponseWriter, req *http.Request) {
	// Code heavily inspired by
	// https://freshman.tech/file-upload-golang/
	if req.Method != "POST" {
		http.Error(w, "HTTP Method not supported. Try again with POST.", http.StatusBadRequest)
		return
	}

	err := os.Mkdir("uploads", os.ModePerm)
	if err != nil && !os.IsExist(err) {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	multipartFile, _, err := req.FormFile("imageupload")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	filename := RandomString(10)
	filename = fmt.Sprintf("%v.png", filename)

	file, err := os.Create(fmt.Sprintf("./uploads/%s", filename))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(file, multipartFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	fmt.Println("Starting API server")
	const port = 8001

	http.HandleFunc("/upload", imageUploaderHandler)

	fmt.Printf("Listening on http://localhost:%v/\n", port)
	var addr = fmt.Sprintf(":%v", port)
	log.Fatal(http.ListenAndServe(addr, nil))
}
