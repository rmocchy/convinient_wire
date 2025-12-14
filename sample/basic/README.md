# Basic DI Sample

このサンプルは、Wireを使った基本的な依存性注入のパターンを示します。

## 構成

- Repository層: データアクセス
- Service層: ビジネスロジック
- Handler層: HTTPハンドラー

## 実行方法

```bash
# 依存関係のインストール
go mod download

# Wireによるコード生成
wire

# 実行
go run .
```

## ファイル構成

- `main.go`: エントリーポイント
- `wire.go`: Wire設定
- `wire_gen.go`: Wireが生成するファイル(自動生成)
- `repository/`: データアクセス層
- `service/`: ビジネスロジック層
- `handler/`: ハンドラー層
