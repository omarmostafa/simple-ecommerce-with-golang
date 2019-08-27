package main

import (
	"fmt"
	"net/http"
)

func main()  {
	http.ListenAndServe(":8080", ChiRouter().InitRouter())
	fmt.Printf("Listen to localhost:8080")
}