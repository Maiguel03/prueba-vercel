package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "¡Hola desde mi servidor en Go!")
	})

	port := ":8080" // Puerto en el que escuchará el servidor
	fmt.Printf("Servidor escuchando en http://localhost%s\n", port)
	http.ListenAndServe(port, nil)
}
