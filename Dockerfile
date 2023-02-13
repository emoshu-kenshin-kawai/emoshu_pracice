FROM golang:1.20-alpine
RUN go install github.com/cosmtrek/air@latest
WORKDIR /go/src/app
COPY . .
RUN go mod download
EXPOSE 3000
CMD ["air", "-c", ".air.toml"]