FROM golang:1.19-alpine as builder

COPY .  /app/
WORKDIR /app/
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o courses-mail-service .

FROM alpine
WORKDIR /app/
COPY --from=builder /app/courses-mail-service .
COPY config/*.yml ./config/
COPY templates/*.html ./templates/

CMD [ "./courses-mail-service" ]
