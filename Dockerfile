######## Start a builder stage #######
FROM golang:1.15-alpine as builder

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o metadata-scraper

######## Start a new stage from scratch #######
FROM alpine:latest

WORKDIR /opt/
COPY --from=builder /app/metadata-scraper .

CMD ["./nginx-conf-generator"]