FROM golang:1.17-alpine
EXPOSE 8080
RUN mkdir /go/src/hw07-graphql
WORKDIR /go/src/hw07-graphql
COPY . .
RUN go mod tidy
CMD go run server/server.go
