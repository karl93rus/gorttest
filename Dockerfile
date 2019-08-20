FROM golang:1.12-buster

RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go get github.com/lib/pq

RUN go build -o server .
CMD ["./server"]
