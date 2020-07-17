FROM golang:1.14-alpine AS build
ENV TZ Asia/Tokyo
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN apk update && apk add git alpine-sdk tzdata
WORKDIR /go/src/github.com/inoue0124/base_app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN make deps clean bin/base_app

FROM alpine:3.10
ENV TZ Asia/Tokyo
RUN apk --update --no-cache add ca-certificates tzdata
COPY --from=build /go/src/github.com/inoue0124/base_app/bin/base_app /base_app
CMD ["/base_app"]
