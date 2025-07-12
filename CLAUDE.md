# DropTube プロジェクトメモリ

## プロジェクト概要
YouTube動画をダウンロードするGoアプリケーション「DropTube」の開発プロジェクト

## ドキュメント準拠
コード生成・修正時は必ず以下のドキュメントに準拠すること：

[specification](docs/specification.md)
[styleguide](docs/styleguide.md)

### 重要な指示
1. **仕様書準拠**: `docs/specification.md`で定義された機能要件、技術仕様、ディレクトリ構成に厳密に従う
2. **コーディング規約**: `docs/styleguide.md`で定義された命名規則、フォーマット、エラーハンドリング等を遵守
3. **アーキテクチャ**: internal/パッケージ構成とcmd/エントリーポイント設計を維持
4. **品質保証**: テスト駆動開発、GitHub ActionsでのCI/CD、gofmt/goimports/golintによる自動品質チェックを実施

## 品質管理の実施状況
### 自動化済み
- GitHub Actionsでのコードフォーマットチェック（gofmt）
- インポート整理チェック（goimports）
- 自動ビルドテスト
- 単体テスト実行
- developおよびmainブランチでのCI実行

### 手動実施推奨
開発中は以下のコマンドで品質を保つこと：
```bash
# フォーマット修正
gofmt -s -w .
goimports -w .

# 静的解析
go vet ./...
golint ./...

# テスト実行
go test -v ./...

# ビルド確認
go build -v ./...
```

## プロジェクト固有ルール
- YouTube利用規約遵守を最優先
- yt-dlpバイナリ依存の明示的エラーハンドリング
- 進捗表示とログ出力の一貫性維持
- セキュリティ考慮（著作権保護コンテンツ対応）