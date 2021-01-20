# 使用するGolangのイメージを指定する
FROM golang:1.15.0

ENV GOBIN=/go/bin \
    GOPATH=/go

# ワーキングディレクトリを指定する
WORKDIR /go/src/github.com/sasa-nori/nyaitter

# 依存ライブラリをダウンロード
COPY go.mod go.sum ./
RUN go mod download

# book-supplement直下のディレクトリをコンテナ上に載せる
COPY . .

# go install で /go/bin 配下にバイナリがビルドされる
RUN GOOS=linux GOARCH=amd64 go install

# image自体に最初からPATHに/go/binのパスが設定されているため、バイナリファイルを書くだけで動かすことができる
ENTRYPOINT [ "nyaitter" ]
