package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	// http://127.0.0.1:8000/add?a=1&b=2
	// 返回的格式化:json{"data":3}
	// calllId的问题:r.Url Path
	// 数据的传输协议:http的url参数
	// 网络协议:http

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm() // 解析参数
		fmt.Println("path:", r.URL.Path)
		a, _ := strconv.Atoi(r.Form["a"][0])
		b, _ := strconv.Atoi(r.Form["b"][0])
		w.Header().Set("Content-Type", "application/json")
		jData, _ := json.Marshal(map[string]int{
			"data": a + b,
		})
		w.Write(jData)
	})

	http.ListenAndServe(":8080", nil)
}
