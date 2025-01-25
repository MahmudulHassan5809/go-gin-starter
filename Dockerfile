FROM golang:1.20-alpine
ENV GOPATH /go
ENV GIN_MODE release
RUN apk add --no-cache git gcc g++ make bash curl
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN curl -fLo /usr/local/bin/air https://raw.githubusercontent.com/cosmtrek/air/master/bin/linux/air && \
    chmod +x /usr/local/bin/air
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]
