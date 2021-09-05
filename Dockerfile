FROM golang:latest as builder
WORKDIR /demo_project_api

COPY go.mod ./

RUN go mod download

COPY . ./

RUN go build    #-o /docker-demo-go

EXPOSE 8081
EXPOSE 27017
EXPOSE 6379

CMD ["./demo_project"]
