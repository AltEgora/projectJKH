FROM golang:1.25.4

COPY . /app

#COPY ./../static /app

WORKDIR /app

#RUN go mod init ConsultantBack

RUN go mod tidy

#RUN go run /app/cmd/main.go