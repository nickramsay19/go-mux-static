FROM golang:1.14

WORKDIR /
COPY . .

RUN go get github.com/gorilla/mux
RUN go build -o main .

CMD ["./main"]