# ginサンプルアプリケーション

このプロジェクトは、ginのシンプルなサンプルアプリケーションです。Go言語で実装されており、PostgreSQLデータベースを使用しています。

## 機能

- ログイン
- メンバーの追加
- メンバー一覧の取得
- メンバーの検索

## 必要条件

- Go 1.22.4以上
- PostgreSQL 13以上

## セットアップ

1. リポジトリをクローンします：
   ```
   git clone https://github.com/fusatsukatsujin/gin-sample.git
   ```

2. 必要な依存関係をインストールします：
   ```
   go mod tidy
   ```

3. アプリケーションとDBを起動します。
```
docker compose up -d
```

4. DBにテーブルを作成します。
以下のDDLを流してください。
docs/db/create.sql

## API エンドポイント

- `POST /login`: ログイン
- `POST /api/members`: 新しいメンバーを追加
- `GET /api/members`: すべてのメンバーを取得
- `GET /api/members/:id`: メンバーを検索

ログイン以外のエンドポイントには、認証が必要です。
`Authorization`ヘッダーに`/login`で発行されたJWTトークンを含める必要があります。

## 使用例

ログイン：
```
curl -X POST http://localhost:8080/api/login \
-d "username=admin" \
-d "password=password"
```

新しいメンバーを追加：
```
curl -X POST http://localhost:8080/api/members \
     -H "Authorization: Bearer <JWTトークン>" \
     -d "name=佐藤" \
     -d "age=25" \
     -d "sex=male"
```