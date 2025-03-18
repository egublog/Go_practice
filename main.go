package main

import (
	"Go_practice/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	userHandler := handlers.NewUserHandler()

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/users", userHandler.ListUsers).Methods("GET")
	api.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	api.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userHandler.GetUser(w, r, vars["id"])
	}).Methods("GET")
	api.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userHandler.UpdateUser(w, r, vars["id"])
	}).Methods("PUT")

	port := ":8080"
	fmt.Printf("サーバーを起動しました。http://localhost%s にアクセスしてください。\n", port)
	fmt.Println("利用可能なエンドポイント:")
	fmt.Println("- GET    /api/users     (ユーザー一覧の取得)")
	fmt.Println("- POST   /api/users     (新規ユーザーの作成)")
	fmt.Println("- GET    /api/users/{id} (特定のユーザーの取得)")
	fmt.Println("- PUT    /api/users/{id} (ユーザー情報の更新)")
	
	log.Fatal(http.ListenAndServe(port, r))
}
