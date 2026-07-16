package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)


var total_uploads, total_bytes int
var mu sync.Mutex


func main(){

	mux := http.NewServeMux()
	mux.HandleFunc("/",homeHandler)
	mux.HandleFunc("/upload",uploadHandler)
	mux.HandleFunc("/stats",statsHandler)
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

	mu.Lock()
	total_bytes+= len(fileData)
	total_uploads+=1
	mu.Unlock()

	
	
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

func statsHandler(w http.ResponseWriter, r *http.Request){
	if r.Method!="GET"{
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	mu.Lock()
	resp := FileStats{
		Uploads: total_uploads,
		TotalBytes: total_bytes,
	}
	mu.Unlock()
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

type FileStats struct{
	Uploads int `json:"total_uploads"`
	TotalBytes int `json:"total_bytes"`
}

