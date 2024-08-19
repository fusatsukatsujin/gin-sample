# ginサンプルアプリケーション

このプロジェクトは、ginのシンプルなサンプルアプリケーションです。Go言語で実装されており、PostgreSQLデータベースを使用しています。

## 機能

- ログイン
- メンバーの追加
- メンバー一覧の取得
- メンバーの検索

## 必要条件

- Go 1.20以上
- PostgreSQL 12以上

## セットアップ

1. リポジトリをクローンします：
   ```
   git clone https://github.com/fusatsukatsujin/gin-sample.git
   ```

2. 必要な依存関係をインストールします：
   ```
   go mod tidy
   ```

3. PostgreSQLデータベースを設定します。

4. `.env`ファイルを作成し、以下の環境変数を設定します：
   ```
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=yourusername
   DB_PASSWORD=yourpassword
   DB_NAME=membersdb
   ```

5. アプリケーションを実行します：
   ```
   go run main.go
   ```

## API エンドポイント

- `POST /login`: ログイン
- `POST /api/members`: 新しいメンバーを追加
- `GET /api/members`: すべてのメンバーを取得
- `GET /api/members/:id`: メンバーを検索

ログイン以外のエンドポイントには、認証が必要です。
`Authorization`ヘッダーにJWTトークンを含める必要があります。

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