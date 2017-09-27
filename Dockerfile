FROM golang:1.8.3 AS build
WORKDIR /go/src/github.com/turnerlabs/turner-defaultbackend
COPY Gopkg.toml .
COPY Gopkg.lock .
COPY main.go .

# install deps
RUN go get -v github.com/golang/dep/cmd/dep
RUN dep ensure -v

# compile server
RUN GOOS=linux GOARCH=386 go build -v -o app .

### run
FROM alpine:3.6
WORKDIR /root/
COPY --from=build /go/src/github.com/turnerlabs/turner-defaultbackend/app .
CMD ["./app"]
###