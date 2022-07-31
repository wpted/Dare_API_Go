FROM golang:latest
COPY go.mod go.sum /Users/edward/Developer/Go_Projects/dareAPI/
WORKDIR /Users/edward/Developer/Go_Projects/dareAPI/
RUN go mod download