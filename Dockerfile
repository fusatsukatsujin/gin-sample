# ビルドステージ
FROM golang:1.22.4-alpine AS builder

# 作業ディレクトリを設定
WORKDIR /app

# 依存関係ファイルをコピー
COPY go.mod go.sum ./

# 依存関係をダウンロード
RUN go mod download

# ソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# 実行ステージ
FROM alpine:latest

# 作業ディレクトリを設定
WORKDIR /root/

# ビルドステージから実行可能ファイルをコピー
COPY --from=builder /app/main .

COPY conf/app.ini ./conf/app.ini

EXPOSE 8080

# アプリケーションの実行
CMD ["./main"]