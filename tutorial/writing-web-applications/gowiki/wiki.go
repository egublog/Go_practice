// Wikiページの作成、閲覧、編集機能を提供するWebアプリケーション
package main

import (
    "fmt"
    "html/template" // HTMLテンプレート処理用
    "log"          // ログ出力用
    "net/http"     // HTTPサーバー機能用
    "os"           // ファイル操作用
)

// メインエントリーポイント
// URLルーティングの設定とHTTPサーバーの起動を行う
func main() {
    http.HandleFunc("/view/", viewHandler) // /view/へのリクエストをviewHandlerで処理
    http.HandleFunc("/edit/", editHandler) // /edit/へのリクエストをeditHandlerで処理
    log.Fatal(http.ListenAndServe(":8080", nil)) // 8080ポートでサーバーを起動
}

// ページ表示用のハンドラー関数
// /view/で始まるURLに対して、指定されたページの内容を表示する
func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):] // URLからページタイトルを抽出
    p, _ := loadPage(title)             // ページを読み込み
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body) // HTMLとしてページを表示
}

// Wikiページの構造体定義
// Title: ページのタイトル
// Body: ページの本文（バイト配列として保存）
type Page struct {
    Title string
    Body  []byte
}

// タイトルを指定してページを読み込む関数
// 指定されたタイトルのテキストファイルからページ内容を読み込み、Page構造体として返す
func loadPage(title string) (*Page, error) {
    filename := title + ".txt"           // ファイル名を生成
    body, err := os.ReadFile(filename)   // ファイルを読み込み
    if err != nil {
        return nil, err                  // エラーが発生した場合はエラーを返す
    }
    return &Page{Title: title, Body: body}, nil // ページオブジェクトを返す
}

// ページ編集用のハンドラー関数
// /edit/で始まるURLに対して、編集フォームを表示する
func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]  // URLからページタイトルを抽出
    p, err := loadPage(title)            // ページを読み込み
    if err != nil {
        p = &Page{Title: title}          // 存在しない場合は新規ページとして扱う
    }
    t, _ := template.ParseFiles("edit.html")  // 編集用テンプレートを読み込み
    t.Execute(w, p)                          // テンプレートを実行してフォームを表示
}
