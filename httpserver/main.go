package main

import "net/http"

//简单的HTTP服务器
func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello world~"))
	})
	http.ListenAndServe(":8080", nil)
}
