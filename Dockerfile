FROM golang:1.19-alpine as builder
WORKDIR /go/app
COPY . .
# RUN go build -v -o app app/main.go
RUN go build -v -o app main.go
FROM alpine
COPY --from=builder /go/app/ .
CMD ["/app"]