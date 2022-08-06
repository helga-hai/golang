FROM golang:latest

RUN mkdir /main
ADD app/server.go /main
ADD go.mod /main
WORKDIR /main
RUN ls /main
RUN go build server.go
# RUN go run main/server.go
CMD ["go", "run", "server.go"]