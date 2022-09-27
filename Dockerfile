FROM golang:1.18.2-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build  -o /out/main ./
ENTRYPOINT ["/out/main"]