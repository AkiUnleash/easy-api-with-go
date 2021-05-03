# 2020/10/14最新versionを取得
FROM golang:1.15.2-alpine
# アップデートとgitのインストール
RUN apk update && apk add git
# appディレクトリの作成
RUN mkdir /go/src/app
# ワーキングディレクトリの設定
WORKDIR /go/src/app
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/app

RUN go get -u github.com/oxequa/realize \
  && go get github.com/go-sql-driver/mysql \
  && go get github.com/labstack/echo \
  && go get github.com/jinzhu/gorm \
  && go get gopkg.in/ini.v1 \
  && go get -u github.com/swaggo/swag/cmd/swag \
  && go get -u github.com/swaggo/echo-swagger

# CMD ["go", "env", "-w", "GO111MODULE=on"] && \
#   ["realize", "start"]