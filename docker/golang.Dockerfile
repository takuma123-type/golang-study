FROM --platform=linux/arm64/v8 golang:1.21

WORKDIR /go/src

# 必要なツールと依存関係をインストール
RUN apt-get update && \
    apt-get install -y git && \
    go install github.com/cosmtrek/air@latest

# src/go.mod と src/go.sum のみを先にコピー
COPY src/go.* .

# 依存関係をインストール
RUN go mod download

# ソースコードと設定ファイルをコピー
COPY src/ .

CMD ["air", "-c", ".air.toml"]
