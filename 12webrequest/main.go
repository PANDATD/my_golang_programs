package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://github.com/pandatd/"				

func main() {
	fmt.Println("websier request")

	responce, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%T", responce)

	defer responce.Body.Close()

	databytes, err := ioutil.ReadAll(responce.Body)
	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)

}
