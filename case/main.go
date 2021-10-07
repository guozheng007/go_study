package main

import (
	"fmt"
	"github.com/golang/glog"
	"io"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header{
		for _, h := range headers {
			//fmt.Fprintf(w, "%v: %v\n", name, h)
			//1、请求头添加到响应头
			w.Header().Add(name,h)

		}
	}

	env := os.Environ()
	for index := range env{
		fmt.Println(env[index])
	}

	//2、读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	w.Header().Add("Version",os.Getenv("GOVERSION"))

	fmt.Fprintf(w, "hello\n")

	//3、Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	io.WriteString(w, fmt.Sprintf("客户端IP:%v，HTTP 返回码:%v\n",req.Host,http.StatusOK))
}

//4、当访问 localhost/healthz 时，应返回200
func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, fmt.Sprintf("状态码：%v",http.StatusOK))
}

func main() {

	glog.V(2).Info("Starting http server...")
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/healthz",healthz)

	http.ListenAndServe(":80", nil)

}