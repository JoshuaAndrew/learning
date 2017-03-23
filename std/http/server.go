package main

import (
	"net/http"
	"fmt"
)



func handler(){
	fmt.Println("df")

}

func main() {
	http.ListenAndServe(":80",http.HandlerFunc(handler))
}
