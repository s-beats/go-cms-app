# HanndaiConnectBackend

## 環境構築

- air（ホットリロード）のインストール

```
$ curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

$ curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

$ air -v
```

### サーバー初期化

- Webサーバー、DB、ストレージ起動

```sh
$ make start
```

### DB初期化

```sh
$ make init-db-local
```

### ストレージ初期化

```sh
$ make init-storage-local
```