# build stage
FROM golang:1.9-alpine AS build-env
RUN mkdir -p /go/src \
  && mkdir -p /go/bin \
  && mkdir -p /go/pkg
RUN apk update && apk add git

WORKDIR /go/src/myapp
ADD . .
RUN go get -u github.com/golang/dep/cmd/dep && \
    dep ensure && \
    go build -v -o greeter-api .

# final stage
FROM alpine:3.2

WORKDIR /app
COPY --from=build-env /go/src/myapp/greeter-api greeter-api
ENTRYPOINT [ "/app/greeter-api" ]
