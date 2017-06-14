package main

import (
	"fmt"
	"net/http"
)

var (
	duration = 0
)

//Test push
func Index(w http.ResponseWriter, r *http.Request) {
	go func() {
		for n := 0; n <= 5; n++ {
			duration++
			fmt.Println(duration)
		}
		duration = 0
		fmt.Println(duration)
	}()
	fmt.Fprintln(w, "Welcome!", duration)
}
