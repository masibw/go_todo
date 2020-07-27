# ベースとなるDockerイメージ指定
FROM golang:latest
# コンテナ内に作業ディレクトリを作成
ENV GO111MODULE=on
ENV DOCKERIZE_VERSION v0.6.1

RUN mkdir /go/src/go_todo \
    && apt-get update \
    && apt-get install -y wget \
    && go get -u github.com/gin-gonic/gin \
    && go get -u github.com/jinzhu/gorm \
    && go get github.com/pilu/fresh \
    && go get github.com/joho/godotenv \
    && go get golang.org/x/crypto/bcrypt \
    && wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz


# コンテナログイン時のディレクトリ指定
WORKDIR /go/src/go_todo/cmd/go_todo


