FROM golang:latest

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /usr/trackster-api/

COPY . .
RUN go mod download

RUN go build -o ./dist/main ./cmd

CMD ["./dist/main"]