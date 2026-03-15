package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Go API")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/health", healthHandler)

	fmt.Println("Servidor rodando em http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erro ao iniciar servidor:", err)
	}
}
