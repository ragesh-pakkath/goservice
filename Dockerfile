FROM golang
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o first .
CMD [""/app/first""]
