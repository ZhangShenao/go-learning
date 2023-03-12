package main

import (
	"flag"
	"io"
	"log"
	"net/http"
)

//declare constants
const (
	HealthCheckResponse = "OK"
	VersionKey          = "version"
)

var version string

//编写一个 HTTP 服务器，大家视个人不同情况决定完成到哪个环节，但尽量把 1 都做完：
//接收客户端 request，并将 request 中带的 header 写入 response header
//读取当前系统的环境变量中的 VERSION 配置，并写入 response header
//Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
//当访问 localhost/healthz 时，应返回 200
func main() {
	//Parse Flags
	version = *flag.String(VersionKey, "v1.0", "specify the version of http server")
	flag.Parse()

	//Register Health Check Handler
	http.HandleFunc("/healthz", healthCheckHandler)

	//Register Http Request Handler
	http.HandleFunc("/", requestHandler)

	//Start HttpServer And Listen On Port
	log.Printf("http server started on port: %d, version: %s\n", 80, version)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}

//Health Check Handler
func healthCheckHandler(writer http.ResponseWriter, req *http.Request) {
	io.WriteString(writer, HealthCheckResponse)
}

//Http Request Handler
func requestHandler(writer http.ResponseWriter, req *http.Request) {
	//handle http headers
	for k, v := range req.Header {
		writer.Header().Add(k, v[0])
	}

	writer.Header().Add(VersionKey, version)

	//get remote address
	addr := req.RemoteAddr
	log.Println("request ip: ", addr)
}
