package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello, world!")
		if err != nil {
			fmt.Println(err)

		}
		fmt.Println(fmt.Sprintf("Number of ytes written: %d", n))
	})

	fmt.Println("starting go web server app...")
	_ = http.ListenAndServe(":8080", nil)
}
