package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloUser)

	http.HandleFunc("/todos", todoFunction)

	http.ListenAndServe(":8080", nil)
	fmt.Println("Server is running on port 8080")
}

func helloUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, User! This is the start of George's Go journey.\n")
}

func todoFunction(w http.ResponseWriter, r *http.Request) {
	todos := []string{
		"Learn Go",
		"Build a web app",
		"Write tests",
		"Deploy to production",
	}

	fmt.Fprintf(w, "Todos:\n")
	for i, todo := range todos {
		fmt.Fprintf(w, "%d: %s\n", i+1, todo)
	}
	fmt.Fprintf(w, "Reward on completion: $100\n")
}
