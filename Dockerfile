FROM golang:alpine AS builder

WORKDIR /go/src

ADD ./main.go .

RUN CGO_ENABLED=0 go build -o aimage -ldflags '-s' main.go

FROM scratch

COPY --from=builder /go/src/aimage /aimage

EXPOSE 8080

CMD ["/aimage"] 