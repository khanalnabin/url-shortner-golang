FROM golang:latest

WORKDIR /output

COPY . /output

RUN go mod tidy

RUN go build -o main .

EXPOSE 8080

CMD ["/output/main"]
