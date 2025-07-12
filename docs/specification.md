# DropTube - YouTube動画ダウンロードアプリ 仕様書

## 概要
コマンドライン引数で指定されたYouTube動画URLをローカルにダウンロードするGoアプリケーション

## 機能要件

### 基本機能
- YouTube動画URLを引数として受け取る
- 動画をローカルに保存する
- 動画の形式選択（MP4、WEBM等）
- 音声のみのダウンロード（MP3、M4A等）

### コマンドライン引数
```
drop-tube [OPTIONS] <YouTube URL>
```

#### オプション

| オプション | 説明 | デフォルト値 |
|------------|------|-------------|
| `-o, --output <PATH>` | 出力ディレクトリの指定 | カレントディレクトリ |
| `-f, --format <FORMAT>` | 動画形式の指定（mp4, webm, best等） | best |
| `-a, --audio-only` | 音声のみダウンロード | false |
| `--audio-format <FORMAT>` | 音声形式の指定（mp3, m4a等） | mp3 |
| `-q, --quality <QUALITY>` | 品質指定（720p, 1080p, best等） | best |
| `--playlist` | プレイリスト全体をダウンロード | false |
| `-v, --verbose` | 詳細ログ出力 | false |
| `-h, --help` | ヘルプ表示 | - |


### 入力例
```bash
# 基本的な使用
drop-tube "https://www.youtube.com/watch?v=dQw4w9WgXcQ"

# 高品質MP4でダウンロード
drop-tube -f mp4 -q 1080p "https://www.youtube.com/watch?v=dQw4w9WgXcQ"

# 音声のみMP3でダウンロード
drop-tube -a --audio-format mp3 "https://www.youtube.com/watch?v=dQw4w9WgXcQ"

# 出力ディレクトリ指定
drop-tube -o ./my-videos "https://www.youtube.com/watch?v=dQw4w9WgXcQ"
```

## 技術仕様

### 使用技術
- Go 1.21以上
- YouTube動画取得: yt-dlp または 同等のライブラリ
- CLI: cobra または flag パッケージ
- 進捗表示: progressbar ライブラリ

### ディレクトリ構成
```
drop-tube/
├── cmd/
│   └── drop-tube/
│       └── main.go
├── internal/
│   ├── downloader/
│   │   └── youtube.go
│   ├── config/
│   │   └── config.go
│   └── cli/
│       └── cmd.go
├── pkg/
│   └── utils/
│       └── file.go
├── docs/
│   └── specification.md
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

### 依存関係
- github.com/spf13/cobra (CLI)
- github.com/schollz/progressbar/v3 (進捗表示)
- yt-dlp バイナリ (システム要件)

## 非機能要件

### パフォーマンス
- 同時ダウンロード数: 1
- ダウンロード進捗の表示
- 大容量ファイルに対応

### エラーハンドリング
- 無効なURL
- ネットワークエラー
- ディスク容量不足
- 権限エラー
- yt-dlpの未インストール

### ログ
- デフォルト: 基本的な進捗情報
- verbose モード: 詳細なデバッグ情報
- エラー情報の適切な出力

## 制約事項
- YouTube利用規約に準拠
- 著作権で保護されたコンテンツの取り扱い注意
- yt-dlpの事前インストールが必要
- Go 1.21以上が必要

## 今後の拡張予定
- 複数URL同時ダウンロード
- 設定ファイル対応
- GUI版の開発
- Docker対応
- 字幕ダウンロード機能