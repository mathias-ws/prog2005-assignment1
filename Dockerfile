FROM golang:1.17-alpine AS builder

WORKDIR /uniinfo

COPY . .
RUN go-wrapper download
RUN go build -v

FROM alpine:3.5

WORKDIR /usr/local/bin

COPY --from=builder /uniinfoapp .

EXPOSE 80

CMD ["./uniinfoapp"]
