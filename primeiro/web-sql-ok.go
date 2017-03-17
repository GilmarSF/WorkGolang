package main

import (
	"net/http"
)

var name string
var	peso = 68


func main() {

	connectionHTTP() // chama função servidor http
}


func connectionHTTP() {
	http.HandleFunc("/", serveHome) // Pagina home (localhost:80/)
	http.HandleFunc("/contato", serveContato) // (localhost:80/contato)

	http.ListenAndServe(":80", nil) // usando porta 80, pode ser 8080, apenas alterar
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Paginal Inicial do Gil"))
}

func serveContato(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pagina de contato..."))
}
