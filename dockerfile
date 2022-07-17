FROM golang:latest


LABEL maintainer="yogie <hello@yogie.tama.com>"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod donwload

COPY . .

RUN go build

CMD ["./go-clean-api"]