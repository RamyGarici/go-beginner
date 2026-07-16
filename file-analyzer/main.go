package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


func main(){

	

	mux := http.NewServeMux()
	mux.HandleFunc("/",homeHandler)
	mux.HandleFunc("/upload",uploadHandler)
	http.ListenAndServe("localhost:8080",mux)


}


func homeHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"Home")
}


func uploadHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST"{
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	file, header, err := r.FormFile("myFile")
	
	if err!= nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()
	fileData, err := io.ReadAll(file)
	if err!= nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	
	
	
	resp := response{
		Filename: header.Filename,
		FileSize: len(fileData),
		DetectedType : http.DetectContentType(fileData),
	}
	w.Header().Set("Content-Type","application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	encoder.Encode(resp)
}

type response struct{
		Filename string `json:"filename"`
		FileSize int    `json:"filesize"`
		DetectedType string `json:"detected_type"`
	}