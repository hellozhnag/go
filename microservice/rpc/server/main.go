package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	// http:127.0.0.1:8001/add?a=1&b=2
	// 返回的格式化：json {"data":3}
	// call id：r.URL.Path
	// 数据传输协议http1.X
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm() //解析参数
		fmt.Println("path: ", r.URL.Path)
		a, _ := strconv.Atoi(r.Form["a"][0])
		b, _ := strconv.Atoi(r.Form["b"][0])
		w.Header().Set("Content-Type", "application/json")
		jData, _ := json.Marshal(map[string]int{
			"data": a + b,
		})
		_, _ = w.Write(jData)
	})

	_ = http.ListenAndServe(":8001", nil)

}
