# ベースとなるDockerイメージ指定
FROM golang:latest
# コンテナ内に作業ディレクトリを作成
RUN mkdir /usr/local/go/src/go-sample
# コンテナログイン時のディレクトリ指定
WORKDIR /usr/local/go/src/go-sample
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /usr/local/go/src/go-sample