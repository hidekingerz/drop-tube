# DropTube コーディングスタイルガイド

## 概要
DropTubeプロジェクトにおけるGoコードの統一されたスタイルガイドです。保守性と可読性を重視したコーディング規約を定めています。

## 基本原則
- シンプルで読みやすいコードを書く
- Goの慣用的な書き方に従う
- パフォーマンスよりも可読性を優先する（ボトルネック以外）
- 一貫性を保つ

## コードフォーマット

### gofmt/goimports
- 全てのコードは`gofmt`でフォーマットする
- インポートの整理には`goimports`を使用する

### 行の長さ
- 1行は120文字以内を推奨
- 長い行は適切な位置で改行する

## 命名規則

### パッケージ名
- 小文字のみ使用
- 短く、意味のある名前
- アンダースコアやキャメルケースは使用しない

```go
// Good
package downloader
package config

// Bad
package downloadManager
package config_util
```

### 変数・関数名
- キャメルケースを使用
- パブリック: 大文字で始まる
- プライベート: 小文字で始まる
- 略語は大文字で統一（URL、HTTP、ID等）

```go
// Good
var userID string
var httpClient *http.Client
func ParseURL(url string) error

// Bad
var userId string
var httpClient *http.Client
func parseUrl(url string) error
```

### 定数
- 大文字のスネークケース
- パッケージレベルの定数はプレフィックスを付ける

```go
const (
    DEFAULT_OUTPUT_DIR = "."
    DEFAULT_FORMAT     = "best"
    MAX_RETRY_COUNT    = 3
)
```

### 構造体
- キャメルケース
- 意味のある名前を付ける

```go
type DownloadConfig struct {
    OutputDir    string
    Format       string
    Quality      string
    AudioOnly    bool
    AudioFormat  string
}
```

## ファイル・ディレクトリ構成

### ファイル名
- 小文字とアンダースコアを使用
- 機能を表す名前を付ける

```
youtube_downloader.go
config_parser.go
file_utils.go
```

### ディレクトリ構成
```
cmd/drop-tube/          # エントリーポイント
internal/               # 内部パッケージ
  ├── downloader/       # ダウンロード機能
  ├── config/           # 設定管理
  └── cli/              # CLI関連
pkg/                    # 公開パッケージ
  └── utils/            # ユーティリティ
```

## エラーハンドリング

### エラーの作成
- `errors.New()`または`fmt.Errorf()`を使用
- エラーメッセージは小文字で始める
- 文脈情報を含める

```go
// Good
return fmt.Errorf("failed to download video %s: %w", videoID, err)

// Bad
return errors.New("Download failed")
```

### エラーのラップ
- `fmt.Errorf()`の`%w`を使用してエラーをラップ
- 呼び出し元でのエラー判定を考慮

```go
if err := downloadVideo(url); err != nil {
    return fmt.Errorf("video download failed: %w", err)
}
```

### パニックの使用
- プログラムが継続できない致命的なエラーのみ
- 通常のエラーハンドリングでpanicは使用しない

## ログ

### ログレベル
- ERROR: エラー情報
- WARN: 警告情報
- INFO: 一般的な情報
- DEBUG: デバッグ情報

### ログフォーマット
- 構造化ログを使用
- 文脈情報を含める

```go
log.Info("starting download",
    "url", videoURL,
    "format", config.Format,
    "output", config.OutputDir)
```

## テスト

### テストファイル命名
- `*_test.go`サフィックス
- テスト対象と同じパッケージに配置

### テスト関数命名
- `Test`プレフィックス
- テスト内容を表す名前

```go
func TestDownloadVideo_Success(t *testing.T) {}
func TestParseURL_InvalidURL(t *testing.T) {}
```

### テーブル駆動テスト
- 複数のテストケースがある場合は使用

```go
func TestParseURL(t *testing.T) {
    tests := []struct {
        name    string
        url     string
        want    *VideoInfo
        wantErr bool
    }{
        {"valid url", "https://youtube.com/watch?v=123", &VideoInfo{ID: "123"}, false},
        {"invalid url", "invalid", nil, true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := ParseURL(tt.url)
            if (err != nil) != tt.wantErr {
                t.Errorf("ParseURL() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("ParseURL() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

## コメント

### パッケージコメント
- パッケージの目的と機能を説明
- `package`文の直前に記述

```go
// Package downloader provides YouTube video downloading functionality.
// It supports various formats and quality options.
package downloader
```

### 関数コメント
- パブリック関数には必須
- 関数の動作、引数、戻り値を説明

```go
// DownloadVideo downloads a YouTube video from the given URL.
// It returns the path to the downloaded file or an error if the download fails.
func DownloadVideo(url, outputDir string) (string, error) {
    // implementation
}
```

### 型コメント
- パブリック型には必須
- 型の目的と使用方法を説明

```go
// Config represents the configuration for video downloading.
// It contains all the necessary parameters for customizing the download process.
type Config struct {
    OutputDir   string
    Format      string
    Quality     string
}
```

## インポート

### インポート順序
1. 標準ライブラリ
2. サードパーティライブラリ
3. 内部パッケージ

```go
import (
    "fmt"
    "os"
    "path/filepath"

    "github.com/spf13/cobra"
    "github.com/schollz/progressbar/v3"

    "github.com/hidekingerz/drop-tube/internal/config"
    "github.com/hidekingerz/drop-tube/pkg/utils"
)
```

## 禁止事項

- `panic()`の不適切な使用
- グローバル変数の乱用
- 長すぎる関数（50行以上は要検討）
- 深すぎるネスト（3レベル以上は要検討）
- マジックナンバーの使用（定数化する）

## 推奨ツール

### 静的解析
- `go vet`: 静的解析
- `golint`: リンティング
- `staticcheck`: 高度な静的解析

### フォーマット
- `gofmt`: コードフォーマット
- `goimports`: インポート整理

### テスト
- `go test`: テスト実行
- `go test -race`: レース条件検出
- `go test -cover`: カバレッジ計測

## 参考資料
- [Effective Go](https://golang.org/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Uber Go Style Guide](https://github.com/uber-go/guide)