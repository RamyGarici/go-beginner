package main

import (
	"fmt"
	"net/http"
	"io"
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
	fmt.Fprintln(w,"File received successfully")
	fileData, err := io.ReadAll(file)
	if err!= nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w,"Filename: %v\nFile size: %v",header.Filename,len(fileData))



	
}