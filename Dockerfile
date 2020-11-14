FROM golang:1.15

WORKDIR /airports-service
COPY . .

ENV GOFLAGS=-mod=mod
RUN go build -o server cmd/main.go

ENTRYPOINT ["./server"]