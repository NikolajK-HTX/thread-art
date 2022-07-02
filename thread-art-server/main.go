package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const UPLOAD_PATH = "/home/nikolaj/Documents/Git/thread-art/thread-art-server/upload"

func imageUploaderHandler(w http.ResponseWriter, req *http.Request) {
	// Code heavily inspired by
	// https://freshman.tech/file-upload-golang/

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "POST")

	if req.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if req.Method != "POST" {
		http.Error(w, "HTTP Method not supported. Try again with POST.", http.StatusBadRequest)
		return
	}

	err := os.Mkdir(UPLOAD_PATH, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	multipartFile, _, err := req.FormFile("image-upload")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := RandomString(10)
	filename := fmt.Sprintf("%v.png", id)
	filepath := filepath.Join(UPLOAD_PATH, filename)

	file, err := os.Create(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(file, multipartFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(id))
}

func main() {
	fmt.Println("Starting API server")
	const port = 8001

	http.HandleFunc("/upload", imageUploaderHandler)

	fmt.Printf("Listening on http://localhost:%v/\n", port)
	var addr = fmt.Sprintf(":%v", port)
	log.Fatal(http.ListenAndServe(addr, nil))
}
