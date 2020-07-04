# ベースとなるDockerイメージ指定
FROM golang:latest
# コンテナ内に作業ディレクトリを作成
ENV GO111MODULE=on
RUN mkdir /go/src/go_todo \
    && apt-get update \
    && go get -u github.com/gin-gonic/gin \
    && go get -u github.com/jinzhu/gorm \
    && go get github.com/pilu/fresh \
    && go get github.com/joho/godotenv \
    && go get golang.org/x/crypto/bcrypt


# コンテナログイン時のディレクトリ指定
WORKDIR /go/src/go_todo/main




CMD ["fresh"]