FROM golang:1.19-buster

WORKDIR /go/src/server

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ./out/server .

EXPOSE 8080 8080

CMD ["./out/server"]