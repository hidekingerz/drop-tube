# DropTube

YouTube動画をダウンロードするコマンドラインアプリケーション

## 概要

DropTubeはYouTube動画をローカルにダウンロードするためのGoアプリケーションです。  
[yt-dlp](https://github.com/yt-dlp/yt-dlp)のGoラッパーとして開発されており、シンプルなコマンドライン操作で、様々な形式・品質の動画や音声をダウンロードできます。

## 機能

- YouTube動画のダウンロード
- 複数の出力形式対応（MP4、WEBM等）
- 音声のみダウンロード（MP3、M4A等）
- 品質選択（720p、1080p等）
- プレイリスト全体のダウンロード
- 進捗表示
- 詳細ログ出力

## インストール

### 前提条件

- Go 1.21以上
- yt-dlp（システムにインストール済みである必要があります）
- ffmpeg（高品質ダウンロードに必要）

### 依存関係のインストール

#### yt-dlpのインストール

```bash
# macOS (Homebrew)
brew install yt-dlp

# Linux (apt)
sudo apt install yt-dlp

#### ffmpegのインストール

ffmpegは映像と音声の結合、フォーマット変換に必要です。高品質な動画をダウンロードする場合は必須です。

```bash
# macOS (Homebrew)
brew install ffmpeg

# Linux (apt)
sudo apt install ffmpeg
```

### DropTubeのビルド

```bash
git clone https://github.com/hidekingerz/drop-tube.git
cd drop-tube
go build -o drop-tube cmd/drop-tube/main.go
```

### DropTubeのインストール

```bash
go install github.com/hidekingerz/drop-tube/cmd/drop-tube@latest
```

## 使用方法

### 基本的な使用

```bash
drop-tube "https://www.youtube.com/watch?v=dQw4w9WgXcQ"
```

### オプション

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

### 使用例

```bash
# 高品質MP4でダウンロード
drop-tube -f mp4 -q 1080p "https://www.youtube.com/watch?v=dQw4w9WgXcQ"

# 音声のみMP3でダウンロード
drop-tube -a --audio-format mp3 "https://www.youtube.com/watch?v=dQw4w9WgXcQ"

# 出力ディレクトリ指定
drop-tube -o ./my-videos "https://www.youtube.com/watch?v=dQw4w9WgXcQ"

# プレイリスト全体をダウンロード
drop-tube --playlist "https://www.youtube.com/playlist?list=PLxxxxxxxxxxxxxx"

# 詳細ログ付きでダウンロード
drop-tube -v "https://www.youtube.com/watch?v=dQw4w9WgXcQ"
```

## プロジェクト構成

```
drop-tube/
├── cmd/
│   └── drop-tube/
│       └── main.go         # エントリーポイント
├── internal/
│   ├── downloader/
│   │   └── youtube.go      # YouTube ダウンロード機能
│   ├── config/
│   │   └── config.go       # 設定管理
│   └── cli/
│       └── cmd.go          # CLI コマンド定義
├── pkg/
│   └── utils/
│       └── file.go         # ファイル操作ユーティリティ
├── docs/
│   ├── specification.md    # 仕様書
│   └── styleguide.md      # コーディングスタイルガイド
├── go.mod
├── go.sum
└── README.md
```

## 開発

### 依存関係

- [github.com/spf13/cobra](https://github.com/spf13/cobra) - CLI フレームワーク
- [github.com/schollz/progressbar/v3](https://github.com/schollz/progressbar) - 進捗表示

### テスト実行

```bash
go test ./...
```

### ビルド

```bash
go build -o drop-tube cmd/drop-tube/main.go
```

### コードフォーマット

```bash
gofmt -s -w .
goimports -w .
```

### 静的解析

```bash
go vet ./...
golint ./...
```

## 制約事項

- YouTube利用規約に準拠してご利用ください
- 著作権で保護されたコンテンツのダウンロードは適切な権限が必要です
- yt-dlpの事前インストールが必要です
- ffmpegの事前インストールが必要です（高品質ダウンロード時）
- Go 1.21以上が必要です

## ライセンス

このプロジェクトはMITライセンスの下で公開されています。詳細は[LICENSE](LICENSE)ファイルをご覧ください。

## 今後の予定

- 複数URL同時ダウンロード
- 設定ファイル対応
- GUI版の開発
- Docker対応
- 字幕ダウンロード機能

## 貢献

プルリクエストやイシューを歓迎します。開発に参加される場合は、[docs/styleguide.md](docs/styleguide.md)をご確認ください。