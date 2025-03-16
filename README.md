# Go言語学習プロジェクト

Go言語の基本的な機能を学習するためのプロジェクトです。
主にWebサーバーの実装と基本的なライブラリの作成を通じて、Go言語の特徴を学びます。

## 実行方法

1. Webサーバーの起動
```bash
go run main.go
```
その後、ブラウザで http://localhost:8080 にアクセス

2. ライブラリの使用例
```go
import "mylib"

numbers := []int{1, 2, 3, 4, 5}
average := mylib.Average(numbers)
```

## 開発環境

- Go言語のインストールが必要です
- `go mod tidy` で依存関係を解決してください