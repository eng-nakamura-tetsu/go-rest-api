# バイナリ作成コンテナ
FROM golang:1.18 as deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpatch -ldflags "-w -s" -o app

# デプロイコンテナ
FROM debian:bullseye-alim as deploy

RUN apt-get update

COPY --from=deploy-builder /app/app .

CMD ["./app"]

# ホットリロード
FROM golang:1.18 as dev
WORKDIR /app
RUN go install github.com/cosmtrek/air@latest
CMD ["air"]
