# TODO API

## 概要
Go/Ginを使った簡単なTODO管理のAPI。PostgreSQLをDBとして使用。
タスクのCRUD操作をHTTPで提供する。

## 技術スタック
| 項目 | 内容 |
|---|---|
| 言語 | Go |
| フレームワーク | gin（github.com/gin-gonic/gin）|
| データベース | PostgreSQL |
| DBドライバ | pgx/v5 （github.com/jackc/pgx/v5）|
| コンテナ | Docker |

## 機能
・タスク追加
・タスク一覧取得
・タスク取得(1件)
・タスク完了
・タスク削除

## 環境構築
### 前提条件
- Go 1.23以上
- Docker / Docker Compose

### 起動手順

1. リポジトリをクローン
```bash
git clone [your-repo-url]
cd todo-api
```

### 1.環境変数の設定
プロジェクトルートに `.env` を作成する。
```env
POSTGRES_USER=your_user
POSTGRES_PASSWORD=your_password
POSTGRES_DB=your_dbname
```

### 2.PostgreSQLの起動
```bash
docker-compose up -d
```
ホストのポート7777で起動する。

### 3.テーブルの作成
PostgreSQLコンテナに接続してテーブルを作成します。
```bash
# PostgreSQLに接続
docker exec -it todo-api-postgres psql -U your_user -d your_dbname

# 以下のSQLを実行
CREATE TABLE todos (
    id        SERIAL PRIMARY KEY,
    title     TEXT NOT NULL,
    completed BOOLEAN NOT NULL DEFAULT false
);

# 確認
\dt

# 終了
\q
```

### 4.アプリの起動
```bash
go run main.go
```
`http://localhost:8888` で起動する。

## API仕様
| メソッド | エンドポイント | 説明 |
|---|---|---|
| POST | /todos | タスクを作成 |
| GET | /todos | タスク一覧の取得 |
| GET | /todos/:id | タスクを1件取得 |
| PUT | /todos/:id | タスクを完了にする |
| DELETE | /todos/:id | タスクを削除する |

### リクエスト / レスポンス例

**POST /todos**
```json
// Request
{ "title": "買い物に行く" }

// Response 201
{ "id": 1, "title": "買い物に行く", "completed": false }
```

**GET /todos**
```json
// Response 200
[
  { "id": 1, "title": "買い物に行く", "completed": false },
  { "id": 2, "title": "本を読む", "completed": true }
]
```

**PUT /todos/:id**
```json
// Response 200
{ "id": 1, "title": "買い物に行く", "completed": true }
```

**DELETE /todos/:id**
```json
// Response 200
{ "message": "削除しました" }
```

## ディレクトリ構造
```
todo-api/
├── db
│   └── db.go          # データベース接続
├── handler
│   └── handler.go     # HTTPハンドラー
├── model
│   └── model.go       # データ構造
├── main.go            # エントリーポイント
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
```

## 工夫した点
### 1. ファイル分割
- `model/`: データ構造の定義
- `handler/`: HTTPリクエスト処理
- `db/`: データベース接続管理
- `main.go`: ルーティング設定のみ

このように責務を分離することで、例えばデータベースをPostgreSQLから
MySQLに変更する場合、`db/db.go`のみ修正すればよく、
他の部分に影響を与えません。

### 2. 環境変数管理
データベース接続情報を.envファイルで管理し、セキュリティを考慮。
例として`.env.example`を配置。

### 3. エラーハンドリング
適切にエラーを処理してHTTPステータスコード/エラーメッセージを返すように実装。

## 今後の改善案
- テストコードの追加
- バリデーションの強化
- ページネーション機能
- 認証機能の追加
- Docker Composeでアプリも起動