package main
import ("fmt"
"net/http")


func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/",homeHandler)
	http.ListenAndServe("localhost:8080",mux)

}


func homeHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"Home")
}