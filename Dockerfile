FROM golang:1.14.2
WORKDIR /go/src/
COPY api.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/app .
CMD ["./app"]  