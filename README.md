# 色当てゲーム (Color Guessing Game)

シンプルな色当てゲームです。ランダムに表示される色を当てるゲームで、10問解答すると結果が表示されます。

## 機能

- ランダムに色が表示され、選択肢から正解を選ぶ
- 5問正解すると選択肢が3つから4つに増える
- 10問解答すると結果が表示される
- 日本語と英語の切り替えが可能

詳細な設計については[詳細設計書](docs/design.md)を参照してください。

## 技術スタック

- フロントエンド: React
- バックエンド: Go
- コンテナ化: Docker, Docker Compose
- テスト: Jest (フロントエンド), Go testing (バックエンド)

## セットアップ

### 前提条件

- Docker と Docker Compose がインストールされていること

### 起動方法

```bash
# アプリケーションの起動
docker-compose up --build
```

ブラウザで http://localhost:3000 にアクセスするとゲームが表示されます。

## 開発

### テスト実行

```bash
# すべてのテストを実行
make test

# フロントエンドのテストのみ実行
make test-frontend

# バックエンドのテストのみ実行
make test-backend
```

### カバレッジレポート生成

```bash
# すべてのカバレッジレポートを生成
make coverage

# フロントエンドのカバレッジレポートのみ生成
make coverage-frontend

# バックエンドのカバレッジレポートのみ生成
make coverage-backend
```

カバレッジレポートは `coverage` ディレクトリに生成されます。

### ローカル環境での開発

```bash
# フロントエンド
cd frontend
npm install
npm start

# バックエンド
cd backend
go run main.go
```

### クリーンアップ

```bash
make clean
```

## プロジェクト構成

```
.
├── backend/             # バックエンドのソースコード (Go)
├── frontend/            # フロントエンドのソースコード (React)
├── docs/                # ドキュメント
│   └── design.md        # 詳細設計書
├── Dockerfile.backend   # バックエンド用Dockerfile
├── Dockerfile.frontend  # フロントエンド用Dockerfile
├── docker-compose.yml   # Docker Compose設定
└── Makefile             # ビルド・テスト用Makefile
```

## ライセンス

MIT