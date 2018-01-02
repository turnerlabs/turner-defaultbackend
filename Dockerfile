FROM golang:1.8.3 AS build
WORKDIR /go/src/github.com/turnerlabs/turner-defaultbackend
COPY Gopkg.toml .
COPY Gopkg.lock .
COPY main.go .

# install deps
RUN go get -v github.com/golang/dep/cmd/dep
RUN dep ensure -v

# compile server
RUN CGO_ENABLED=0 GOOS=linux go build -v -o app .

### run
FROM scratch
WORKDIR /root/
COPY --from=build /go/src/github.com/turnerlabs/turner-defaultbackend/app .
CMD ["./app"]
###
