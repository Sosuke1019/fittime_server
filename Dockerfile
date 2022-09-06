# ベースとなるDockerイメージ指定
FROM golang:1.19-alpine as server-build

# コンテナ内の作業ディレクトリを作成し、そこを指定
WORKDIR  /go/src/fittime_server
# ローカルの現在のディレクトリから、コンテナの作業ディレクトリにコピー
COPY . .

RUN go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]