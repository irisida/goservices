package main

import "net/http"

/*
	simple demo of a webserver (inbuilt) into the core of Go.
	here we import the net/http package and create a HandleFunc
	that takes a route and a function that will be executed when
	we are on that route.

	We need to specify the listenAndServe with the port number
	and how to handle any error raised.
*/

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello world"))
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}
