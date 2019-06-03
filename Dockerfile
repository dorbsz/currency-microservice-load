FROM golang:1

RUN curl -sSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 && chmod +x /usr/local/bin/dep

COPY . /app
WORKDIR /app/
ENV GOPATH /app

#get dependencies from Gopkg.toml
RUN dep ensure

RUN GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build ./src/main.go




FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=0 /app/main .
CMD ["./main"]