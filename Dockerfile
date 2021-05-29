FROM golang:alpine AS builder
VOLUME [ "/home" ]
WORKDIR /home
RUN go build main.go

FROM scratch
COPY --from=builder /home/main /home
ENTRYPOINT [ "/home/main" ]
