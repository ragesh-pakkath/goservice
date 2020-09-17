FROM golang
COPY . /app
WORKDIR /app
CMD ["go run first.go"]
