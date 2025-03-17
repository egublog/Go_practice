package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h2>Hello World</h2>")
}
func main() {
    http.HandleFunc("/", handler)
    fmt.Println("サーバーを起動しました。http://localhost:8081 にアクセスしてください。")
    err := http.ListenAndServe(":8081", nil)
    if err != nil {
        fmt.Printf("サーバーの起動に失敗しました: %v\n", err)
    }
}
