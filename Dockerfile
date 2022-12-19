FROM golang:1.19 as builder

WORKDIR /build
COPY . .
RUN go mod download && make compile


FROM gcr.io/distroless/static

COPY --from=builder /build/outbox-pattern /outbox-pattern
CMD ["/outbox-pattern"]
